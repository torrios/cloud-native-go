# cloud-native-go

This repo contains the exercises and examples described in the online course "Getting Started with Cloud Native Go"

## Build the Linux executable

In order to build the Docker image you'll first need to build Linux executable so that it can be copied into the image.

Change into the *json_marshalling* directory and run

```text
env GOOS=linux GOARCH=amd64 go build -v .
```

## Build Docker image

Change into the **json_marshalling** directory
```text
docker build -t cloud-native-go:1.0.0-alpine .
```
## Push image to DockerHub

### Tag the built image for DockerHub
In order to push to DockerHub you'll need to have a DockerHub account which is free to create.

First of all, you'll need to use the *docker tag* command to tag the image you want to push up

Assuming that my DockerHub username is: torrios , then the command would be...

```text
docker tag cloud-native-go:1.0.0-alpine torrios/cloud-native-go:1.0.0-alpine 
``` 

which takes the existing cloud-native-go:1.0.0-alpine image and tags it as torrios/cloud-native-go:1.0.0-alpine and in the process creates a new image or clones the original image

### Login to Docker 

In order to push to DockerHub you'll need to login which you can do via the *docker login* command.

```text
docker login
```
### Push image up
Once you've successfully provided your credentials you can use the *docker push* to push the image up to DockerHub

```text
docker push torrios/cloud-native-go:1.0.0-alpine 
```
