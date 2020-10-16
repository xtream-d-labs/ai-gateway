# Singularity

https://www.sylabs.io/singularity/

[![aigateway/singularity](http://dockeri.co/image/aigateway/singularity)](https://hub.docker.com/r/aigateway/singularity)

## Supported tags and respective `Dockerfile` links

・3.4 ([singularity/Dockerfile](https://github.com/xtream-d-labs/ai-gateway/blob/master/tools/prod/docker/singularity/Dockerfile))  
・2.6-d2s ([doc2sin/Dockerfile](https://github.com/xtream-d-labs/ai-gateway/blob/master/tools/prod/docker/doc2sin/Dockerfile))  

## Usage

singularity:2.6-d2s

```console
docker run --rm --privileged -v $(pwd):/output \
    -v /var/run/docker.sock:/var/run/docker.sock \
    aigateway/singularity:2.6-d2s \
    --name app.simg alpine:3.12
```

singularity:3.4

```console
sudo docker run --rm -it --privileged -v $(pwd):/work \
    aigateway/singularity:3.4 exec app.simg echo hello
```
