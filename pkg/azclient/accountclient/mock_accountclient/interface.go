// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/cloud-provider-azure/pkg/azclient/accountclient (interfaces: Interface)

// Package mock_accountclient is a generated GoMock package.
package mock_accountclient

import (
	context "context"
	reflect "reflect"

	armstorage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockInterface) Create(arg0 context.Context, arg1, arg2 string, arg3 *armstorage.AccountCreateParameters) (*armstorage.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*armstorage.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockInterfaceMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInterface)(nil).Create), arg0, arg1, arg2, arg3)
}

// Delete mocks base method.
func (m *MockInterface) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), arg0, arg1, arg2)
}

// GetProperties mocks base method.
func (m *MockInterface) GetProperties(arg0 context.Context, arg1, arg2 string, arg3 *armstorage.AccountsClientGetPropertiesOptions) (*armstorage.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperties", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*armstorage.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProperties indicates an expected call of GetProperties.
func (mr *MockInterfaceMockRecorder) GetProperties(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperties", reflect.TypeOf((*MockInterface)(nil).GetProperties), arg0, arg1, arg2, arg3)
}

// List mocks base method.
func (m *MockInterface) List(arg0 context.Context, arg1 string) ([]*armstorage.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*armstorage.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), arg0, arg1)
}

// ListKeys mocks base method.
func (m *MockInterface) ListKeys(arg0 context.Context, arg1, arg2 string) ([]*armstorage.AccountKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*armstorage.AccountKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockInterfaceMockRecorder) ListKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockInterface)(nil).ListKeys), arg0, arg1, arg2)
}
