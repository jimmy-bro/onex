// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/superproj/onex/internal/fakeserver/biz/order (interfaces: OrderBiz)

// Package order is a generated GoMock package.
package order

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/superproj/onex/pkg/api/fakeserver/v1"
)

// MockOrderBiz is a mock of OrderBiz interface.
type MockOrderBiz struct {
	ctrl     *gomock.Controller
	recorder *MockOrderBizMockRecorder
}

// MockOrderBizMockRecorder is the mock recorder for MockOrderBiz.
type MockOrderBizMockRecorder struct {
	mock *MockOrderBiz
}

// NewMockOrderBiz creates a new mock instance.
func NewMockOrderBiz(ctrl *gomock.Controller) *MockOrderBiz {
	mock := &MockOrderBiz{ctrl: ctrl}
	mock.recorder = &MockOrderBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderBiz) EXPECT() *MockOrderBizMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderBiz) Create(arg0 context.Context, arg1 *v1.CreateOrderRequest) (*v1.CreateOrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*v1.CreateOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderBizMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderBiz)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockOrderBiz) Delete(arg0 context.Context, arg1 *v1.DeleteOrderRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrderBizMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrderBiz)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockOrderBiz) Get(arg0 context.Context, arg1 *v1.GetOrderRequest) (*v1.OrderReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.OrderReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrderBizMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrderBiz)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockOrderBiz) List(arg0 context.Context, arg1 *v1.ListOrderRequest) (*v1.ListOrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ListOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrderBizMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrderBiz)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockOrderBiz) Update(arg0 context.Context, arg1 *v1.UpdateOrderRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOrderBizMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderBiz)(nil).Update), arg0, arg1)
}
