#!/bin/sh
set -u

wk_dir="$(pwd)"
if [ ! -d "${wk_dir}/.git" ]; then
  echo 'This script must be executed on local git repository root dir.' 1>&2
  exit 1
fi

DOCKER_REPO="scaleshift"
tag=$( git describe --tags --abbrev=0 2>/dev/null )
version="${tag:-$( git rev-parse --abbrev-ref @ )}"
commit=$( git rev-parse --short --verify HEAD )

cat << EOT

[ Environment variables ]

DOCKER_REPO: ${DOCKER_REPO}
VERSION:     ${version}

EOT
set -e

cat << EOT > tools/prod/latest-version.json
{
  "version": "${version}",
  "commit": "${commit}",
  "date": "$(date +%Y-%m-%d)"
}
EOT

docker build -f tools/prod/api.Dockerfile -t "${DOCKER_REPO}/api:${version}-${commit}" \
    --build-arg API_VERSION="${version}" --build-arg API_COMMIT="${commit}" api
cp -f spec/openapi.yaml web/src/
docker build -f tools/prod/web.Dockerfile -t "${DOCKER_REPO}/web:${version}-${commit}" \
    --build-arg WEB_VERSION="${version}-${commit}" web/src

docker tag "${DOCKER_REPO}/api:${version}-${commit}" "${DOCKER_REPO}/api:latest"
docker tag "${DOCKER_REPO}/web:${version}-${commit}" "${DOCKER_REPO}/web:latest"

docker push "${DOCKER_REPO}/api:${version}-${commit}"
docker push "${DOCKER_REPO}/web:${version}-${commit}"
docker push "${DOCKER_REPO}/api:latest"
docker push "${DOCKER_REPO}/web:latest"

# shellcheck disable=SC1083,SC2086
sed -e s/%{version}/${version}-${commit}/g tools/prod/docker-compose.template.yml > tools/prod/docker-compose.tmp.yml
sed -e s/%{publish_port}/80/g    tools/prod/docker-compose.tmp.yml >      tools/prod/docker-compose-80.yml
sed -e s/%{publish_port}/8080/g  tools/prod/docker-compose.tmp.yml >      tools/prod/docker-compose-8080.yml

rm -f tools/prod/docker-compose.tmp.yml
