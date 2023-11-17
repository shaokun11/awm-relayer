package cmd

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/awm-relayer/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

var warpSignedMsgCmd = &cobra.Command{
	Use:   "warp-signed-msg",
	Short: "warp the send unsigned message to signed warp message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("the params should be: payload")
		}
		decode, err := hexutil.Decode(args[0])
		if err != nil {
			return err
		}
		// 这里是的bytes的是传递的数据
		message, err := warp.ParseUnsignedMessage(decode)
		if err != nil {
			return err
		}
		// subnet中通过rpc 像avalanchego请求签名
		// sig, err := b.warpSigner.Sign(unsignedMessage)
		// 此处直接返回签名
		keyBytes, _ := hex.DecodeString("0040c137287b7d169e076cb80d69bf98222ea780693b3bb4989c5998508490ff")
		if err != nil {
			return err
		}
		pk, err := bls.SecretKeyFromBytes(keyBytes)
		if err != nil {
			return err
		}
		signature := bls.Sign(pk, message.Bytes())
		count := 5
		signatures := make([]*bls.Signature, 0, count)
		bitSet := set.NewBits()
		for i := 0; i < count; i++ {
			signatures = append(signatures, signature)
			bitSet.Add(i)
		}
		aggSig, err := bls.AggregateSignatures(signatures)
		signedMsg, err := warp.NewMessage(message, &warp.BitSetSignature{
			Signers:   bitSet.Bytes(),
			Signature: *(*[bls.SignatureLen]byte)(bls.SignatureToBytes(aggSig)),
		})
		var ret = struct {
			Signature string `json:"signature"`
		}{
			Signature: hexutil.Encode(signedMsg.Bytes()),
		}
		marshal, err := json.Marshal(ret)
		if err != nil {
			return err
		}
		utils.ToNode(string(marshal))
		return nil
	},
}
