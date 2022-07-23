LOCAL_BIN=$(CURDIR)/bin
PROJECT_NAME=witcher


include bin-deps.mk

.PHONY: lint
lint: $(GOLANGCI_BIN)
	$(GOENV) $(GOLANGCI_BIN) run --fix ./...

.PHONY: test
test:
	$(GOENV) go test -race ./...


