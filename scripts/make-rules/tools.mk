# ==============================================================================
#  Makefile helper functions for tools
#
# Rules name starting with `_` mean that it is not recommended to call directly through make command, 
# like `make _install.gotests`, you should run `make tools.install.gotests` instead.
#
# Specify tools category.
CODE_GENERATOR_TOOLS= client-gen lister-gen informer-gen defaulter-gen deepcopy-gen prerelease-lifecycle-gen conversion-gen openapi-gen
# code-generator is a makefile target not a real tool.
CI_WORKFLOW_TOOLS := code-generator golangci-lint goimports wire 
# unused tools in this project: gentool
OTHER_TOOLS := mockgen uplift git-chglog addlicense kratos kind go-apidiff gotests \
			   cfssl go-gitlint kustomize kafkactl kube-linter kubeconform kubectl \
			   helm-docs db2struct gentool air swagger license gothanks kubebuilder go-junit-report controller-gen
MANUAL_INSTALL_TOOLS := kafka

.PHONY: tools.install
tools.install: install.ci _install.other tools.print-manual-tool ## Install all tools.

.PHONY: tools.print-manual-tool
tools.print-manual-tool: 
	@echo "===========> The following tools may need to be installed manually:"
	@echo $(MANUAL_INSTALL_TOOLS) | awk 'BEGIN{RS=" "} {printf("%15s%s\n","- ",$$0)}'

.PHONY: tools.install.%
tools.install.%: ## Install a specified tool.
	@echo "===========> Installing $*"
	@$(MAKE) _install.$*

.PHONY: tools.verify.%
tools.verify.%: ## Verify a specified tool.
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: tools.verify.code-generator
tools.verify.code-generator: $(addprefix _verify.code-generator., $(CODE_GENERATOR_TOOLS)) ## Verify a specified tool.

.PHONY: _verify.code-generator.%
_verify.code-generator.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.code-generator.$*; fi

.PHONY: _install.ci
_install.ci: $(addprefix tools.install., $(CI_WORKFLOW_TOOLS)) ## Install necessary tools used by CI/CD workflow.

.PHONY: _install.other
_install.other: $(addprefix tools.install., $(OTHER_TOOLS))

.PHONY: _install.code-generator
_install.code-generator: $(addprefix tools.install.code-generator., $(CODE_GENERATOR_TOOLS)) ## Install all necessary code-generator tools.

.PHONY: _install.code-generator.%
_install.code-generator.%: ## Install specified code-generator tool.
	@$(GO) install k8s.io/code-generator/cmd/$*@$(CODE_GENERATOR_VERSION)
	#@$(GO) install github.com/colin404/code-generator/cmd/$*@$(CODE_GENERATOR_VERSION)

.PHONY: _install.swagger
_install.swagger:
	@$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: _install.golangci-lint
_install.golangci-lint: ## Install golangci-lint.
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh golangci-lint bash

.PHONY: _install.go-junit-report
_install.go-junit-report:
	@$(GO) install github.com/jstemmer/go-junit-report@latest

.PHONY: _install.wire
_install.wire: ## Install wire.
	@$(GO) install github.com/google/wire/cmd/wire@$(WIRE_VERSION)

.PHONY: _install.mockgen
_install.mockgen: ## Install mockgen.
	@$(GO) install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)

.PHONY: _install.gotests
_install.gotests: ## Install gotests.
	@$(GO) install github.com/cweill/gotests/gotests@latest

.PHONY: _install.goimports
_install.goimports: ## Install goimports.
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: _install.go-gitlint
_install.go-gitlint: ## Install go-gitlint.
	@$(GO) install github.com/marmotedu/go-gitlint/cmd/go-gitlint@latest

.PHONY: _install.gsemver
_install.gsemver: ## Install gsemver.
	@$(GO) install github.com/arnaud-deprez/gsemver@latest
	@$(SCRIPTS_DIR)/add-completion.sh gsemver bash

.PHONY: _install.uplift
_install.uplift: ## Install uplift.
	@export UPLIFT_INSTALL_DIR=$(HOME)/bin && \
		curl --retry 10 -sL https://raw.githubusercontent.com/gembaadvantage/uplift/main/scripts/install | bash -s -- '--no-sudo'
	@$(SCRIPTS_DIR)/add-completion.sh uplift bash

.PHONY: _install.git-chglog
_install.git-chglog: ## Install git-chglog tool which is used to generate CHANGELOG.
	@$(GO) install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

.PHONY: _install.cfssl
_install.cfssl: ## Install cfssl toolkit.
	@$(SCRIPTS_DIR)/install.sh onex::install::install_cfssl

.PHONY: _install.addlicense
_install.addlicense: ## Install addlicense.
	@$(GO) install github.com/superproj/addlicense@latest

.PHONY: _install.kustomize
_install.kustomize: ## Install kustomize.
	@$(GO) install sigs.k8s.io/kustomize/kustomize/v5@latest
	@$(SCRIPTS_DIR)/add-completion.sh kustomize bash

.PHONY: _install.controller-gen
_install.controller-gen: ## Install controller-gen.
	@$(GO) install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

