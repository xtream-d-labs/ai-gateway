# ScaleShift

A tool which assists us in making deep learning models locally or on cloud.

## Try on AWS

[![Launch Stack](https://cdn.rawgit.com/buildkite/cloudformation-launch-stack-button-svg/master/launch-stack.svg)](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=scaleshift&templateURL=https://s3-ap-northeast-1.amazonaws.com/scaleshift/template.yaml)

## How to use locally

### Install dependent softwares

- [Docker](https://docs.docker.com/install/#get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Download the docker-compose.yml to your working directory

[docker-compose-8080.yml](https://s3-ap-northeast-1.amazonaws.com/scaleshift/docker-compose-8080.yml)

### Start services

Set your host name as an environment variable `SS_API_ENDPOINT` and start the services!

```console
export SS_API_ENDPOINT=http://localhost:8080
docker-compose --file docker-compose-8080.yml up
```

### Access the web UI

```console
open http://localhost:8080
```

### You can use APIs directly as well

```console
docker pull tensorflow/tensorflow:1.13.1-py3
curl -i -X POST -H "Content-Type: application/json" \
    -d '{"image": "tensorflow/tensorflow:1.13.1-py3"}' \
    http://localhost:8080/api/v1/notebooks
```

## Contribution

1. Fork ([https://github.com/rescale-labs/scaleshift/fork](https://github.com/rescale-labs/scaleshift/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Create new Pull Request

## Copyright and license

Code released under the [MIT license](https://github.com/rescale-labs/scaleshift/blob/master/LICENSE).
