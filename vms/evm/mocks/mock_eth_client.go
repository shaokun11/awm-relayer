// Code generated by MockGen. DO NOT EDIT.
// Source: destination_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	types "github.com/ava-labs/subnet-evm/core/types"
	interfaces "github.com/ava-labs/subnet-evm/interfaces"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AcceptedCallContract mocks base method.
func (m *MockClient) AcceptedCallContract(arg0 context.Context, arg1 interfaces.CallMsg) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedCallContract", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptedCallContract indicates an expected call of AcceptedCallContract.
func (mr *MockClientMockRecorder) AcceptedCallContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedCallContract", reflect.TypeOf((*MockClient)(nil).AcceptedCallContract), arg0, arg1)
}

// AcceptedCodeAt mocks base method.
func (m *MockClient) AcceptedCodeAt(arg0 context.Context, arg1 common.Address) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedCodeAt", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptedCodeAt indicates an expected call of AcceptedCodeAt.
func (mr *MockClientMockRecorder) AcceptedCodeAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedCodeAt", reflect.TypeOf((*MockClient)(nil).AcceptedCodeAt), arg0, arg1)
}

// AcceptedNonceAt mocks base method.
func (m *MockClient) AcceptedNonceAt(arg0 context.Context, arg1 common.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedNonceAt", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptedNonceAt indicates an expected call of AcceptedNonceAt.
func (mr *MockClientMockRecorder) AcceptedNonceAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedNonceAt", reflect.TypeOf((*MockClient)(nil).AcceptedNonceAt), arg0, arg1)
}

// AssetBalanceAt mocks base method.
func (m *MockClient) AssetBalanceAt(arg0 context.Context, arg1 common.Address, arg2 ids.ID, arg3 *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssetBalanceAt", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssetBalanceAt indicates an expected call of AssetBalanceAt.
func (mr *MockClientMockRecorder) AssetBalanceAt(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssetBalanceAt", reflect.TypeOf((*MockClient)(nil).AssetBalanceAt), arg0, arg1, arg2, arg3)
}

// BalanceAt mocks base method.
func (m *MockClient) BalanceAt(arg0 context.Context, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceAt indicates an expected call of BalanceAt.
func (mr *MockClientMockRecorder) BalanceAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceAt", reflect.TypeOf((*MockClient)(nil).BalanceAt), arg0, arg1, arg2)
}

// BlockByHash mocks base method.
func (m *MockClient) BlockByHash(arg0 context.Context, arg1 common.Hash) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", arg0, arg1)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockClientMockRecorder) BlockByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockClient)(nil).BlockByHash), arg0, arg1)
}

// BlockByNumber mocks base method.
func (m *MockClient) BlockByNumber(arg0 context.Context, arg1 *big.Int) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockClientMockRecorder) BlockByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockClient)(nil).BlockByNumber), arg0, arg1)
}

// BlockNumber mocks base method.
func (m *MockClient) BlockNumber(arg0 context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockNumber indicates an expected call of BlockNumber.
func (mr *MockClientMockRecorder) BlockNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockClient)(nil).BlockNumber), arg0)
}

// CallContract mocks base method.
func (m *MockClient) CallContract(arg0 context.Context, arg1 interfaces.CallMsg, arg2 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockClientMockRecorder) CallContract(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockClient)(nil).CallContract), arg0, arg1, arg2)
}

// CallContractAtHash mocks base method.
func (m *MockClient) CallContractAtHash(ctx context.Context, msg interfaces.CallMsg, blockHash common.Hash) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContractAtHash", ctx, msg, blockHash)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContractAtHash indicates an expected call of CallContractAtHash.
func (mr *MockClientMockRecorder) CallContractAtHash(ctx, msg, blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContractAtHash", reflect.TypeOf((*MockClient)(nil).CallContractAtHash), ctx, msg, blockHash)
}

// ChainID mocks base method.
func (m *MockClient) ChainID(arg0 context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainID indicates an expected call of ChainID.
func (mr *MockClientMockRecorder) ChainID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockClient)(nil).ChainID), arg0)
}

// Close mocks base method.
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// CodeAt mocks base method.
func (m *MockClient) CodeAt(arg0 context.Context, arg1 common.Address, arg2 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockClientMockRecorder) CodeAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockClient)(nil).CodeAt), arg0, arg1, arg2)
}

