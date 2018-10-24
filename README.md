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

## Kubernetes and Minikube commands

In the course the instructor demonstrated several concepts related to Kubernetes using Minikube, which is a single Node Kubernetes cluster that you can run locally.

### Starting and defining Docker environment variables

Start Minikube by issuing the following command:

```text
minikube start
```

Minikube will start and you should see output similar to

```text
Starting local Kubernetes v1.10.0 cluster...
Starting VM...
Getting VM IP address...
Moving files into cluster...
Setting up certs...
Connecting to cluster...
Setting up kubeconfig...
Starting cluster components...
Kubectl is now configured to use the cluster.
Loading cached images from config file.
```

Once started, one of the first things you'll want to do is to switch your local Docker environment to the Minikube Docker environment. Therefore you can first issue the following command:

```text
minikube docker-env
```

which should return something like this...

```text
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://192.168.99.100:2376"
export DOCKER_CERT_PATH="/Users/hectorrios/.minikube/certs"
export DOCKER_API_VERSION="1.35"
# Run this command to configure your shell:
# eval $(minikube docker-env)
```

which displays the environment variables that would be set to switch to the Minikube Docker registry in the Minikube VM. In addition it also gives you the command to enter into the terminal to perform the switch. So to actually perform the switch run the following:

```text
eval $(minikube docker-env)
```
Once complete, if you were to run a 

```text
docker ps
```

Then you should get a listing of all running Kubernetes containers.

### Verify Minikube cluster information

The command to display which Kubernetes cluster you're working against is...

```text
kubectl cluster-info
``` 
### Deploying a Pod, the smallest and easiest deployment unit

A "Pod" represents either a single container or set of containers but I believe that the best practice is to deploy a container per Pod.

One way of deploying resources in Kubernetes is via a YAML configuration file. For our "Pod" we can use the file,

```text
k8s-pod.yml
```
Without going into much detail, some of the important bits are

* **kind** This is the resource type and for our pod the value is **Pod**
* **labels** represent key value pair and you are free to decide what these should be.
* **spec -> containers** Defines the image that will be used to create the container for this Pod

In order to deploy the Pod you would issue the following command:

```text
kubectl create -f k8s-pod.yml
```

Once deployed you can issue a

```text
kubectl get pods
```

to see a list of all running Pods. You should see an entry for a Pod called **cloud-native-go**

You can also some more detailed information on the Pod by issuing a 

```text
kubectl describe pod cloud-native-go
```
which should provide a plethora of information on the running Pod.

#### Accessing the deployed Pod

In order to access the running Pod you can use "port-forward" command. Be aware that this command blocks when issued. For example lets listen locally on port 8080 from the Pod's port 8080

```text
kubectl port-forward <Pod Name> 8080:8080
```
In our case the "Pod Name" is **cloud-native-go**

In order to get the labels on a Pod issue something like

```text
kubectl get pods cloud-native-go --show-labels
```

To add a new label dynamically you can issue the following:

```text
kubectl label pod cloud-native-go hello=world
```
Which adds a new key (hello) value (world) pair to the **cloud-native-go** Pod.

You can also change existing label values such as changing the value of our "hello" label...

```text
kubectl label pod cloud-native-go-844dcb4dd5-jlf7z hello=hector --overwrite
```
### Namespaces

A Namespace is kind of like a "package" in Java and namespaces in PHP

To see all existing namespaces you can issue:

```text
kubectl get ns
```

On my system I get:

```text
NAME          STATUS   AGE
default       Active   5d
kube-public   Active   5d
kube-system   Active   5d
```

If I want to see all Pods in the "kube-system" name space then I could issue:

```text
kubectl get pods --namespace kube-system 
```

The fallback namespace for resource is the **default** namespace. If you issue a **describe** command on our **cloud-native-go** Pod then you'll see that it's in the **default** namespace.

```text
kubectl describe pod cloud-native-go | grep Namespace
```
In order to create a new Namespace you need a YAML file which declares the new Namespace. For example, check out the **k8s-namespace.yml** file to see how to delcare a Namespace.

We use the **create** to create one once we have the YAML file defined.

```text
kubectl create -f k8s-namespace.yml
```
With our newly created Namespace we can create our **cloud-native-go** Pod again but this time in our new namespace

```text
$ kubectl create -f k8s-pod.yml --namespace cloud-native-go
```

and now if we have a look at our Namespace again we'll see a Pod in there. In addition, our original Pod in the **default** is still present so we have two Pods but in different namespaces

```text
$ kubectl get pods --namespace cloud-native-go
NAME              READY   STATUS    RESTARTS   AGE
cloud-native-go   1/1     Running   0          50s
```
The interesting thing about namespaces is that if you delete a namespace, then everything in the Namespace will be deleted as well

```text
$ kubectl delete -f k8s-namespace.yml
namespace "cloud-native-go" deleted
```
**Sooooo, be careful when deleting Namespaces because the CLI will not warn you that the resources within the namespace will also be deleted.**


