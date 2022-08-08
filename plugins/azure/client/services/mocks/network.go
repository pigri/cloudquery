// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: ExpressRouteCircuitsClient,ExpressRouteGatewaysClient,ExpressRoutePortsClient,InterfacesClient,PublicIPAddressesClient,RouteFiltersClient,RouteTablesClient,SecurityGroupsClient,VirtualNetworkGatewaysClient,VirtualNetworksClient,WatchersClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	gomock "github.com/golang/mock/gomock"
)

// MockExpressRouteCircuitsClient is a mock of ExpressRouteCircuitsClient interface.
type MockExpressRouteCircuitsClient struct {
	ctrl     *gomock.Controller
	recorder *MockExpressRouteCircuitsClientMockRecorder
}

// MockExpressRouteCircuitsClientMockRecorder is the mock recorder for MockExpressRouteCircuitsClient.
type MockExpressRouteCircuitsClientMockRecorder struct {
	mock *MockExpressRouteCircuitsClient
}

// NewMockExpressRouteCircuitsClient creates a new mock instance.
func NewMockExpressRouteCircuitsClient(ctrl *gomock.Controller) *MockExpressRouteCircuitsClient {
	mock := &MockExpressRouteCircuitsClient{ctrl: ctrl}
	mock.recorder = &MockExpressRouteCircuitsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpressRouteCircuitsClient) EXPECT() *MockExpressRouteCircuitsClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockExpressRouteCircuitsClient) ListAll(arg0 context.Context) (network.ExpressRouteCircuitListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.ExpressRouteCircuitListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockExpressRouteCircuitsClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockExpressRouteCircuitsClient)(nil).ListAll), arg0)
}

// MockExpressRouteGatewaysClient is a mock of ExpressRouteGatewaysClient interface.
type MockExpressRouteGatewaysClient struct {
	ctrl     *gomock.Controller
	recorder *MockExpressRouteGatewaysClientMockRecorder
}

// MockExpressRouteGatewaysClientMockRecorder is the mock recorder for MockExpressRouteGatewaysClient.
type MockExpressRouteGatewaysClientMockRecorder struct {
	mock *MockExpressRouteGatewaysClient
}

// NewMockExpressRouteGatewaysClient creates a new mock instance.
func NewMockExpressRouteGatewaysClient(ctrl *gomock.Controller) *MockExpressRouteGatewaysClient {
	mock := &MockExpressRouteGatewaysClient{ctrl: ctrl}
	mock.recorder = &MockExpressRouteGatewaysClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpressRouteGatewaysClient) EXPECT() *MockExpressRouteGatewaysClientMockRecorder {
	return m.recorder
}

// ListBySubscription mocks base method.
func (m *MockExpressRouteGatewaysClient) ListBySubscription(arg0 context.Context) (network.ExpressRouteGatewayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0)
	ret0, _ := ret[0].(network.ExpressRouteGatewayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscription indicates an expected call of ListBySubscription.
func (mr *MockExpressRouteGatewaysClientMockRecorder) ListBySubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockExpressRouteGatewaysClient)(nil).ListBySubscription), arg0)
}

// MockExpressRoutePortsClient is a mock of ExpressRoutePortsClient interface.
type MockExpressRoutePortsClient struct {
	ctrl     *gomock.Controller
	recorder *MockExpressRoutePortsClientMockRecorder
}

// MockExpressRoutePortsClientMockRecorder is the mock recorder for MockExpressRoutePortsClient.
type MockExpressRoutePortsClientMockRecorder struct {
	mock *MockExpressRoutePortsClient
}

// NewMockExpressRoutePortsClient creates a new mock instance.
func NewMockExpressRoutePortsClient(ctrl *gomock.Controller) *MockExpressRoutePortsClient {
	mock := &MockExpressRoutePortsClient{ctrl: ctrl}
	mock.recorder = &MockExpressRoutePortsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpressRoutePortsClient) EXPECT() *MockExpressRoutePortsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockExpressRoutePortsClient) List(arg0 context.Context) (network.ExpressRoutePortListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(network.ExpressRoutePortListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockExpressRoutePortsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockExpressRoutePortsClient)(nil).List), arg0)
}

// MockInterfacesClient is a mock of InterfacesClient interface.
type MockInterfacesClient struct {
	ctrl     *gomock.Controller
	recorder *MockInterfacesClientMockRecorder
}

