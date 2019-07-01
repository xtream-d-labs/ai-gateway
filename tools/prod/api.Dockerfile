FROM golang:1.12.6-alpine3.9 AS build-app
RUN apk --no-cache add build-base
COPY src /go/src/github.com/rescale-labs/scaleshift/api/src
WORKDIR /go/src/github.com/rescale-labs/scaleshift/api/src
ARG API_VERSION
ARG API_COMMIT
RUN go build -ldflags \
    "-s -w -X github.com/rescale-labs/scaleshift/api/src/config.date=$(date +%Y-%m-%d) -X github.com/rescale-labs/scaleshift/api/src/config.version=${API_VERSION} -X github.com/rescale-labs/scaleshift/api/src/config.commit=${API_COMMIT}" \
    generated/cmd/scale-shift-server/main.go
RUN mv main /app

FROM alpine:3.9
RUN apk --no-cache add "ca-certificates=20190108-r0" "openssl=1.1.1b-r1" "bash=4.4.19-r1"
ARG API_VERSION
ENV SS_API_VERSION=${API_VERSION:-unknown} \
    DOCKER_HOST=unix:///var/run/docker.sock \
    GOPATH=/go
VOLUME ["/tmp/badger", "/tmp/work", "/tmp/simg"]
COPY templates /go/src/github.com/rescale-labs/scaleshift/api/templates
COPY --from=build-app /app /app
COPY src/entrypoint.sh /
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
