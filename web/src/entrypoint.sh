#!/bin/bash

find /usr/share/nginx/html/ -type f -name "*.html" -print0 | \
  xargs -0 sed -i "s|http://localhost:8080||g"

sed -i "s|localhost:9000||" /usr/share/nginx/html/spec.yaml
sed -i "s|https://petstore.swagger.io/v2/swagger.json|/spec.yaml|g" \
  /usr/share/nginx/html/doc/api/index.html

exec nginx -g "daemon off;"
