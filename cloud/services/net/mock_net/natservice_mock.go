// Code generated by MockGen. DO NOT EDIT.
// Source: ./natservice.go

// Package mock_net is a generated GoMock package.
package mock_net

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "github.com/outscale/osc-sdk-go/v2"
)

// MockOscNatServiceInterface is a mock of OscNatServiceInterface interface.
type MockOscNatServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOscNatServiceInterfaceMockRecorder
}

// MockOscNatServiceInterfaceMockRecorder is the mock recorder for MockOscNatServiceInterface.
type MockOscNatServiceInterfaceMockRecorder struct {
	mock *MockOscNatServiceInterface
}

// NewMockOscNatServiceInterface creates a new mock instance.
func NewMockOscNatServiceInterface(ctrl *gomock.Controller) *MockOscNatServiceInterface {
	mock := &MockOscNatServiceInterface{ctrl: ctrl}
	mock.recorder = &MockOscNatServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOscNatServiceInterface) EXPECT() *MockOscNatServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateNatService mocks base method.
func (m *MockOscNatServiceInterface) CreateNatService(publicIpId, subnetId, natServiceName string) (*osc.NatService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNatService", publicIpId, subnetId, natServiceName)
	ret0, _ := ret[0].(*osc.NatService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNatService indicates an expected call of CreateNatService.
func (mr *MockOscNatServiceInterfaceMockRecorder) CreateNatService(publicIpId, subnetId, natServiceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNatService", reflect.TypeOf((*MockOscNatServiceInterface)(nil).CreateNatService), publicIpId, subnetId, natServiceName)
}

// DeleteNatService mocks base method.
func (m *MockOscNatServiceInterface) DeleteNatService(natServiceId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNatService", natServiceId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNatService indicates an expected call of DeleteNatService.
func (mr *MockOscNatServiceInterfaceMockRecorder) DeleteNatService(natServiceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNatService", reflect.TypeOf((*MockOscNatServiceInterface)(nil).DeleteNatService), natServiceId)
}

// GetNatService mocks base method.
func (m *MockOscNatServiceInterface) GetNatService(natServiceId string) (*osc.NatService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNatService", natServiceId)
	ret0, _ := ret[0].(*osc.NatService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNatService indicates an expected call of GetNatService.
func (mr *MockOscNatServiceInterfaceMockRecorder) GetNatService(natServiceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNatService", reflect.TypeOf((*MockOscNatServiceInterface)(nil).GetNatService), natServiceId)
}
