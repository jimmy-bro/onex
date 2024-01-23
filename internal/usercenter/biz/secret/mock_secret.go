// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/superproj/onex/internal/usercenter/biz/secret (interfaces: SecretBiz)

// Package secret is a generated GoMock package.
package secret

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/superproj/onex/pkg/api/usercenter/v1"
)

// MockSecretBiz is a mock of SecretBiz interface.
type MockSecretBiz struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBizMockRecorder
}

// MockSecretBizMockRecorder is the mock recorder for MockSecretBiz.
type MockSecretBizMockRecorder struct {
	mock *MockSecretBiz
}

// NewMockSecretBiz creates a new mock instance.
func NewMockSecretBiz(ctrl *gomock.Controller) *MockSecretBiz {
	mock := &MockSecretBiz{ctrl: ctrl}
	mock.recorder = &MockSecretBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBiz) EXPECT() *MockSecretBizMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSecretBiz) Create(arg0 context.Context, arg1 *v1.CreateSecretRequest) (*v1.SecretReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*v1.SecretReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSecretBizMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSecretBiz)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockSecretBiz) Delete(arg0 context.Context, arg1 *v1.DeleteSecretRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSecretBizMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSecretBiz)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockSecretBiz) Get(arg0 context.Context, arg1 *v1.GetSecretRequest) (*v1.SecretReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.SecretReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSecretBizMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecretBiz)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockSecretBiz) List(arg0 context.Context, arg1 *v1.ListSecretRequest) (*v1.ListSecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ListSecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSecretBizMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretBiz)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockSecretBiz) Update(arg0 context.Context, arg1 *v1.UpdateSecretRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSecretBizMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSecretBiz)(nil).Update), arg0, arg1)
}
