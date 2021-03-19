GOBUILD=go build

.PHONY: all test clean wrappers build

wrappers:
	cd truffle;npx truffle build
	go run ./build --json truffle/build/contracts --pkg wrappers --out wrappers
	#build/solc/solc-static-linux $(CONTRACTSRC)
#	go run ./build --sol truffle/contracts --pkg wrappers --out wrappers

clean:
	rm ./wrappers/*.go || rm ./truffle/build/contracts/*.json || rm ./ea-starter ./build/build

build:
	$(GOBUILD)

start:
	./ea-starter

test:
	go test ./test -v
all: clean build wrappers start

CONTRACTSRC=$(shell find truffle/contracts -name '*.sol' || true)


SOLCURL=https://github.com/ethereum/solidity/releases/download/v0.6.9

solc:
	mkdir -p compilers/solc
	wget -O compilers/solc/solc-static-linux $(SOLCURL)/solc-static-linux



