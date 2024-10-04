.PHONY: lint
lint:
	golangci-lint run --config .github/golangci.yaml

.PHONY: bench
bench:
	go test -bench=. -benchmem

.PHONY: gofumpt
gofumpt:
	go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .
