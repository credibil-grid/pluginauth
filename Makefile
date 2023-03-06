.PHONY: lint test vendor clean

default: lint test

.PHONY: lint 
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: yaegi 
yaegi:
	yaegi test -v .

vendor:
	go mod vendor

clean:
	rm -rf ./vendor