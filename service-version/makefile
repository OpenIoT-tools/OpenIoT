GO_PATHS =  $(shell go list -f '{{ .Dir }}' ./... | grep -E -v 'docs|cmd|mocks')

.PHONY: install
install:
	bash shell/install.sh

.PHONY: mocks
mocks:
	rm -R tests/mocks || true
	mockery

test:
	go test $(GO_PATHS)
