// Code generated by MockGen. DO NOT EDIT.
// Source: grogu/signaller/expected_types.go
//
// Generated by this command:
//
//	mockgen -source=grogu/signaller/expected_types.go -package testutil -destination grogu/signaller/testutil/expected_types_mock.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	proto "github.com/bandprotocol/bothan/bothan-api/client/go-client/proto/bothan/v1"
	types "github.com/bandprotocol/chain/v3/x/feeds/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "go.uber.org/mock/gomock"
)

// MockBothanClient is a mock of BothanClient interface.
type MockBothanClient struct {
	ctrl     *gomock.Controller
	recorder *MockBothanClientMockRecorder
	isgomock struct{}
}

// MockBothanClientMockRecorder is the mock recorder for MockBothanClient.
type MockBothanClientMockRecorder struct {
	mock *MockBothanClient
}

// NewMockBothanClient creates a new mock instance.
func NewMockBothanClient(ctrl *gomock.Controller) *MockBothanClient {
	mock := &MockBothanClient{ctrl: ctrl}
	mock.recorder = &MockBothanClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBothanClient) EXPECT() *MockBothanClientMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockBothanClient) GetInfo() (*proto.GetInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(*proto.GetInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockBothanClientMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockBothanClient)(nil).GetInfo))
}

// GetPrices mocks base method.
func (m *MockBothanClient) GetPrices(signalIDs []string) (*proto.GetPricesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrices", signalIDs)
	ret0, _ := ret[0].(*proto.GetPricesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrices indicates an expected call of GetPrices.
func (mr *MockBothanClientMockRecorder) GetPrices(signalIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrices", reflect.TypeOf((*MockBothanClient)(nil).GetPrices), signalIDs)
}

// PushMonitoringRecords mocks base method.
func (m *MockBothanClient) PushMonitoringRecords(uuid, txHash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushMonitoringRecords", uuid, txHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushMonitoringRecords indicates an expected call of PushMonitoringRecords.
func (mr *MockBothanClientMockRecorder) PushMonitoringRecords(uuid, txHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushMonitoringRecords", reflect.TypeOf((*MockBothanClient)(nil).PushMonitoringRecords), uuid, txHash)
}

// UpdateRegistry mocks base method.
func (m *MockBothanClient) UpdateRegistry(ipfsHash, version string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRegistry", ipfsHash, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRegistry indicates an expected call of UpdateRegistry.
func (mr *MockBothanClientMockRecorder) UpdateRegistry(ipfsHash, version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRegistry", reflect.TypeOf((*MockBothanClient)(nil).UpdateRegistry), ipfsHash, version)
}

// MockFeedQuerier is a mock of FeedQuerier interface.
type MockFeedQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockFeedQuerierMockRecorder
	isgomock struct{}
}

// MockFeedQuerierMockRecorder is the mock recorder for MockFeedQuerier.
type MockFeedQuerierMockRecorder struct {
	mock *MockFeedQuerier
}

// NewMockFeedQuerier creates a new mock instance.
func NewMockFeedQuerier(ctrl *gomock.Controller) *MockFeedQuerier {
	mock := &MockFeedQuerier{ctrl: ctrl}
	mock.recorder = &MockFeedQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeedQuerier) EXPECT() *MockFeedQuerierMockRecorder {
	return m.recorder
}

// QueryCurrentFeeds mocks base method.
func (m *MockFeedQuerier) QueryCurrentFeeds() (*types.QueryCurrentFeedsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryCurrentFeeds")
	ret0, _ := ret[0].(*types.QueryCurrentFeedsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryCurrentFeeds indicates an expected call of QueryCurrentFeeds.
func (mr *MockFeedQuerierMockRecorder) QueryCurrentFeeds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCurrentFeeds", reflect.TypeOf((*MockFeedQuerier)(nil).QueryCurrentFeeds))
}

// QueryParams mocks base method.
func (m *MockFeedQuerier) QueryParams() (*types.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParams")
	ret0, _ := ret[0].(*types.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryParams indicates an expected call of QueryParams.
func (mr *MockFeedQuerierMockRecorder) QueryParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParams", reflect.TypeOf((*MockFeedQuerier)(nil).QueryParams))
}

// QueryValidValidator mocks base method.
func (m *MockFeedQuerier) QueryValidValidator(valAddress types0.ValAddress) (*types.QueryValidValidatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryValidValidator", valAddress)
	ret0, _ := ret[0].(*types.QueryValidValidatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryValidValidator indicates an expected call of QueryValidValidator.
func (mr *MockFeedQuerierMockRecorder) QueryValidValidator(valAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryValidValidator", reflect.TypeOf((*MockFeedQuerier)(nil).QueryValidValidator), valAddress)
}

// QueryValidatorPrices mocks base method.
func (m *MockFeedQuerier) QueryValidatorPrices(valAddress types0.ValAddress) (*types.QueryValidatorPricesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryValidatorPrices", valAddress)
	ret0, _ := ret[0].(*types.QueryValidatorPricesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryValidatorPrices indicates an expected call of QueryValidatorPrices.
func (mr *MockFeedQuerierMockRecorder) QueryValidatorPrices(valAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryValidatorPrices", reflect.TypeOf((*MockFeedQuerier)(nil).QueryValidatorPrices), valAddress)
}