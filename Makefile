.PHONY: lint test vendor clean

default: lint test

.PHONY: lint 
lint:
	golangci-lint run

test:
	go test -v -cover ./...

yaegi_test:
	yaegi test -v .

vendor:
	go mod vendor

clean:
	rm -rf ./vendor