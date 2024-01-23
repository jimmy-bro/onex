// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 页面未找到错误，请求的页面不存在
func IsPageNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PageNotFound.String() && e.Code == 404
}

// 页面未找到错误，请求的页面不存在
func ErrorPageNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_PageNotFound.String(), fmt.Sprintf(format, args...))
}
