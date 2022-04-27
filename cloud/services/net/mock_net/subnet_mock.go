// Code generated by MockGen. DO NOT EDIT.
// Source: ./subnet.go

// Package mock_net is a generated GoMock package.
package mock_net

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/outscale-dev/cluster-api-provider-outscale.git/api/v1beta1"
	osc "github.com/outscale/osc-sdk-go/v2"
)

// MockOscSubnetInterface is a mock of OscSubnetInterface interface.
type MockOscSubnetInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOscSubnetInterfaceMockRecorder
}

// MockOscSubnetInterfaceMockRecorder is the mock recorder for MockOscSubnetInterface.
type MockOscSubnetInterfaceMockRecorder struct {
	mock *MockOscSubnetInterface
}

// NewMockOscSubnetInterface creates a new mock instance.
func NewMockOscSubnetInterface(ctrl *gomock.Controller) *MockOscSubnetInterface {
	mock := &MockOscSubnetInterface{ctrl: ctrl}
	mock.recorder = &MockOscSubnetInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOscSubnetInterface) EXPECT() *MockOscSubnetInterfaceMockRecorder {
	return m.recorder
}

// CreateSubnet mocks base method.
func (m *MockOscSubnetInterface) CreateSubnet(spec *v1beta1.OscSubnet, netId, subnetName string) (*osc.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubnet", spec, netId, subnetName)
	ret0, _ := ret[0].(*osc.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubnet indicates an expected call of CreateSubnet.
func (mr *MockOscSubnetInterfaceMockRecorder) CreateSubnet(spec, netId, subnetName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubnet", reflect.TypeOf((*MockOscSubnetInterface)(nil).CreateSubnet), spec, netId, subnetName)
}

// DeleteSubnet mocks base method.
func (m *MockOscSubnetInterface) DeleteSubnet(subnetId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubnet", subnetId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubnet indicates an expected call of DeleteSubnet.
func (mr *MockOscSubnetInterfaceMockRecorder) DeleteSubnet(subnetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubnet", reflect.TypeOf((*MockOscSubnetInterface)(nil).DeleteSubnet), subnetId)
}

// GetSubnet mocks base method.
func (m *MockOscSubnetInterface) GetSubnet(subnetId string) (*osc.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnet", subnetId)
	ret0, _ := ret[0].(*osc.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnet indicates an expected call of GetSubnet.
func (mr *MockOscSubnetInterfaceMockRecorder) GetSubnet(subnetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnet", reflect.TypeOf((*MockOscSubnetInterface)(nil).GetSubnet), subnetId)
}

// GetSubnetIdsFromNetIds mocks base method.
func (m *MockOscSubnetInterface) GetSubnetIdsFromNetIds(netId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetIdsFromNetIds", netId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetIdsFromNetIds indicates an expected call of GetSubnetIdsFromNetIds.
func (mr *MockOscSubnetInterfaceMockRecorder) GetSubnetIdsFromNetIds(netId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetIdsFromNetIds", reflect.TypeOf((*MockOscSubnetInterface)(nil).GetSubnetIdsFromNetIds), netId)
}