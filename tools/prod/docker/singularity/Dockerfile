# Singularity v3.2
# docker run --rm -it scaleshift/singularity:3.2

FROM alpine:3.10 AS build-gosu
RUN apk --no-cache add "curl=7.65.1-r0"
RUN curl --location --silent --show-error --out /gosu \
    https://github.com/tianon/gosu/releases/download/1.11/gosu-amd64
RUN chmod +x /gosu

FROM golang:1.12.7-alpine3.10 AS build-app
ENV SINGULARITY_VERESION=v3.3.0
RUN apk --no-cache add "build-base=0.5-r1" "linux-headers=4.19.36-r0" \
       "openssl-dev=1.1.1c-r0" "util-linux-dev=2.33.2-r0" "gawk=5.0.1-r0" \
       "git=2.22.0-r0" \
    && mkdir -p /go/src/github.com/sylabs \
    && cd /go/src/github.com/sylabs \
    && git clone --depth=1 -b "${SINGULARITY_VERESION}" https://github.com/sylabs/singularity.git \
    && cd singularity \
    && ./mconfig \
    && cd builddir \
    && make "-j$(grep -c ^processor /proc/cpuinfo 2>/dev/null || 1)" \
    && make install \
    && rm -rf /usr/local/go /usr/local/etc/bash_completion.d

FROM alpine:3.10
RUN apk --no-cache add "bash=5.0.0-r0" "ca-certificates=20190108-r0" \
       "squashfs-tools=4.3-r5" "libseccomp=2.4.1-r0" \
    && rm -rf /usr/share/terminfo /lib/apk /etc/apk /etc/terminfo
COPY --from=build-app /usr/local /usr/local
COPY --from=build-gosu /gosu /usr/local/bin/gosu
WORKDIR /work
COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
CMD ["bash"]
