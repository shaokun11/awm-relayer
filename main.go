// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ava-labs/subnet-evm/core/types"
	"math/big"
)

//export parseTx
func parseTx(str string) string {
	var tx types.Transaction
	decodeString, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	err = tx.UnmarshalBinary(decodeString)
	if err != nil {
		panic(err)
	}
	v, r, s := tx.RawSignatureValues()
	tx_ := struct {
		Hash       string           `json:"hash"`
		ChainId    *big.Int         `json:"chain_id"`
		To         string           `json:"to"`
		Data       string           `json:"data"`
		AccessList types.AccessList `json:"access_list"`
		GasPrice   *big.Int         `json:"gas_price"`
		Gas        uint64           `json:"gas"`
		Value      *big.Int         `json:"value"`
		Nonce      uint64           `json:"nonce"`
		GasFeeCap  *big.Int         `json:"gas_fee_cap"`
		GasTipCap  *big.Int         `json:"gas_tip_cap"`
		V          *big.Int         `json:"v"`
		S          *big.Int         `json:"s"`
		R          *big.Int         `json:"r"`
	}{

		Hash:       tx.Hash().Hex(),
		AccessList: tx.AccessList(),
		ChainId:    tx.ChainId(),
		Value:      tx.Value(),
		Nonce:      tx.Nonce(),
		To:         tx.To().Hex(),
		GasFeeCap:  tx.GasFeeCap(),
		R:          r,
		S:          s,
		V:          v,
		GasTipCap:  tx.GasTipCap(),
		GasPrice:   tx.GasPrice(),
		Gas:        tx.Gas(),
		Data:       hex.EncodeToString(tx.Data()),
	}
	marshal, err := json.Marshal(tx_)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}

func main() {

}
