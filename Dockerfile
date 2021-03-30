FROM golang:alpine as build

RUN apk add --no-cache ca-certificates build-base

WORKDIR /p2p-bridge

ADD ./adapter/p2p-bridge .

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux \
    go build -ldflags '-extldflags "-static"' -o bridge cmd/node.go

FROM golang:alpine

COPY --from=build /etc/ssl/certs/ca-certificates.crt \
     /etc/ssl/certs/ca-certificates.crt

COPY --from=build /p2p-bridge/bridge /bridge

COPY --from=build /p2p-bridge/config/my.yaml  config/

COPY --from=build /p2p-bridge/keys/srv3-ecdsa.key  keys/

COPY --from=build /p2p-bridge/keys/srv3-rsa.key  keys/

EXPOSE ${PORT}

ENTRYPOINT ["/bridge"]
