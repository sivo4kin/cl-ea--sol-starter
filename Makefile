.DEFAULT_GOAL := all

GOBUILD=go build

.PHONY: all test clean wrappers build keys

wrappers:
	cd truffle;npx truffle build;
	go run ./wrappers-builder --json truffle/build/contracts --pkg wrappers --out wrappers
	#build/solc/solc-static-linux $(CONTRACTSRC)
#	go run ./build --sol truffle/contracts --pkg wrappers --out wrappers

clean:
	rm ./wrappers/*.go || rm ./truffle/build/contracts/*.json || rm ./ea-starter ./wrappers-builder/wrappers-builder || rm keys/*.key || rm logs/*.log

build:
	cd ./wrappers-builder;$(GOBUILD)
	cd ./cmd;$(GOBUILD)

start:
	./cmd/cmd


testbls:
	#go test -v ./p2p/pub_sub_bls/libp2p_pubsub -run ^TestWithMinorFailure$
	#go test -v ./p2p/pub_sub_bls/libp2p_pubsub -run ^TestWithThreeGroups$
	go test -v ./p2p/pub_sub_bls/libp2p_pubsub -run ^TestBLS$

test:
	go test ./test -v

all: clean wrappers build start

CONTRACTSRC=$(shell find truffle/contracts -name '*.sol' || true)


SOLCURL=https://github.com/ethereum/solidity/releases/download/v0.6.9

solc:
	mkdir -p compilers/solc
	wget -O compilers/solc/solc-static-linux $(SOLCURL)/solc-static-linux
keys:
	go run key/keygen.go --prefix tls1
	go run key/keygen.go --prefix tls2
	go run key/keygen.go --prefix srv1
	go run key/keygen.go --prefix srv2
	go run key/keygen.go --prefix srv3
	go run key/keygen.go --prefix cl1
	go run key/keygen.go --prefix dht1
	go run key/keygen.go --prefix dht2






