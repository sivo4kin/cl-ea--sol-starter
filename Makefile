GOBUILD=go build

.PHONY: all test clean bindings build build2


bindings:
	cd build && go build && ./build --json contracts --pkg bindings --out bindings

clean:
	rm ./build/bindings/*.go || rm ./truffle/build/contracts/*.json || rm ./digiu-cross-chain ./build/build

build:
	$(GOBUILD)

start:
	./digiu-cross-chain

test:
	go test ./test -v
all: clean bindings build start

CONTRACTSRC=$(shell find truffle/contracts -name '*.sol' || true)
build2:
	cd truffle;npx truffle build
	go run ./build --json truffle/build/contracts --pkg bindings --out bindings
	#build/solc/solc-static-linux $(CONTRACTSRC)
#	go run ./build --sol truffle/contracts --pkg bindings --out bindings


SOLCURL=https://github.com/ethereum/solidity/releases/download/v0.6.9

solc:
	mkdir -p compilers/solc
	wget -O compilers/solc/solc-static-linux $(SOLCURL)/solc-static-linux



