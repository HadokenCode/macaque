PROJECT  = macaque
PACKAGE  = github.com/wildnature/macaque
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell cat .version)
GOPATH   = $(CURDIR)/.gopath
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
PROTO    = $(BASE)/protobuf
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/" | grep -v "^$(PACKAGE)/testInt/"))
TESTPKGS = $(shell env GOPATH=$(GOPATH) $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))
TESTINTPKGS = $(shell env GOPATH=$(GOPATH) $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))
GO      = go
GODOC   = godoc
GOFMT   = gofmt
GLIDE   = glide
TIMEOUT = 15
V = 0	
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

$(BASE): ; $(info $(M) setting GOPATH…)
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

# Tools

GOGRPC = $(BIN)/grpc
$(BIN)/grpc: | $(BASE) ; $(info $(M) building proto's)
	$Q go get google.golang.org/grpc


GOLINT = $(BIN)/golint
$(BIN)/golint: | $(BASE) ; $(info $(M) building golint…)
	$Q go get github.com/golang/lint/golint

GOCOVMERGE = $(BIN)/gocovmerge
$(BIN)/gocovmerge: | $(BASE) ; $(info $(M) building gocovmerge…)
	$Q go get github.com/wadey/gocovmerge

GOCOV = $(BIN)/gocov
$(BIN)/gocov: | $(BASE) ; $(info $(M) building gocov…)
	$Q go get github.com/axw/gocov/...

GOCOVXML = $(BIN)/gocov-xml
$(BIN)/gocov-xml: | $(BASE) ; $(info $(M) building gocov-xml…)
	$Q go get github.com/AlekSi/gocov-xml

GO2XUNIT = $(BIN)/go2xunit
$(BIN)/go2xunit: | $(BASE) ; $(info $(M) building go2xunit…)
	$Q go get github.com/tebeka/go2xunit

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short -v       ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: vendor | $(BASE) ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q cd $(BASE) && $(GO) test -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

test-xml: fmt lint vendor | $(BASE) $(GO2XUNIT) ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests with xUnit output
	$Q cd $(BASE) && 2>&1 $(GO) test -timeout 20s -v $(TESTPKGS) | tee test/tests.output
	$(GO2XUNIT) -fail -input test/tests.output -output test/tests.xml

COVERAGE_MODE = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
COVERAGE_XML = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML = $(COVERAGE_DIR)/index.html
.PHONY: test-coverage test-coverage-tools
test-coverage-tools: | $(GOCOVMERGE) $(GOCOV) $(GOCOVXML)
test-coverage: COVERAGE_DIR := $(CURDIR)/test/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
test-coverage: fmt lint vendor test-coverage-tools | $(BASE) ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)/coverage
	$Q cd $(BASE) && for pkg in $(TESTPKGS); do \
		$(GO) test \
			-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $$pkg | \
					grep '^$(PACKAGE)/' | grep -v '^$(PACKAGE)/vendor/' | \
					tr '\n' ',')$$pkg \
			-covermode=$(COVERAGE_MODE) \
			-coverprofile="$(COVERAGE_DIR)/coverage/`echo $$pkg | tr "/" "-"`.cover" $$pkg ;\
	 done
	$Q $(GOCOVMERGE) $(COVERAGE_DIR)/coverage/*.cover > $(COVERAGE_PROFILE)
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: lint
lint: vendor | $(BASE) $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q cd $(BASE) && ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=1 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

.PHONY: vet
vet: ; $(info $(M) running go vet...) @ ## Run go vet on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GO) vet  $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

# Dependency api

glide.lock: glide.yaml | $(BASE) ; $(info $(M) updating dependencies…)
	$Q cd $(BASE) && $(GLIDE) update
	@touch $@
vendor: glide.lock | $(BASE) ; $(info $(M) retrieving dependencies…)
	$Q cd $(BASE) && $(GLIDE) --quiet install
	@ln -nsf . vendor/src
	@touch $@

# Misc
.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(GOPATH)
	@rm -rf bin
	@rm -rf build
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: release
release: release-api  release-store-mongodb| $(BASE) $(GOGRPC) ; $(info $(M) Building docker images…) @ ## Build docker images

.PHONY: release-api
release-api: clean vendor  | $(BASE) $(GOGRPC) ; $(info $(M) building executables…) @ ## Build docker image (api)
	@echo "Building macaque-api..."
	$Q cd $(BASE) && sh scripts/release.sh api cmd/api/main.go

.PHONY: release-store-mongodb
release-store-mongodb: clean vendor | $(BASE) $(GOGRPC) ; $(info $(M) building executables…) @ ## Build docker image (store-mongodb)
	@echo "Building macaque-store-mongodb..."
	$Q cd $(BASE) && sh scripts/release.sh store-mongodb cmd/store/mongodb/main.go

.PHONY: publish-api
publish-api: release-api | $(BASE) ; $(info $(M) publishing dcker image) @ ## Publish docker image (api)
	@echo "Publishing macaque-api..."
	$Q cd $(BASE) && sh scripts/publish.sh store-mongodb

.PHONY: publish-store-mongodb
publish-store-mongodb: release-store-mongodb | $(BASE) ; $(info $(M) publishing dcker image) @ ## Publish docker image (store-mongodb)
	@echo "Publishing macaque-store-mongodb..."
	$Q cd $(BASE) && sh scripts/publish.sh store-mongodb

.PHONY: test-int
test-int: clean vendor test-int-api | $(BASE) ; $(info $(M) Running integration tests…) @ ## Run integration tests

.PHONY: testInt-api
testInt-api: clean vendor  | $(BASE) ; $(info $(M) Running integration tests (api)…) @ ## Run integration tests for module api
	$Q cd $(BASE) sh scripts/run-integration-tests.sh api 3031 cmd/api/main.go

.PHONY: testInt-store-mongodb
testInt-store-mongodb: clean vendor  | $(BASE) ; $(info $(M) Running integration tests (store-mongodb)…) @ ## Run integration tests for module store-mongodb
	$Q cd $(BASE) && sh scripts/run-integration-tests.sh store-mongodb 11001 cmd/store/mongodb/main.go

	
.PHONY: proto
proto: $(PROTO); $(info $(M) Running pb generation) @ ## Generating golang code from protobuffers definition
	$Q cd $(BASE) && sh scripts/proto2golang.sh
