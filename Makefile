
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /tests)

.PHONY: lint
lint:
	gometalinter --config=config.json ./...

.PHONY: test
test:
	@go test -cover $(PKGS)

.PHONY: flag-gen 
flag-gen: 
	go run ./cmd/flag-gen/main.go -i github.com/zoumo/mamba/cmd/flag-gen/types \
	  -p github.com/zoumo/mamba \
	  -v 5
