.DEFAULT_GOAL := all

GOBUILD=go build

.PHONY: all test clean wrappers build keys

npm:
	@if [ -d truffle/node_modules ]; then \
  			echo "installed"; \
  			else \
  			cd truffle;npm i;fi

wrappers: npm
	cd truffle;npx truffle build;
	go run ./wrappers-builder --json truffle/build/contracts --pkg wrappers --out wrappers

clean:
	rm ./wrappers/*.go || rm ./truffle/build/contracts/*.json || rm ./ea-starter ./wrappers-builder/wrappers-builder || rm keys/*.key || rm logs/*.log

build: keys
	cd ./wrappers-builder;$(GOBUILD)
	go build -o node  cmd/node.go

start: build
	./node


testbls:
	#go test -v ./p2p/pub_sub_bls/libp2p_pubsub -run ^TestWithMinorFailure$
	#go test -v ./p2p/pub_sub_bls/libp2p_pubsub -run ^TestWithThreeGroups$
	go test -v ./p2p/pub_sub_bls/libp2p_pubsub -run ^TestBLS$

test:
	go test ./test -v

all: build start

CONTRACTSRC=$(shell find truffle/contracts -name '*.sol' || true)


SOLCURL=https://github.com/ethereum/solidity/releases/download/v0.6.9

solc:
	mkdir -p compilers/solc
	wget -O compilers/solc/solc-static-linux $(SOLCURL)/solc-static-linux
keys:
	go run key/keygen.go --prefix srv3






