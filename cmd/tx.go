package cmd

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ava-labs/awm-relayer/utils"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"math/big"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "parse receive cross chain raw message to struct info ",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var tx types.Transaction
		// 02f903a3a0dd084a5ba12bbdae7ccda7d2322009b6ce21267530467f9aa57eac8e71ce6f3f80843b9aca0085010c388d008321a1609469b5c70fabf28b48ec130984cd439efb11ca634780b844ccb5f8090000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a100ff48a37cab9f87c8b5da933da46ea1a5fb80f902d0f902cd940200000000000000000000000000000000000005f902b5a00000000005391bd1deea214bd454066d1073305ef65c9b99d716c98756b2851da0f62df469c292000001f200000000000069b5c70fabf28b48ec130984cd439efba011ca6347dd084a5ba12bbdae7ccda7d2322009b6ce21267530467f9aa57eac8ea071ce6f3f69b5c70fabf28b48ec130984cd439efb11ca6347000001a000000000a00000000000000000000000000000000000000000000000000000002000000000a00000000000000000000000000000000000000000000000000000001400000000a000000000000000008db97c7cece249c2b98bdc0226cc4c2a57bf52fc00000000a00000000000000000abcedf1234abcedf1234abcedf1234abcedf123400000000a0000000000000000000000000000000000000000000000000000186a000000000a0000000000000000000000000000000000000000000000000000000e000000000a00000000000000000000000000000000000000000000000000000010000000000a00000000000000000000000000000000000000000000000000000012000000000a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000002acafebabea0cafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabea0cafebabecafe0000000000000000000000000000000000000000000000000000a0000000011790ffa1094327a6512013a87e0d8269f42c290cf3782737b480de2ba02bc9774a03ded8112df3e38d7753bc0ddb1dc548221084288c86ad05169ec0f7a0237f561a84348764e92d8000a36b57f7550ad8b68e1261c3f575f4e82583d3d8a07ee88ee5eaff000000000000000000000000000000000000000000000000000080a0712a564dc43965e0266c991d766c730e89623cc647dd07ad35bf482159a75213a04cf730a1c063e6e435fe607ce976cdc71844dda437997835580a8caed6013fc6
		decodeString, err := hexutil.Decode(args[0])
		if err != nil {
			return err
		}
		err = tx.UnmarshalBinary(decodeString)
		if err != nil {
			return err
		}
		v, r, s := tx.RawSignatureValues()
		tx_ := struct {
			Hash       string           `json:"hash"`
			ChainId    string           `json:"chain_id"`
			To         string           `json:"to"`
			Data       string           `json:"data"`
			AccessList types.AccessList `json:"access_list"`
			GasPrice   *big.Int         `json:"gas_price"`
			Gas        uint64           `json:"gas"`
			Value      *big.Int         `json:"value"`
			Nonce      uint64           `json:"nonce"`
			GasFeeCap  *big.Int         `json:"gas_fee_cap"`
			GasTipCap  *big.Int         `json:"gas_tip_cap"`
			V          string           `json:"v"`
			S          string           `json:"s"`
			R          string           `json:"r"`
		}{

			Hash:       tx.Hash().Hex(),
			AccessList: tx.AccessList(),
			ChainId:    tx.ChainId().String(),
			Value:      tx.Value(),
			Nonce:      tx.Nonce(),
			To:         tx.To().Hex(),
			GasFeeCap:  tx.GasFeeCap(),
			R:          r.String(),
			S:          s.String(),
			V:          v.String(),
			GasTipCap:  tx.GasTipCap(),
			GasPrice:   tx.GasPrice(),
			Gas:        tx.Gas(),
			Data:       hex.EncodeToString(tx.Data()),
		}
		marshal, err := json.Marshal(tx_)
		if err != nil {
			return err
		}
		utils.ToNode(string(marshal))
		return nil
	},
}
