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

var accessListCmd = &cobra.Command{
	Use:   "access-list",
	Short: "parse to access list data to warp message bytes",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		decodeString, err := hexutil.Decode(args[0])
		if err != nil {
			return err
		}
		unpackedPredicateBytes, err := predicate.UnpackPredicate(decodeString)
		if err != nil {
			return err
		}
		warpMessage, err := warp.ParseMessage(unpackedPredicateBytes)
		if err != nil {
			return err
		}
		addressedPayload, err := payload.ParseAddressedPayload(warpMessage.UnsignedMessage.Payload)
		message := warp2.WarpMessage{
			SourceChainID: common.Hash(warpMessage.SourceChainID),
			//OriginSenderAddress: addressedPayload.SourceAddress,
			OriginSenderAddress: common.HexToAddress(args[1]),
			DestinationChainID:  addressedPayload.DestinationChainID,
			//DestinationAddress:  addressedPayload.DestinationAddress,
			// 因为改过TeleporterMessager的sol文件,所以这个地址可能与其他链的地址不一样
			DestinationAddress: common.HexToAddress(args[1]),
			Payload:            addressedPayload.Payload,
		}
		//println("SourceChainID", message.SourceChainID.Hex())
		//println("OriginSenderAddress", message.OriginSenderAddress.Hex())
		//println("DestinationChainID", message.DestinationChainID.Hex())
		//println("DestinationAddress", message.DestinationAddress.Hex())
		output, err := warp2.WarpABI.PackOutput("getVerifiedWarpMessage",
			message,
			true,
		)
		if err != nil {
			return err
		}
		var warp = struct {
			Data string `json:"data"`
		}{
			Data: hexutil.Encode(output),
		}
		marshal, err := json.Marshal(warp)
		if err != nil {
			return err
		}
		utils.ToNode(string(marshal))
		return nil
	},
}
