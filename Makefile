PROJECT_NAME=almighty-test-runner

PACKAGE_NAME := github.com/almighty/almighty-test-runner
BINARY=alm-test

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

# Build configuration
BUILD_TIME=`date -u '+%Y-%m-%dT%H:%M:%SZ'`
COMMIT=$(shell git rev-parse HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	COMMIT := $(COMMIT)-dirty
endif

# Pass in build time variables to main
LDFLAGS="-X main.Commit=${COMMIT} -X main.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL := all

help: ## Get help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Removes binary
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: test
test: ## Runs ginkgo tests
	ginkgo -r -p -cover

.PHONY: watch
watch: ## Continuously run tests whenever source code changes
	ginkgo watch -r -p

.PHONY: install
install: ## Fetches all dependencies using Glide
	glide --verbose install

.PHONY: up
up: ## Updates all dependencies defined for glide
	glide up


.PHONY: check
check: ## Concurrently runs a whole bunch of static analysis tools
	gometalinter --vendor --deadline 100s ./...

.PHONY: all
all: clean install $(BINARY) test ## (default) Performs clean deps build test

$(BINARY): $(SOURCES)
	go build -v -ldflags ${LDFLAGS} -o ${BINARY}
