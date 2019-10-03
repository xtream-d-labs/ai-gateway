ARG BASE_IMAGE="libs"

FROM golang:1.13.1-alpine3.10 AS libs
RUN apk --no-cache add g++ git

FROM ${BASE_IMAGE} as builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
ARG API_VERSION="dev"
ARG API_COMMIT="unknown"
COPY  . /go/src/github.com/rescale-labs/scaleshift/
WORKDIR /go/src/github.com/rescale-labs/scaleshift/api/src
RUN go build -o app -mod=vendor -ldflags "-s -w -X github.com/rescale-labs/scaleshift/api/src/config.date=$(date +%Y-%m-%d) -X github.com/rescale-labs/scaleshift/api/src/config.version=${API_VERSION} -X github.com/rescale-labs/scaleshift/api/src/config.commit=${API_COMMIT}"
RUN mv app /

FROM golang:1.13.1-alpine3.10 as cache
RUN apk --no-cache add tini
COPY --from=libs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=libs /usr /usr
COPY --from=builder /root/.cache /root/.cache

FROM alpine:3.10 as prod
RUN apk --no-cache add bash openssl
COPY --from=libs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY api/templates /go/src/github.com/rescale-labs/scaleshift/api/templates
COPY api/src/entrypoint.sh /
VOLUME ["/tmp/badger", "/tmp/work", "/tmp/simg"]
ARG API_VERSION="dev"
ENV SS_API_VERSION=${API_VERSION} \
    DOCKER_HOST=unix:///var/run/docker.sock \
    GOPATH=/go
COPY --from=builder /app /
CMD ["/entrypoint.sh"]
