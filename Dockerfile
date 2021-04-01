FROM golang:alpine as build

RUN apk add --no-cache git gcc musl-dev linux-headers build-base

WORKDIR /p2p-bridge

ADD ./adapter/p2p-bridge .

RUN make keys

RUN make

FROM golang:alpine

COPY --from=build /p2p-bridge/bridge /bridge

COPY --from=build /p2p-bridge/config/my.yaml  config/

COPY --from=build /p2p-bridge/keys/srv3-ecdsa.key  keys/

COPY --from=build /p2p-bridge/keys/srv3-rsa.key  keys/

EXPOSE ${PORT}

ENTRYPOINT ["/bridge"]
