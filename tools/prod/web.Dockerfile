FROM supinf/hugo:0.x AS build-app
ARG WEB_VERSION
ENV VERSION=${WEB_VERSION:-unknown}
COPY . /work/
WORKDIR /work
RUN sed -i -e "s/%{CI_BUILD_KEY}/${VERSION}/" layouts/_default/single.html
RUN sed -i -e "s/%{CI_BUILD_KEY}/${VERSION}/" layouts/index.html
RUN hugo --config config.toml --destination dist --baseURL='/' --cleanDestinationDir --ignoreCache
RUN find dist/js -maxdepth 1 -name "[^lm]*" -not -name '*js' -exec rm -rf {} \;
RUN rm -rf dist/scss dist/*.xml dist/**/*.xml /usr/share/nginx/html/50x.html

FROM swaggerapi/swagger-ui:v3.20.5 AS doc-gen

FROM nginx:1.15.11-alpine
COPY --from=doc-gen /usr/share/nginx/html /usr/share/nginx/html/doc/api
COPY --from=build-app /work/dist /usr/share/nginx/html
COPY openapi.yaml /usr/share/nginx/html/spec.yaml
COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh && apk add --no-cache bash
ENTRYPOINT ["/entrypoint.sh"]
