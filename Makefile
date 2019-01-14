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

BINARY := mytool
VERSION ?= vlatest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: windows linux darwin