// MockInterfacesClientMockRecorder is the mock recorder for MockInterfacesClient.
type MockInterfacesClientMockRecorder struct {
	mock *MockInterfacesClient
}

// NewMockInterfacesClient creates a new mock instance.
func NewMockInterfacesClient(ctrl *gomock.Controller) *MockInterfacesClient {
	mock := &MockInterfacesClient{ctrl: ctrl}
	mock.recorder = &MockInterfacesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfacesClient) EXPECT() *MockInterfacesClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockInterfacesClient) ListAll(arg0 context.Context) (network.InterfaceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.InterfaceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockInterfacesClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockInterfacesClient)(nil).ListAll), arg0)
}

// MockPublicIPAddressesClient is a mock of PublicIPAddressesClient interface.
type MockPublicIPAddressesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPublicIPAddressesClientMockRecorder
}

// MockPublicIPAddressesClientMockRecorder is the mock recorder for MockPublicIPAddressesClient.
type MockPublicIPAddressesClientMockRecorder struct {
	mock *MockPublicIPAddressesClient
}

// NewMockPublicIPAddressesClient creates a new mock instance.
func NewMockPublicIPAddressesClient(ctrl *gomock.Controller) *MockPublicIPAddressesClient {
	mock := &MockPublicIPAddressesClient{ctrl: ctrl}
	mock.recorder = &MockPublicIPAddressesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublicIPAddressesClient) EXPECT() *MockPublicIPAddressesClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockPublicIPAddressesClient) ListAll(arg0 context.Context) (network.PublicIPAddressListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.PublicIPAddressListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockPublicIPAddressesClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).ListAll), arg0)
}

// MockRouteFiltersClient is a mock of RouteFiltersClient interface.
type MockRouteFiltersClient struct {
	ctrl     *gomock.Controller
	recorder *MockRouteFiltersClientMockRecorder
}

// MockRouteFiltersClientMockRecorder is the mock recorder for MockRouteFiltersClient.
type MockRouteFiltersClientMockRecorder struct {
	mock *MockRouteFiltersClient
}

// NewMockRouteFiltersClient creates a new mock instance.
func NewMockRouteFiltersClient(ctrl *gomock.Controller) *MockRouteFiltersClient {
	mock := &MockRouteFiltersClient{ctrl: ctrl}
	mock.recorder = &MockRouteFiltersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteFiltersClient) EXPECT() *MockRouteFiltersClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockRouteFiltersClient) List(arg0 context.Context) (network.RouteFilterListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(network.RouteFilterListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRouteFiltersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRouteFiltersClient)(nil).List), arg0)
}

// MockRouteTablesClient is a mock of RouteTablesClient interface.
type MockRouteTablesClient struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTablesClientMockRecorder
}

// MockRouteTablesClientMockRecorder is the mock recorder for MockRouteTablesClient.
type MockRouteTablesClientMockRecorder struct {
	mock *MockRouteTablesClient
}

// NewMockRouteTablesClient creates a new mock instance.
func NewMockRouteTablesClient(ctrl *gomock.Controller) *MockRouteTablesClient {
	mock := &MockRouteTablesClient{ctrl: ctrl}
	mock.recorder = &MockRouteTablesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteTablesClient) EXPECT() *MockRouteTablesClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockRouteTablesClient) ListAll(arg0 context.Context) (network.RouteTableListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.RouteTableListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockRouteTablesClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockRouteTablesClient)(nil).ListAll), arg0)
}

// MockSecurityGroupsClient is a mock of SecurityGroupsClient interface.
type MockSecurityGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityGroupsClientMockRecorder
}

// MockSecurityGroupsClientMockRecorder is the mock recorder for MockSecurityGroupsClient.
type MockSecurityGroupsClientMockRecorder struct {
	mock *MockSecurityGroupsClient
}

// NewMockSecurityGroupsClient creates a new mock instance.
func NewMockSecurityGroupsClient(ctrl *gomock.Controller) *MockSecurityGroupsClient {
	mock := &MockSecurityGroupsClient{ctrl: ctrl}
	mock.recorder = &MockSecurityGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecurityGroupsClient) EXPECT() *MockSecurityGroupsClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockSecurityGroupsClient) ListAll(arg0 context.Context) (network.SecurityGroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.SecurityGroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockSecurityGroupsClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockSecurityGroupsClient)(nil).ListAll), arg0)
}