.PHONY: _install.kind
_install.kind: ## Install kind cluster command line tool.
	@$(GO) install sigs.k8s.io/kind@$(KIND_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh kind bash

.PHONY: _install.go-apidiff
_install.go-apidiff: ## Install go-apidiff.
	@$(GO) install github.com/joelanford/go-apidiff@latest

.PHONY: _install.helm
_install.helm: ## Install helm command line tool.
	@curl --retry 3 -fsSL -o $(TMP_DIR)/get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
	@chmod 700 $(TMP_DIR)/get_helm.sh
	@USE_SUDO=false HELM_INSTALL_DIR=$(GOBIN) DESIRED_VERSION=$(HELM_VERSION) BINARY_NAME=helm $(TMP_DIR)/get_helm.sh
	@$(SCRIPTS_DIR)/add-completion.sh helm bash

.PHONY: _install.grpc
_install.grpc:
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@$(GRPC_GATEWAY_VERSION)
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@$(GRPC_GATEWAY_VERSION)
	@$(SCRIPTS_DIR)/install-protoc.sh

.PHONY: _install.kratos
_install.kratos: _install.grpc ## Install kratos toolkit, includes multiple protoc plugins.
	@$(GO) install github.com/joelanford/go-apidiff@latest
	@$(GO) install github.com/envoyproxy/protoc-gen-validate@$(PROTOC_GEN_VALIDATE_VERSION)
	@$(GO) install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@$(GO) install github.com/go-kratos/kratos/cmd/kratos/v2@$(KRATOS_VERSION)
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@$(KRATOS_VERSION)
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@$(KRATOS_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh kratos bash

.PHONY: _install.buf
_install.buf: ## Install buf command line tool.
	@$(GO) install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: _install.kafkactl
_install.kafkactl: ## Install kafkactl command line tool.
	@$(GO) install github.com/deviceinsight/kafkactl@latest
	@$(SCRIPTS_DIR)/add-completion.sh kafkactl bash

# kube-linter reference: https://docs.kubelinter.io/#/
.PHONY: _install.kube-linter
_install.kube-linter: ## Install kube-linter command line tool.
	@$(GO) install golang.stackrox.io/kube-linter/cmd/kube-linter@latest
	@$(SCRIPTS_DIR)/add-completion.sh kube-linter bash

.PHONY: _install.kubeconform
_install.kubeconform: ## Install kubeconform command line tool.
	@$(GO) install github.com/yannh/kubeconform/cmd/kubeconform@latest

.PHONY: _install.kubectl
_install.kubectl: ## Install kubectl command line tool.
	@curl --create-dirs -L -o $$HOME/bin/kubectl "https://dl.k8s.io/release/$(shell curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
	@chmod +x $$HOME/bin/kubectl
	@$(SCRIPTS_DIR)/add-completion.sh kubectl bash

.PHONY: _install.helm-docs
_install.helm-docs: ## Install helm-docs which is a tool to generating markdown documentation for helm charts.
	@$(GO) install github.com/norwoodj/helm-docs/cmd/helm-docs@latest

.PHONY: _install.gentool
_install.gentool: ## Install gentool which is a tool used to generate gorm model and query code.
	@$(GO) install gorm.io/gen/tools/gentool@latest

# db2struct --gorm --json -H 127.0.0.1 -d onex -t secret --package model --struct SecretM -u gateway -p 'proj(#)666' --target=secret.go
.PHONY: _install.db2struct
_install.db2struct: ## Install db2struct which is a tool used to converts a mysql table into a golang struct.
	@$(GO) install github.com/Shelnutt2/db2struct/cmd/db2struct@latest

.PHONY: _install.protoc-go-inject-tag
_install.protoc-go-inject-tag:
	@$(GO) install github.com/favadi/protoc-go-inject-tag@latest

.PHONY: _install.air
_install.air: ## Install air tool which is used to live reload your go apps.
	@$(GO) install github.com/cosmtrek/air@latest

.PHONY: _install.license 
_install.license : ## Install license tool which is used to generate LICENSE file as you want.
	@$(GO) install github.com/nishanths/license/v5@latest

.PHONY: _install.gothanks
_install.gothanks: ## Install gothanks tool which is used to automatically stars your go.mod github dependencies.
	@$(GO) install github.com/psampaz/gothanks@latest

.PHONY: _install.kubebuilder
_install.kubebuilder : ## Install kubebuilder tool which is used to building Kubernetes APIs using CRDs.
	# download kubebuilder and install locally.
	@curl -sL -o kubebuilder https://go.kubebuilder.io/dl/latest/$(shell $(GO) env GOOS)/$(shell $(GO) env GOARCH)
	@mkdir -p ${HOME}/bin
	@chmod +x kubebuilder && mv kubebuilder ${HOME}/bin

# gomodifytags -all -add-tags json -w -transform camelcase --skip-unexported -file *.go
.PHONY: _install.gomodifytags
_install.gomodifytags: ## Install gomodifytags tool which is used to modify struct field tags.
	@$(GO) install github.com/fatih/gomodifytags@latest

.PHONY: _install.yq
_install.yq:
	@$(GO) install github.com/mikefarah/yq/v4@$(YQ_VERSION)

.PHONY: _install.gotestsum
_install.gotestsum:
	@$(GO) install gotest.tools/gotestsum@$(GO_TESTS_SUM_VERSION)

.PHONY: _install.gofumpt
_install.gofumpt:
	@$(GO) install mvdan.cc/gofumpt@latest

.PHONY: _install.grpcurl
_install.grpcurl:
	@$(GO) install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

.PHONY: _install.logcheck
_install.logcheck:
	@$(GO) install sigs.k8s.io/logtools/logcheck@latest
