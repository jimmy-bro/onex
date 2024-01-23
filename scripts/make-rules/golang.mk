# ==============================================================================
# Makefile helper functions for golang
#

GO := go
GO_SUPPORTED_VERSIONS ?= 1.13|1.14|1.15|1.16|1.17|1.18|1.19|1.20|1.21

GOPATH := $(shell go env GOPATH)
ifeq ($(origin GOBIN), undefined)
	GOBIN := $(GOPATH)/bin
endif

COMMANDS ?= $(filter-out %.md, $(wildcard ${ONEX_ROOT}/cmd/*))
BINS ?= $(foreach cmd,${COMMANDS},$(notdir ${cmd}))

ifeq (${COMMANDS},)
  $(error Could not determine COMMANDS, set ONEX_ROOT or run in source dir)
endif
ifeq (${BINS},)
  $(error Could not determine BINS, set ONEX_ROOT or run in source dir)
endif

EXCLUDE_TESTS=github.com/superproj/onex/pkg/db,manifests

.PHONY: go.build.verify
go.build.verify: ## Verify supported go versions.
ifneq ($(shell $(GO) version | grep -q -E '\bgo($(GO_SUPPORTED_VERSIONS))\b' && echo 0 || echo 1), 0)
	$(error unsupported go version. Please make install one of the following supported version: '$(GO_SUPPORTED_VERSIONS)')
endif

.PHONY: go.build.%
go.build.%: ## Build specified applications with platform, os and arch.
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	#@ONEX_GIT_VERSION=$(VERSION) $(SCRIPTS_DIR)/build.sh $(COMMAND) $(PLATFORM)
	@if grep -q "func main()" $(ONEX_ROOT)/cmd/$(COMMAND)/*.go &>/dev/null; then \
		echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)" ; \
		CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) \
		-o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(COMMAND)$(GO_OUT_EXT) $(PRJ_SRC_PATH)/cmd/$(COMMAND) ; \
	fi

.PHONY: go.build
go.build: $(addprefix go.build., $(addprefix $(PLATFORM)., $(BINS))) ## Build all applications.

.PHONY: go.build.multiarch
go.build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix go.build., $(addprefix $(p)., $(BINS)))) ## Build all applications with all supported arch.

.PHONY: go.test
go.test: tools.verify.go-junit-report ## Run unit test and generate run report.
	@echo "===========> Run unit test"
	@set -o pipefail; $(GO) test -race -cover -coverprofile=$(OUTPUT_DIR)/coverage.out \
		-timeout=10m -shuffle=on -short -v `go list ./...|\
		egrep -v $(subst $(SPACE),'|',$(sort $(EXCLUDE_TESTS)))` 2>&1 | \
		tee >(go-junit-report --set-exit-code >$(OUTPUT_DIR)/report.xml)
	@sed -i '/mock_.*.go/d' $(OUTPUT_DIR)/coverage.out # remove mock_.*.go files from test coverage
	@$(GO) tool cover -html=$(OUTPUT_DIR)/coverage.out -o $(OUTPUT_DIR)/coverage.html

.PHONY: go.test.cover
go.test.cover: go.test ## Calculate test coverage.
	@$(GO) tool cover -func=$(OUTPUT_DIR)/coverage.out | \
		awk -v target=$(COVERAGE) -f $(SCRIPTS_DIR)/coverage.awk

.PHONY: go.updates
go.updates: tools.verify.go-mod-outdated ## Find outdated dependencies.
	@$(GO) list -u -m -json all | go-mod-outdated -update -direct