// MockVirtualNetworkGatewaysClient is a mock of VirtualNetworkGatewaysClient interface.
type MockVirtualNetworkGatewaysClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworkGatewaysClientMockRecorder
}

// MockVirtualNetworkGatewaysClientMockRecorder is the mock recorder for MockVirtualNetworkGatewaysClient.
type MockVirtualNetworkGatewaysClientMockRecorder struct {
	mock *MockVirtualNetworkGatewaysClient
}

// NewMockVirtualNetworkGatewaysClient creates a new mock instance.
func NewMockVirtualNetworkGatewaysClient(ctrl *gomock.Controller) *MockVirtualNetworkGatewaysClient {
	mock := &MockVirtualNetworkGatewaysClient{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworkGatewaysClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualNetworkGatewaysClient) EXPECT() *MockVirtualNetworkGatewaysClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockVirtualNetworkGatewaysClient) List(arg0 context.Context, arg1 string) (network.VirtualNetworkGatewayListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(network.VirtualNetworkGatewayListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVirtualNetworkGatewaysClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualNetworkGatewaysClient)(nil).List), arg0, arg1)
}

// ListConnections mocks base method.
func (m *MockVirtualNetworkGatewaysClient) ListConnections(arg0 context.Context, arg1, arg2 string) (network.VirtualNetworkGatewayListConnectionsResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConnections", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.VirtualNetworkGatewayListConnectionsResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnections indicates an expected call of ListConnections.
func (mr *MockVirtualNetworkGatewaysClientMockRecorder) ListConnections(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnections", reflect.TypeOf((*MockVirtualNetworkGatewaysClient)(nil).ListConnections), arg0, arg1, arg2)
}

// MockVirtualNetworksClient is a mock of VirtualNetworksClient interface.
type MockVirtualNetworksClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworksClientMockRecorder
}

// MockVirtualNetworksClientMockRecorder is the mock recorder for MockVirtualNetworksClient.
type MockVirtualNetworksClientMockRecorder struct {
	mock *MockVirtualNetworksClient
}

// NewMockVirtualNetworksClient creates a new mock instance.
func NewMockVirtualNetworksClient(ctrl *gomock.Controller) *MockVirtualNetworksClient {
	mock := &MockVirtualNetworksClient{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualNetworksClient) EXPECT() *MockVirtualNetworksClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockVirtualNetworksClient) ListAll(arg0 context.Context) (network.VirtualNetworkListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.VirtualNetworkListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockVirtualNetworksClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockVirtualNetworksClient)(nil).ListAll), arg0)
}

// MockWatchersClient is a mock of WatchersClient interface.
type MockWatchersClient struct {
	ctrl     *gomock.Controller
	recorder *MockWatchersClientMockRecorder
}

// MockWatchersClientMockRecorder is the mock recorder for MockWatchersClient.
type MockWatchersClientMockRecorder struct {
	mock *MockWatchersClient
}

// NewMockWatchersClient creates a new mock instance.
func NewMockWatchersClient(ctrl *gomock.Controller) *MockWatchersClient {
	mock := &MockWatchersClient{ctrl: ctrl}
	mock.recorder = &MockWatchersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchersClient) EXPECT() *MockWatchersClientMockRecorder {
	return m.recorder
}

// GetFlowLogStatus mocks base method.
func (m *MockWatchersClient) GetFlowLogStatus(arg0 context.Context, arg1, arg2 string, arg3 network.FlowLogStatusParameters) (network.WatchersGetFlowLogStatusFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowLogStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.WatchersGetFlowLogStatusFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlowLogStatus indicates an expected call of GetFlowLogStatus.
func (mr *MockWatchersClientMockRecorder) GetFlowLogStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowLogStatus", reflect.TypeOf((*MockWatchersClient)(nil).GetFlowLogStatus), arg0, arg1, arg2, arg3)
}

// ListAll mocks base method.
func (m *MockWatchersClient) ListAll(arg0 context.Context) (network.WatcherListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.WatcherListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockWatchersClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockWatchersClient)(nil).ListAll), arg0)
}
