
default: lint test yaegi

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: yaegi 
yaegi:
	yaegi test -v .

.PHONY: vendor
vendor:
	go mod vendor
	
.PHONY: clean
clean:
	rm -rf ./vendor