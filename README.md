<div align="center">
  <img src="https://xtreme-d.net/wp-content/uploads/2020/05/logo.svg">
  <h1>AI Gateway</h1>
  <span>An Open Source Machine Learning Tool for making & training models.</span>
</div>

-----------------

## Overview

AI Gateway is a client web application which makes it easier for researchers to build machine learning models locally and to train them on premise/cloud!

## Key Capabilities

- A client application that has web-based user-friendly interfaces
- Can be run anywhere (Linux, macOS, Windows..)
- Supports any Python-based dockerized algorithms or software
- Fully integrated with [NGC](https://ngc.nvidia.com/), [Rescale platform](https://www.rescale.com/) & [Kubernetes](https://kubernetes.io/)

## How It Works

### 1. Setup an AI Gateway client

<img src="https://raw.github.com/wiki/rescale-labs/scaleshift/img/how-it-works-1.png">

With just only [3 steps!](https://github.com/xtream-d-labs/ai-gateway#local-installation)

### 2. Download machine learning software

<img src="https://raw.github.com/wiki/rescale-labs/scaleshift/img/how-it-works-2.png">

You can pull any docker images to the client from NGC, DockerHub or your private registry with **just one click**.

### 3. Create a new workspace

<img src="https://raw.github.com/wiki/rescale-labs/scaleshift/img/how-it-works-3.png">

When you click a `run` button, AI Gateway wraps the image with [Jupyter notebook](https://jupyter.org/) & run it as a docker container. Then you can build your own models on your specified software on it.

### 4. Train your models

<img src="https://raw.github.com/wiki/rescale-labs/scaleshift/img/how-it-works-5.png">

In order to train the models in a parallel and distributed way, you can choose a Kubernetes cluster or Rescale platform. AI Gateway converts the image to [Singularity](https://www.sylabs.io/docs/) automatically if it’s needed.

## Get Started

Official Builds

[![aigateway/api](http://dockeri.co/image/aigateway/api)](https://hub.docker.com/r/aigateway/api/)

### Try AI Gateway on AWS

#### 1. Create an EC2 instance with CloudFormation

[![Launch Stack](https://cdn.rawgit.com/buildkite/cloudformation-launch-stack-button-svg/master/launch-stack.svg)](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/new?stackName=aigateway&templateURL=https://s3-ap-northeast-1.amazonaws.com/ai-gateway/template.yaml)

#### 2. Access the Web UI

```console
public_ip=$( sh -c "$( aws cloudformation describe-stacks --stack-name "aigateway" \
  --query 'Stacks[0].Outputs[?OutputKey==`PublicIPs`].OutputValue' \
  --output text )" )
open "http://${public_ip}"
```

### Local Installation

#### 1. Install dependent softwares

- [Docker](https://docs.docker.com/install/#get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

#### 2. Download the latest configuration

```console
curl -so docker-compose.yml https://s3-ap-northeast-1.amazonaws.com/ai-gateway/docker-compose-8080.yml
```

#### 3. Start services

```console
docker-compose up
```

### Usage

#### Access the Web UI

[http://localhost:8080](http://localhost:8080)

#### Access APIs directly

```console
curl -sX POST -H "Content-Type: application/json" \
    -d '{"image": "tensorflow/tensorflow:2.3.0"}' \
    http://localhost:8080/api/v1/notebooks
curl -sX GET -H "Content-Type: application/json" \
    http://localhost:8080/api/v1/notebooks
```

### Contribution

1. Fork ([https://github.com/xtream-d-labs/ai-gateway/fork](https://github.com/xtream-d-labs/ai-gateway/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Create new Pull Request

### Copyright and license

Code released under the [MIT license](https://github.com/xtream-d-labs/ai-gateway/blob/master/LICENSE).
