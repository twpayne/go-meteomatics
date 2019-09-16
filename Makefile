.PHONY: none
none:

.PHONY: coverage.out
coverage.out:
	go test -cover -covermode=count -coverprofile=$@ ./...

.PHONY: format
format:
	find . -name \*.go | xargs gofumports -w

.PHONY: html-coverage
html-coverage: coverage.out
	go tool cover -html=$<

.PHONY: install-tools
install-tools:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- v1.18.0
	GO111MODULE=off go get -u \
		golang.org/x/tools/cmd/cover \
		github.com/mattn/goveralls \
		mvdan.cc/gofumpt/gofumports

.PHONY: lint
lint:
	go vet ./...
	./bin/golangci-lint run

.PHONY: test
test:
	go test ./...