# ScaleShift

A tool which assists us in making deep learning models locally or on cloud.

## Try on AWS

[![Launch Stack](https://cdn.rawgit.com/buildkite/cloudformation-launch-stack-button-svg/master/launch-stack.svg)](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=scaleshift&templateURL=https://scaleshift.s3.amazonaws.com/template.yaml)

## How to use locally

### Install dependent softwares

- [Docker](https://docs.docker.com/install/#get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Download the docker-compose.yml to your working directory

[docker-compose.yml](https://scaleshift.s3.amazonaws.com/docker-compose.yml)

### Start services

```console
docker-compose up
```

### Access the web UI

```console
open http://localhost:8080
```

## Contribution

1. Fork ([https://github.com/rescale/scaleshift/fork](https://github.com/rescale/scaleshift/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Create new Pull Request

## Copyright and license

Code released under the [MIT license](https://github.com/rescale/scaleshift/blob/master/LICENSE).
