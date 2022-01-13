FROM golang:alpine AS builder

WORKDIR /build
COPY . /build/chia_exporter
RUN apk add --update --no-cache --virtual build-dependencies \
 && cd chia_exporter \
 && go build -tags netgo

FROM alpine
COPY --from=builder /build/chia_exporter/chia_exporter /usr/bin/chia_exporter

EXPOSE 9133

#ENV FULL_NODE_CERT
#ENV FULL_NODE_KEY
#ENV FULL_NODE_RPC_ENDPOINT
#ENV WALLET_RPC_ENDPOINT
ENV CHIA_FORK=chia

CMD /usr/bin/chia_exporter
