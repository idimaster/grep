PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test: lint
	go test $(PKGS)

BIN_DIR := $(GOPATH)/bin
GOLINT := $(BIN_DIR)/golint

$(GOLINT):
	go get -u golang.org/x/lint/golint

.PHONY: lint
lint: $(GOLINT)
	$(GOLINT) ./...

BINARY := grep
VERSION ?= vlatest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	mkdir -p $(os)-amd64
	GOOS=$(os) GOARCH=amd64 go build -o release/$(os)-amd64/$(BINARY)$(if $(filter windows,$(os)),.exe)

.PHONY: release
release: windows linux darwin
