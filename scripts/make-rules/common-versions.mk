# ==============================================================================
# Versions used by all Makefiles
#

KUSTOMIZE_VERSION ?= v3.8.7
CONTROLLER_TOOLS_VERSION ?= v0.9.2
CONTROLLER_TOOLS_VERSION ?= v0.8.0
HADOLINT_VER ?= v2.10.0
CHART_VERSION ?= 1.0.0
KIND_VERSION ?= v0.19.0
HELM_VERSION ?= v3.8.1
GO_TESTS_SUM_VERSION ?=v1.6.4
GOLANGCI_LINT_VERSION := v1.55.2
YQ_VERSION ?= v4.25.2
GRPC_GATEWAY_VERSION ?= $(call get_go_version,github.com/grpc-ecosystem/grpc-gateway/v2)
KRATOS_VERSION ?= $(call get_go_version,github.com/go-kratos/kratos/v2)
PROTOC_GEN_VALIDATE_VERSION ?= $(call get_go_version,github.com/envoyproxy/protoc-gen-validate)
PROTOC_GEN_GO_VERSION ?= $(call get_go_version,google.golang.org/protobuf)
PROTOC_GEN_GO_GRPC_VERSION ?= $(call get_go_version,google.golang.org/grpc)
CODE_GENERATOR_VERSION ?= $(call get_go_version,k8s.io/code-generator)
WIRE_VERSION ?= $(call get_go_version,github.com/google/wire)
MOCKGEN_VERSION ?= $(call get_go_version,github.com/golang/mock)
