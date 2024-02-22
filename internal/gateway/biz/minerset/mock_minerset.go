// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/superproj/onex/internal/gateway/biz/minerset (interfaces: MinerSetBiz)

// Package minerset is a generated GoMock package.
package minerset

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/superproj/onex/pkg/api/gateway/v1"
	v1beta1 "github.com/superproj/onex/pkg/apis/apps/v1beta1"
)

// MockMinerSetBiz is a mock of MinerSetBiz interface.
type MockMinerSetBiz struct {
	ctrl     *gomock.Controller
	recorder *MockMinerSetBizMockRecorder
}

// MockMinerSetBizMockRecorder is the mock recorder for MockMinerSetBiz.
type MockMinerSetBizMockRecorder struct {
	mock *MockMinerSetBiz
}

// NewMockMinerSetBiz creates a new mock instance.
func NewMockMinerSetBiz(ctrl *gomock.Controller) *MockMinerSetBiz {
	mock := &MockMinerSetBiz{ctrl: ctrl}
	mock.recorder = &MockMinerSetBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMinerSetBiz) EXPECT() *MockMinerSetBizMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMinerSetBiz) Create(arg0 context.Context, arg1 string, arg2 *v1beta1.MinerSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMinerSetBizMockRecorder) Create(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMinerSetBiz)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockMinerSetBiz) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMinerSetBizMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMinerSetBiz)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockMinerSetBiz) Get(arg0 context.Context, arg1, arg2 string) (*v1beta1.MinerSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.MinerSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMinerSetBizMockRecorder) Get(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMinerSetBiz)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockMinerSetBiz) List(arg0 context.Context, arg1 string, arg2 *v1.ListMinerSetRequest) (*v1.ListMinerSetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ListMinerSetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMinerSetBizMockRecorder) List(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMinerSetBiz)(nil).List), arg0, arg1, arg2)
}

// Scale mocks base method.
func (m *MockMinerSetBiz) Scale(arg0 context.Context, arg1, arg2 string, arg3 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scale", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scale indicates an expected call of Scale.
func (mr *MockMinerSetBizMockRecorder) Scale(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockMinerSetBiz)(nil).Scale), arg0, arg1, arg2, arg3)
}

// Update mocks base method.
func (m *MockMinerSetBiz) Update(arg0 context.Context, arg1 string, arg2 *v1beta1.MinerSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMinerSetBizMockRecorder) Update(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMinerSetBiz)(nil).Update), arg0, arg1, arg2)
}
