// Code generated by MockGen. DO NOT EDIT.
// Source: destination_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	warp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockDestinationClient is a mock of DestinationClient interface.
type MockDestinationClient struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationClientMockRecorder
}

// MockDestinationClientMockRecorder is the mock recorder for MockDestinationClient.
type MockDestinationClientMockRecorder struct {
	mock *MockDestinationClient
}

// NewMockDestinationClient creates a new mock instance.
func NewMockDestinationClient(ctrl *gomock.Controller) *MockDestinationClient {
	mock := &MockDestinationClient{ctrl: ctrl}
	mock.recorder = &MockDestinationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationClient) EXPECT() *MockDestinationClientMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockDestinationClient) Client() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockDestinationClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockDestinationClient)(nil).Client))
}

// DestinationChainID mocks base method.
func (m *MockDestinationClient) DestinationChainID() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestinationChainID")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// DestinationChainID indicates an expected call of DestinationChainID.
func (mr *MockDestinationClientMockRecorder) DestinationChainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestinationChainID", reflect.TypeOf((*MockDestinationClient)(nil).DestinationChainID))
}

// SendTx mocks base method.
func (m *MockDestinationClient) SendTx(signedMessage *warp.Message, toAddress string, gasLimit uint64, callData []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTx", signedMessage, toAddress, gasLimit, callData)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTx indicates an expected call of SendTx.
func (mr *MockDestinationClientMockRecorder) SendTx(signedMessage, toAddress, gasLimit, callData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTx", reflect.TypeOf((*MockDestinationClient)(nil).SendTx), signedMessage, toAddress, gasLimit, callData)
}

// SenderAddress mocks base method.
func (m *MockDestinationClient) SenderAddress() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SenderAddress")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// SenderAddress indicates an expected call of SenderAddress.
func (mr *MockDestinationClientMockRecorder) SenderAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SenderAddress", reflect.TypeOf((*MockDestinationClient)(nil).SenderAddress))
}
