ARG BASE_IMAGE="libs"

FROM golang:1.15.6-alpine3.12 AS libs
RUN apk --no-cache add g++ git

FROM ${BASE_IMAGE} as builder
ENV GO111MODULE=on
RUN go get -u "github.com/cosmtrek/air@v1.15.1"
RUN go get -u "github.com/go-delve/delve/cmd/dlv@v1.5.1"
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
ARG API_VERSION="dev"
ARG API_COMMIT="unknown"
COPY  . /go/src/github.com/xtream-d-labs/ai-gateway/
WORKDIR /go/src/github.com/xtream-d-labs/ai-gateway/api/src
RUN go build -o app -mod=vendor -ldflags "-s -w -X github.com/xtream-d-labs/ai-gateway/api/src/config.date=$(date +%Y-%m-%d) -X github.com/xtream-d-labs/ai-gateway/api/src/config.version=${API_VERSION} -X github.com/xtream-d-labs/ai-gateway/api/src/config.commit=${API_COMMIT}"
RUN mv app /

FROM golang:1.15.6-alpine3.12 as cache
RUN apk --no-cache add tini
COPY --from=libs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=libs /usr /usr
COPY --from=builder /root/.cache /root/.cache

FROM alpine:3.12 as prod
RUN apk --no-cache add bash openssl
COPY --from=libs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY api/templates /go/src/github.com/xtream-d-labs/ai-gateway/api/templates
COPY api/src/entrypoint.sh /
VOLUME ["/tmp/badger", "/tmp/work", "/tmp/simg"]
ARG API_VERSION="dev"
ENV AIG_API_VERSION=${API_VERSION} \
    DOCKER_HOST=unix:///var/run/docker.sock \
    GOPATH=/go
COPY --from=builder /app /
CMD ["/entrypoint.sh"]
