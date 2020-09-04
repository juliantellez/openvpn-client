GOPATH:=$(shell go env GOPATH)

# ----- Installing -----

.PHONY: install
install:
	go mod download

.PHONY: lint
lint:
	@ [ -e ./bin/golangci-lint ] || wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
	./bin/golangci-lint run

# ----- Testing -----

BUILDENV := CGO_ENABLED=0
TESTFLAGS := -short -cover

.PHONY: test
test:
	$(BUILDENV) go test $(TESTFLAGS) ./...

