#!/bin/sh

swagger-codegen generate -i /work/spec/openapi.yaml -o /work/src/generated -l javascript

cd /work/src/generated || exit
rm -rf v1
mv src v1
rm -rf .swagger-codegen src test ./*.json ./.* git* ./*.opts ./*.md

cd /work/src || exit
yarn install
browserify generated/v1.js | uglifyjs --compress --mangle > static/js/lib/api.v1.js
