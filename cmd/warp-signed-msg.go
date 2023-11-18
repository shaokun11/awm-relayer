package cmd

import (
	"encoding/json"
	"errors"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/awm-relayer/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

func newValidator() *bls.SecretKey {
	// windows不可用
	sk, _ := bls.NewSecretKey()
	return sk
}

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
		//keyBytes, _ := hex.DecodeString("0040c137287b7d169e076cb80d69bf98222ea780693b3bb4989c5998508490ff")
		//if err != nil {
		//	return err
		//}
		//pk, err := bls.SecretKeyFromBytes(keyBytes)
		//if err != nil {
		//	return err
		//}
		//signature := bls.Sign(pk, message.Bytes())

		vdr1sk := newValidator()
		vdr2sk := newValidator()
		vdr3sk := newValidator()
		count := 3
		signatures := make([]*bls.Signature, 0, count)
		sig1 := bls.Sign(vdr1sk, message.Bytes())
		sig2 := bls.Sign(vdr2sk, message.Bytes())
		sig3 := bls.Sign(vdr3sk, message.Bytes())
		signatures = append(signatures, sig1)
		signatures = append(signatures, sig2)
		signatures = append(signatures, sig3)
		bitSet := set.NewBits()
		bitSet.Add(0)
		bitSet.Add(1)
		bitSet.Add(2)
		aggSig, err := bls.AggregateSignatures(signatures)
		//signedMsg, err := warp.NewMessage(message, &warp.BitSetSignature{
		//	Signers:   bitSet.Bytes(),
		//	Signature: [96]byte(bls.SignatureToBytes(aggSig)),
		//})
		var ret = struct {
			Signature string `json:"signature"`
			BitSet    string `json:"bit_set"`
		}{
			Signature: hexutil.Encode(bls.SignatureToBytes(aggSig)),
			BitSet:    hexutil.Encode(bitSet.Bytes()),
		}
		marshal, err := json.Marshal(ret)
		if err != nil {
			return err
		}
		utils.ToNode(string(marshal))
		return nil
	},
}
