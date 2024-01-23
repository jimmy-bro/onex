//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLConfiguration) DeepCopyInto(out *MySQLConfiguration) {
	*out = *in
	out.MaxConnectionLifeTime = in.MaxConnectionLifeTime
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLConfiguration.
func (in *MySQLConfiguration) DeepCopy() *MySQLConfiguration {
	if in == nil {
		return nil
	}
	out := new(MySQLConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConfiguration) DeepCopyInto(out *RedisConfiguration) {
	*out = *in
	out.Timeout = in.Timeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConfiguration.
func (in *RedisConfiguration) DeepCopy() *RedisConfiguration {
	if in == nil {
		return nil
	}
	out := new(RedisConfiguration)
	in.DeepCopyInto(out)
	return out
}
