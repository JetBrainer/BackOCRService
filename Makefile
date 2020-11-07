BIN := server
.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...
## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

.PHONY: build-tokenizer
## build-tokenizer: build the tokenizer application
build-tokenizer:
	${MAKE} -c tokenizer build
.PHONY: build
## Build project: build command
build-server:
	cd cmd/server && go build -o ../../bin/server
build-client:
	cd cmd/client-rest && go build -o ../../bin/client
.PHONY: run
## Run project: run command
run-server:
	./bin/server
run-client:
	./bin/client
.PHONY: clean
## Clean: clean command
clean:
	go clean
	clear
.PHONY: all
## Cleans binary file clears terminal and build/run project: all make file
all-server-run:
	cd cmd/server && go build . ../../bin/${BIN} && cd ../../ && ./bin/${BIN}

.PHONY: linter
## golangci-linter: go linter
lint:
	golangci-lint run