// EstimateBaseFee mocks base method.
func (m *MockClient) EstimateBaseFee(arg0 context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateBaseFee", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateBaseFee indicates an expected call of EstimateBaseFee.
func (mr *MockClientMockRecorder) EstimateBaseFee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateBaseFee", reflect.TypeOf((*MockClient)(nil).EstimateBaseFee), arg0)
}

// EstimateGas mocks base method.
func (m *MockClient) EstimateGas(arg0 context.Context, arg1 interfaces.CallMsg) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas.
func (mr *MockClientMockRecorder) EstimateGas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockClient)(nil).EstimateGas), arg0, arg1)
}

// FeeHistory mocks base method.
func (m *MockClient) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*interfaces.FeeHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeeHistory", ctx, blockCount, lastBlock, rewardPercentiles)
	ret0, _ := ret[0].(*interfaces.FeeHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FeeHistory indicates an expected call of FeeHistory.
func (mr *MockClientMockRecorder) FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeeHistory", reflect.TypeOf((*MockClient)(nil).FeeHistory), ctx, blockCount, lastBlock, rewardPercentiles)
}

// FilterLogs mocks base method.
func (m *MockClient) FilterLogs(arg0 context.Context, arg1 interfaces.FilterQuery) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterLogs", arg0, arg1)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterLogs indicates an expected call of FilterLogs.
func (mr *MockClientMockRecorder) FilterLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterLogs", reflect.TypeOf((*MockClient)(nil).FilterLogs), arg0, arg1)
}

// HeaderByHash mocks base method.
func (m *MockClient) HeaderByHash(arg0 context.Context, arg1 common.Hash) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByHash", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByHash indicates an expected call of HeaderByHash.
func (mr *MockClientMockRecorder) HeaderByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByHash", reflect.TypeOf((*MockClient)(nil).HeaderByHash), arg0, arg1)
}

// HeaderByNumber mocks base method.
func (m *MockClient) HeaderByNumber(arg0 context.Context, arg1 *big.Int) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber.
func (mr *MockClientMockRecorder) HeaderByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockClient)(nil).HeaderByNumber), arg0, arg1)
}

// NetworkID mocks base method.
func (m *MockClient) NetworkID(arg0 context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkID", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkID indicates an expected call of NetworkID.
func (mr *MockClientMockRecorder) NetworkID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkID", reflect.TypeOf((*MockClient)(nil).NetworkID), arg0)
}

// NonceAt mocks base method.
func (m *MockClient) NonceAt(arg0 context.Context, arg1 common.Address, arg2 *big.Int) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NonceAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NonceAt indicates an expected call of NonceAt.
func (mr *MockClientMockRecorder) NonceAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NonceAt", reflect.TypeOf((*MockClient)(nil).NonceAt), arg0, arg1, arg2)
}

// SendTransaction mocks base method.
func (m *MockClient) SendTransaction(arg0 context.Context, arg1 *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockClientMockRecorder) SendTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockClient)(nil).SendTransaction), arg0, arg1)
}

