# Singularity v3.4
# docker run --rm -it aigateway/singularity:3.4

FROM alpine:3.12 AS build-gosu
RUN apk --no-cache add curl
RUN curl --location --silent --show-error --out /gosu \
    https://github.com/tianon/gosu/releases/download/1.11/gosu-amd64
RUN chmod +x /gosu

FROM golang:1.15.2-alpine3.12 AS build-app
ENV SINGULARITY_VERESION=v3.6.3
RUN apk --no-cache add build-base linux-headers util-linux-dev \
        openssl-dev cryptsetup gawk git \
    && mkdir -p /go/src/github.com/sylabs \
    && cd /go/src/github.com/sylabs \
    && git clone --depth=1 -b "${SINGULARITY_VERESION}" https://github.com/hpcng/singularity.git \
    && cd singularity \
    && ./mconfig \
    && cd builddir \
    && make "-j$(grep -c ^processor /proc/cpuinfo 2>/dev/null || 1)" \
    && make install \
    && rm -rf /usr/local/go /usr/local/etc/bash_completion.d

FROM alpine:3.12
RUN apk --no-cache add bash ca-certificates squashfs-tools libseccomp \
    && rm -rf /usr/share/terminfo /lib/apk /etc/apk /etc/terminfo
COPY --from=build-app /usr/local /usr/local
COPY --from=build-gosu /gosu /usr/local/bin/gosu
WORKDIR /work
COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
CMD ["sh"]
