package relayer

import (
	"bytes"
	"encoding/json"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"net/http"
)

// RPCRequest 表示 JSON-RPC 请求的结构
type RPCRequest struct {
	JSONRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  []map[string]any `json:"params"`
	ID      int              `json:"id"`
}

// RPCResponse 表示 JSON-RPC 响应的结构
type RPCResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *RPCError       `json:"error"`
	ID     int             `json:"id"`
}

// RPCError 表示 JSON-RPC 错误的结构
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SignatureResponse struct {
	Signature string `json:"signature"`
}

func GetSignature(msg []byte) (*warp.Message, error) {
	params := []map[string]any{
		{
			"unsigned_message": hexutil.Encode(msg),
		},
	}
	request := RPCRequest{
		JSONRPC: "2.0",
		Method:  "request_signature",
		Params:  params,
		ID:      1,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	url := "https://mevm.bbd.sh/v1"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 读取响应
	var response RPCResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	// 处理响应
	if response.Error != nil {
		return nil, err
	} else {
		var result SignatureResponse
		err := json.Unmarshal(response.Result, &result)
		if err != nil {
			return nil, err
		}
		decode, err := hexutil.Decode(result.Signature)
		if err != nil {
			return nil, err
		}
		signature, err := bls.SignatureFromBytes(decode)
		return signature, nil
		//println("------newMessage------", signedMsg.ID().Hex())
	}
}
