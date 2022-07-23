GOLANGCI_BIN=$(LOCAL_BIN)/golangci-lint
$(GOLANGCI_BIN):
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint