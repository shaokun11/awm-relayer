// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"go.uber.org/zap"
)

var _ validators.State = &CanonicalValidatorClient{}

// CanonicalValidatorClient wraps platformvm.Client and implements validators.State
type CanonicalValidatorClient struct {
	client platformvm.Client
	logger logging.Logger
}

func NewCanonicalValidatorClient(logger logging.Logger, client platformvm.Client) *CanonicalValidatorClient {
	return &CanonicalValidatorClient{
		client: client,
		logger: logger,
	}
}

func (v *CanonicalValidatorClient) GetMinimumHeight(ctx context.Context) (uint64, error) {
	return v.client.GetHeight(ctx)
}

func (v *CanonicalValidatorClient) GetCurrentHeight(ctx context.Context) (uint64, error) {
	return v.client.GetHeight(ctx)
}

func (v *CanonicalValidatorClient) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	return v.client.ValidatedBy(ctx, chainID)
}

// Gets the current validator set of the given subnet ID, including the validators' BLS public
// keys. The implementation currently makes two RPC requests, one to get the subnet validators,
// and another to get their BLS public keys. This is necessary in order to enable the use of
// the public APIs (which don't support "GetValidatorsAt") because BLS keys are currently only
// associated with primary network validation periods. If ACP-13 is implementated in the future
// (https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/13-subnet-only-validators.md), it
// may become possible to reduce this to a single RPC request that returns both the subnet validators
// as well as their BLS public keys.
func (v *CanonicalValidatorClient) getCurrentValidatorSet(
	ctx context.Context,
	subnetID ids.ID) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	// Get the current subnet validators. These validators are not expected to include
	// BLS signing information given that addPermissionlessValidatorTx is only used to
	// add primary network validators.
	subnetVdrs, err := v.client.GetCurrentValidators(ctx, subnetID, nil)
	if err != nil {
		return nil, err
	}

	// Look up the primary network validators of the NodeIDs validating the subnet
	// in order to get their BLS keys.
	res := make(map[ids.NodeID]*validators.GetValidatorOutput, len(subnetVdrs))
	subnetNodeIDs := make([]ids.NodeID, 0, len(subnetVdrs))
	for _, subnetVdr := range subnetVdrs {
		subnetNodeIDs = append(subnetNodeIDs, subnetVdr.NodeID)
		res[subnetVdr.NodeID] = &validators.GetValidatorOutput{
			NodeID: subnetVdr.NodeID,
			Weight: subnetVdr.Weight,
		}
	}
	primaryVdrs, err := v.client.GetCurrentValidators(ctx, ids.Empty, subnetNodeIDs)
	if err != nil {
		return nil, err
	}

	// Set the BLS keys of the result.
	for _, primaryVdr := range primaryVdrs {
		// We expect all of the primary network validators to already be in `res` because
		// we filtered the request to node IDs that were identified as validators of the
		// specific subnet ID.
		vdr, ok := res[primaryVdr.NodeID]
		if !ok {
			v.logger.Warn(
				"Unexpected primary network validator returned by getCurrentValidators request",
				zap.String("subnetID", subnetID.String()),
				zap.String("nodeID", primaryVdr.NodeID.String()))
			continue
		}

		// Validators that do not have a BLS public key registered on the P-chain are still
		// included in the result because they affect the stake weight of the subnet validators.
		// Such validators will not be queried for BLS signatures of warp messages. As long as
		// sufficient stake percentage of subnet validators have registered BLS public keys,
		// messages can still be successfully relayed.
		if primaryVdr.Signer != nil {
			vdr.PublicKey = primaryVdr.Signer.Key()
		}
	}

	return res, nil
}

// Gets the validator set of the given subnet at the given P-chain block height.
// Attempts to use the "getValidatorsAt" API first. If not available, falls back
// to use "getCurrentValidators", ignoring the specified P-chain block height.
func (v *CanonicalValidatorClient) GetValidatorSet(
	ctx context.Context,
	height uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	// First, attempt to use the "getValidatorsAt" RPC method. This method may not be available on
	// all API nodes, in which case we can fall back to using "getCurrentValidators" if needed.
	res, err := v.client.GetValidatorsAt(ctx, subnetID, height)
	if err != nil {
		v.logger.Debug(
			"P-chain RPC to getValidatorAt returned error. Falling back to getCurrentValidators",
			zap.String("subnetID", subnetID.String()),
			zap.Uint64("pChainHeight", height),
			zap.Error(err))
		return v.getCurrentValidatorSet(ctx, subnetID)
	}
	return res, nil
}