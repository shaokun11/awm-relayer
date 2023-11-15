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
		if len(args) != 4 {
			return errors.New("the four params should be: payload sourceChainID networkID sourceAddress")
		}
		//0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000040000000000000000000000008db97c7cece249c2b98bdc0226cc4c2a57bf52fc00000000000000000000000069b5c70fabf28b48ec130984cd439efb11ca634700000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
		decode, err := hexutil.Decode(args[0])
		if err != nil {
			return err
		}
		input, err := warp.UnpackSendWarpMessageInput(decode)
		if err != nil {
			return err
		}
		//sourceChainID, err := ids.FromString("DFcwVrEKk46moY2iLTBA6r72HtWEMPf8sr81zXGK5Mno66AKo")
		sourceChainID, err := ids.FromString(args[1])
		if err != nil {
			return err
		}
		networkID, err := strconv.ParseUint(args[2], 10, 32)
		if err != nil {
			return err
		}
		var (
			destinationChainID = input.DestinationChainID
			//sourceAddress      = common.HexToAddress("0xAb72527b0F669B18f8322437b4eB98C2DaD92bf7")
			sourceAddress      = common.HexToAddress(args[3])
			destinationAddress = input.DestinationAddress
			payload            = input.Payload
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
			Topic   []common.Hash `json:"topic"`
			Payload string        `json:"payload"`
		}{

			Topic:   topic,
			Payload: hexutil.Encode(unsignedWarpMessage.Bytes()),
		}
		marshal, err := json.Marshal(data)
		if err != nil {
			return err
		}
		utils.ToNode(string(marshal))
		return nil
	},
}
