# Singularity

https://www.sylabs.io/singularity/

[![scaleshift/singularity](http://dockeri.co/image/scaleshift/singularity)](https://hub.docker.com/r/scaleshift/singularity)

## Supported tags and respective `Dockerfile` links

・3.1 ([singularity/Dockerfile](https://github.com/rescale-labs/scaleshift/blob/master/tools/prod/docker/singularity/Dockerfile))  
・2.6-d2s ([doc2sin/Dockerfile](https://github.com/rescale-labs/scaleshift/blob/master/tools/prod/docker/doc2sin/Dockerfile))  

## Usage

singularity:2.6-d2s

```console
docker run --rm --privileged -v $(pwd):/output \
    -v /var/run/docker.sock:/var/run/docker.sock \
    scaleshift/singularity:2.6-d2s \
    --name app.simg alpine:3.9
```

singularity:3.1

```console
sudo docker run --rm -it --privileged -v $(pwd):/work \
    scaleshift/singularity:3.1 exec app.simg echo hello
```
