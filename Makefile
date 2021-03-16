GOBUILD=go build

.PHONY: all test clean bindings build


bindings:
	cd build && go buildls  && ./build --json contracts --pkg bindings --out bindings

clean:
	rm ./build/bindings/*.go ./digiu-cross-chain ./build/build

build:
	$(GOBUILD)

start:
	./digiu-cross-chain

test:
	go test ./test -v
all: clean bindings build start



