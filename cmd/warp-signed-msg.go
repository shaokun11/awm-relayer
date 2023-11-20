package cmd

import (
	"encoding/json"
	"errors"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/awm-relayer/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func NewValidator(p string) *bls.SecretKey {
	file, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	sk, _ := bls.SecretKeyFromBytes(content)
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
		//message, err := warp.ParseUnsignedMessage(decode)
		//if err != nil {
		//	return err
		//}
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

		vdr1sk := NewValidator("keys/n1.key")
		vdr2sk := NewValidator("keys/n2.key")
		vdr3sk := NewValidator("keys/n3.key")
		vdr4sk := NewValidator("keys/n4.key")
		count := 4
		signatures := make([]*bls.Signature, 0, count)
		sig1 := bls.Sign(vdr1sk, decode)
		sig2 := bls.Sign(vdr2sk, decode)
		sig3 := bls.Sign(vdr3sk, decode)
		sig4 := bls.Sign(vdr4sk, decode)
		signatures = append(signatures, sig1)
		signatures = append(signatures, sig2)
		signatures = append(signatures, sig3)
		signatures = append(signatures, sig4)
		bitSet := set.NewBits()
		bitSet.Add(0)
		bitSet.Add(1)
		bitSet.Add(2)
		bitSet.Add(3)
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
