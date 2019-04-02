#!/bin/bash

find /usr/share/nginx/html/ -type f -name "*.html" | \
  xargs sed -i "s|http://localhost:8080|${SS_API_ENDPOINT}|g"

sed -i "s|localhost:9000|$( echo ${SS_API_ENDPOINT} | sed s/http:\\/\\/// )|" \
  /usr/share/nginx/html/spec.yaml
sed -i "s|https://petstore.swagger.io/v2/swagger.json|${SS_API_ENDPOINT}/spec.yaml|g" \
  /usr/share/nginx/html/doc/api/index.html

exec nginx -g "daemon off;"
