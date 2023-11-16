package cmd

import (
	"encoding/json"
	"errors"
	"github.com/ava-labs/avalanchego/ids"
	warp2 "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/awm-relayer/utils"
	paylaod2 "github.com/ava-labs/subnet-evm/warp/payload"
	"github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"strconv"
)

var warpUnsignedMsgCmd = &cobra.Command{
	Use:   "warp-unsigned-msg",
	Short: "warp the send payload to unsigned message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 6 {
			return errors.New("the params should be: payload sourceChainID networkID sourceAddress destinationChainID destinationAddress")
		}
		payload, err := hexutil.Decode(args[0])
		if err != nil {
			return err
		}
		sourceChainID, err := ids.FromString(args[1])
		if err != nil {
			return err
		}
		networkID, err := strconv.ParseUint(args[2], 10, 32)
		if err != nil {
			return err
		}
		var (
			sourceAddress      = common.HexToAddress(args[3])
			destinationChainID = common.HexToHash(args[4])
			destinationAddress = common.HexToAddress(args[5])
		)
		addressedPayload, err := paylaod2.NewAddressedPayload(
			sourceAddress,
			destinationChainID,
			destinationAddress,
			payload,
		)
		unsignedWarpMessage, err := warp2.NewUnsignedMessage(
			uint32(networkID),
			sourceChainID,
			addressedPayload.Bytes(),
		)
		if err != nil {
			return err
		}
		topic := []common.Hash{
			warp.WarpABI.Events["SendWarpMessage"].ID,
			destinationChainID,
			destinationAddress.Hash(),
			sourceAddress.Hash(),
		}

		var data = struct {
			Topics []common.Hash `json:"topics"`
			Data   string        `json:"data"`
		}{
			Topics: topic,
			Data:   hexutil.Encode(unsignedWarpMessage.Bytes()),
		}
		marshal, err := json.Marshal(data)
		if err != nil {
			return err
		}
		utils.ToNode(string(marshal))
		return nil
	},
}
