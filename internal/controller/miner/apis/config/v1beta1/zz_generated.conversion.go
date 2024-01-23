//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "k8s.io/component-base/config/v1alpha1"

	config "github.com/superproj/onex/internal/controller/miner/apis/config"
	configv1beta1 "github.com/superproj/onex/pkg/config/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*MinerControllerConfiguration)(nil), (*config.MinerControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerControllerConfiguration_To_config_MinerControllerConfiguration(a.(*MinerControllerConfiguration), b.(*config.MinerControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MinerControllerConfiguration)(nil), (*MinerControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MinerControllerConfiguration_To_v1beta1_MinerControllerConfiguration(a.(*config.MinerControllerConfiguration), b.(*MinerControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerProfile)(nil), (*config.MinerProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerProfile_To_config_MinerProfile(a.(*MinerProfile), b.(*config.MinerProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MinerProfile)(nil), (*MinerProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MinerProfile_To_v1beta1_MinerProfile(a.(*config.MinerProfile), b.(*MinerProfile), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_MinerControllerConfiguration_To_config_MinerControllerConfiguration(in *MinerControllerConfiguration, out *config.MinerControllerConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.Parallelism = in.Parallelism
	out.DryRun = in.DryRun
	out.ProviderKubeconfig = in.ProviderKubeconfig
	out.InCluster = in.InCluster
	out.SyncPeriod = in.SyncPeriod
	out.WatchFilterValue = in.WatchFilterValue
	if err := v1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.MetricsBindAddress = in.MetricsBindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.Types = *(*map[string]config.MinerProfile)(unsafe.Pointer(&in.Types))
	if err := configv1beta1.Convert_v1beta1_RedisConfiguration_To_config_RedisConfiguration(&in.Redis, &out.Redis, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_MinerControllerConfiguration_To_config_MinerControllerConfiguration is an autogenerated conversion function.
func Convert_v1beta1_MinerControllerConfiguration_To_config_MinerControllerConfiguration(in *MinerControllerConfiguration, out *config.MinerControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerControllerConfiguration_To_config_MinerControllerConfiguration(in, out, s)
}

func autoConvert_config_MinerControllerConfiguration_To_v1beta1_MinerControllerConfiguration(in *config.MinerControllerConfiguration, out *MinerControllerConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.Parallelism = in.Parallelism
	out.DryRun = in.DryRun
	out.ProviderKubeconfig = in.ProviderKubeconfig
	out.InCluster = in.InCluster
	out.SyncPeriod = in.SyncPeriod
	out.WatchFilterValue = in.WatchFilterValue
	if err := v1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.MetricsBindAddress = in.MetricsBindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.Types = *(*map[string]MinerProfile)(unsafe.Pointer(&in.Types))
	if err := configv1beta1.Convert_config_RedisConfiguration_To_v1beta1_RedisConfiguration(&in.Redis, &out.Redis, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_MinerControllerConfiguration_To_v1beta1_MinerControllerConfiguration is an autogenerated conversion function.
func Convert_config_MinerControllerConfiguration_To_v1beta1_MinerControllerConfiguration(in *config.MinerControllerConfiguration, out *MinerControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_MinerControllerConfiguration_To_v1beta1_MinerControllerConfiguration(in, out, s)
}

func autoConvert_v1beta1_MinerProfile_To_config_MinerProfile(in *MinerProfile, out *config.MinerProfile, s conversion.Scope) error {
	out.CPU = in.CPU
	out.Memory = in.Memory
	out.MiningDifficulty = in.MiningDifficulty
	return nil
}

// Convert_v1beta1_MinerProfile_To_config_MinerProfile is an autogenerated conversion function.
func Convert_v1beta1_MinerProfile_To_config_MinerProfile(in *MinerProfile, out *config.MinerProfile, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerProfile_To_config_MinerProfile(in, out, s)
}

func autoConvert_config_MinerProfile_To_v1beta1_MinerProfile(in *config.MinerProfile, out *MinerProfile, s conversion.Scope) error {
	out.CPU = in.CPU
	out.Memory = in.Memory
	out.MiningDifficulty = in.MiningDifficulty
	return nil
}

// Convert_config_MinerProfile_To_v1beta1_MinerProfile is an autogenerated conversion function.
func Convert_config_MinerProfile_To_v1beta1_MinerProfile(in *config.MinerProfile, out *MinerProfile, s conversion.Scope) error {
	return autoConvert_config_MinerProfile_To_v1beta1_MinerProfile(in, out, s)
}
