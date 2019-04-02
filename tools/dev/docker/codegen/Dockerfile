FROM supinf/swagger-codegen:2.3

RUN apk --no-cache add nodejs yarn \
    && yarn install \
    && yarn global add browserify uglify-js

ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh

VOLUME /spec
ENTRYPOINT ["/entrypoint.sh"]