// StorageAt mocks base method.
func (m *MockClient) StorageAt(arg0 context.Context, arg1 common.Address, arg2 common.Hash, arg3 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAt", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAt indicates an expected call of StorageAt.
func (mr *MockClientMockRecorder) StorageAt(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAt", reflect.TypeOf((*MockClient)(nil).StorageAt), arg0, arg1, arg2, arg3)
}

// SubscribeFilterLogs mocks base method.
func (m *MockClient) SubscribeFilterLogs(arg0 context.Context, arg1 interfaces.FilterQuery, arg2 chan<- types.Log) (interfaces.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeFilterLogs", arg0, arg1, arg2)
	ret0, _ := ret[0].(interfaces.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeFilterLogs indicates an expected call of SubscribeFilterLogs.
func (mr *MockClientMockRecorder) SubscribeFilterLogs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeFilterLogs", reflect.TypeOf((*MockClient)(nil).SubscribeFilterLogs), arg0, arg1, arg2)
}

// SubscribeNewAcceptedTransactions mocks base method.
func (m *MockClient) SubscribeNewAcceptedTransactions(arg0 context.Context, arg1 chan<- *common.Hash) (interfaces.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewAcceptedTransactions", arg0, arg1)
	ret0, _ := ret[0].(interfaces.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeNewAcceptedTransactions indicates an expected call of SubscribeNewAcceptedTransactions.
func (mr *MockClientMockRecorder) SubscribeNewAcceptedTransactions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewAcceptedTransactions", reflect.TypeOf((*MockClient)(nil).SubscribeNewAcceptedTransactions), arg0, arg1)
}

// SubscribeNewHead mocks base method.
func (m *MockClient) SubscribeNewHead(arg0 context.Context, arg1 chan<- *types.Header) (interfaces.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewHead", arg0, arg1)
	ret0, _ := ret[0].(interfaces.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeNewHead indicates an expected call of SubscribeNewHead.
func (mr *MockClientMockRecorder) SubscribeNewHead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewHead", reflect.TypeOf((*MockClient)(nil).SubscribeNewHead), arg0, arg1)
}

// SubscribeNewPendingTransactions mocks base method.
func (m *MockClient) SubscribeNewPendingTransactions(arg0 context.Context, arg1 chan<- *common.Hash) (interfaces.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewPendingTransactions", arg0, arg1)
	ret0, _ := ret[0].(interfaces.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeNewPendingTransactions indicates an expected call of SubscribeNewPendingTransactions.
func (mr *MockClientMockRecorder) SubscribeNewPendingTransactions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewPendingTransactions", reflect.TypeOf((*MockClient)(nil).SubscribeNewPendingTransactions), arg0, arg1)
}

// SuggestGasPrice mocks base method.
func (m *MockClient) SuggestGasPrice(arg0 context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasPrice", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasPrice indicates an expected call of SuggestGasPrice.
func (mr *MockClientMockRecorder) SuggestGasPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasPrice", reflect.TypeOf((*MockClient)(nil).SuggestGasPrice), arg0)
}

// SuggestGasTipCap mocks base method.
func (m *MockClient) SuggestGasTipCap(arg0 context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasTipCap", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasTipCap indicates an expected call of SuggestGasTipCap.
func (mr *MockClientMockRecorder) SuggestGasTipCap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasTipCap", reflect.TypeOf((*MockClient)(nil).SuggestGasTipCap), arg0)
}

// SyncProgress mocks base method.
func (m *MockClient) SyncProgress(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncProgress", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncProgress indicates an expected call of SyncProgress.
func (mr *MockClientMockRecorder) SyncProgress(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncProgress", reflect.TypeOf((*MockClient)(nil).SyncProgress), ctx)
}

// TransactionByHash mocks base method.
func (m *MockClient) TransactionByHash(arg0 context.Context, arg1 common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", arg0, arg1)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockClientMockRecorder) TransactionByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockClient)(nil).TransactionByHash), arg0, arg1)
}

// TransactionCount mocks base method.
func (m *MockClient) TransactionCount(arg0 context.Context, arg1 common.Hash) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionCount", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionCount indicates an expected call of TransactionCount.
func (mr *MockClientMockRecorder) TransactionCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionCount", reflect.TypeOf((*MockClient)(nil).TransactionCount), arg0, arg1)
}

// TransactionInBlock mocks base method.
func (m *MockClient) TransactionInBlock(arg0 context.Context, arg1 common.Hash, arg2 uint) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionInBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionInBlock indicates an expected call of TransactionInBlock.
func (mr *MockClientMockRecorder) TransactionInBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionInBlock", reflect.TypeOf((*MockClient)(nil).TransactionInBlock), arg0, arg1, arg2)
}

// TransactionReceipt mocks base method.
func (m *MockClient) TransactionReceipt(arg0 context.Context, arg1 common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", arg0, arg1)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockClientMockRecorder) TransactionReceipt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockClient)(nil).TransactionReceipt), arg0, arg1)
}

// TransactionSender mocks base method.
func (m *MockClient) TransactionSender(arg0 context.Context, arg1 *types.Transaction, arg2 common.Hash, arg3 uint) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionSender", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionSender indicates an expected call of TransactionSender.
func (mr *MockClientMockRecorder) TransactionSender(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionSender", reflect.TypeOf((*MockClient)(nil).TransactionSender), arg0, arg1, arg2, arg3)
}