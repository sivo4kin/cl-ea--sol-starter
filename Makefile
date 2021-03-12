GOBUILD=go build

.PHONY: all test clean build


bindings:
	cd build;./build --json contracts --pkg bindings --out bindings

clean:
	rm build/bindings/*.go
	rm ./digiu-cross-chain

build:
	$(GOBUILD)

start:
	./digiu-cross-chain

test:

all:
	make build
	make start


