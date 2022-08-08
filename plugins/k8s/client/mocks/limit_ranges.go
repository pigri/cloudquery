// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-k8s/client (interfaces: LimitRangesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockLimitRangesClient is a mock of LimitRangesClient interface.
type MockLimitRangesClient struct {
	ctrl     *gomock.Controller
	recorder *MockLimitRangesClientMockRecorder
}

// MockLimitRangesClientMockRecorder is the mock recorder for MockLimitRangesClient.
type MockLimitRangesClientMockRecorder struct {
	mock *MockLimitRangesClient
}

// NewMockLimitRangesClient creates a new mock instance.
func NewMockLimitRangesClient(ctrl *gomock.Controller) *MockLimitRangesClient {
	mock := &MockLimitRangesClient{ctrl: ctrl}
	mock.recorder = &MockLimitRangesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLimitRangesClient) EXPECT() *MockLimitRangesClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockLimitRangesClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.LimitRangeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.LimitRangeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLimitRangesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLimitRangesClient)(nil).List), arg0, arg1)
}
