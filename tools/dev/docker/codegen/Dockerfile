FROM supinf/swagger-codegen:2.3

RUN apk --no-cache add "nodejs=10.14.2-r0" "yarn=1.12.3-r0" \
    && yarn global add browserify uglify-js

COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

VOLUME /spec
ENTRYPOINT ["/entrypoint.sh"]
