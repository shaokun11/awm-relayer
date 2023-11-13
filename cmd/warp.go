package cmd

import (
	"encoding/json"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/awm-relayer/utils"
	"github.com/ava-labs/subnet-evm/utils/predicate"
	"github.com/ava-labs/subnet-evm/warp/payload"
	warp2 "github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

var warpCmd = &cobra.Command{
	Use:   "warp",
	Short: "parse to access list data to warp message bytes",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		decodeString, err := hexutil.Decode(args[0])
		unpackedPredicateBytes, err := predicate.UnpackPredicate(decodeString)
		warpMessage, err := warp.ParseMessage(unpackedPredicateBytes)
		addressedPayload, err := payload.ParseAddressedPayload(warpMessage.UnsignedMessage.Payload)
		output, err := warp2.WarpABI.PackOutput("getVerifiedWarpMessage",
			warp2.WarpMessage{
				SourceChainID:       common.Hash(warpMessage.SourceChainID),
				OriginSenderAddress: addressedPayload.SourceAddress,
				DestinationChainID:  addressedPayload.DestinationChainID,
				DestinationAddress:  addressedPayload.DestinationAddress,
				Payload:             addressedPayload.Payload,
			},
			true,
		)
		if err != nil {
			panic(err)
		}
		var warp = struct {
			Data string `json:"data"`
		}{
			Data: hexutil.Encode(output),
		}
		marshal, err := json.Marshal(warp)
		if err != nil {
			panic(err)
		}
		utils.ToNode(string(marshal))
	},
}
