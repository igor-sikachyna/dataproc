BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

# Update the ldflags with the app, client & server names
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=dataproc \
	-X github.com/cosmos/cosmos-sdk/version.AppName=dataprocd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

###########
# Install #
###########

all: install

install:
	@echo "--> ensure dependencies have not been modified"
	@go mod verify
	@echo "--> installing dataprocd"
	@go install $(BUILD_FLAGS) -mod=readonly ./cmd/dataprocd

init:
	./scripts/init.sh

##################
###  Protobuf  ###
##################

proto-all:
	@cd ./x/dataproc && make proto-all

proto-gen:
	@cd ./x/dataproc && make proto-gen

proto-format:
	@cd ./x/dataproc && make proto-format

proto-lint:
	@cd ./x/dataproc && make proto-lint

.PHONY: proto-all proto-gen proto-format proto-lint

###################
### Development ###
###################

mock-expected-keepers:
	@mockgen -source=x/dataproc/types/expected_keepers.go \
		-package testutil \
		-destination=x/dataproc/testutil/expected_keepers_mocks.go