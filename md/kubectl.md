# Kubectl

## Get pods

~~~bash
$ kubectl get pods
~~~

## Select the context

~~~bash
$ kubectl config use-context *context-name*
~~~

## get list of namespaces in current context

~~~bash
$ kubectl get namespace
~~~

## get pods in namespace

~~~bash
$ kubectl get po -n *namespace name*
~~~

## get *pod-name* logs

~~~bash
$ kubectl logs *pod-name* -n *namespace-name*
~~~

## get *pod-name* logs blocked

~~~bash
$ kubectl logs *pod-name* -n *namespace-name*
~~~


Here is a list of the basic Docker commands from this page,
and some related ones if you’d like to explore a bit before moving on.

# Create image using this directory's Dockerfile
docker build -t friendlyname .  
# Run "friendlyname" mapping port 4000 to 80
docker run -p 4000:80 friendlyname  
docker run -d -p 4000:80 friendlyname         # Same thing, but in detached mode
docker container ls                                # List all running containers
docker container ls -a             # List all containers, even those not running
docker container stop <hash>           # Gracefully stop the specified container
docker container kill <hash>         # Force shutdown of the specified container
docker container rm <hash>        # Remove specified container from this machine
docker container rm $(docker container ls -a -q)         # Remove all containers
docker image ls -a                             # List all images on this machine
docker image rm <image id>            # Remove specified image from this machine
docker image rm $(docker image ls -a -q)   # Remove all images from this machine
docker login             # Log in this CLI session using your Docker credentials
docker tag <image> username/repository:tag  # Tag <image> for upload to registry
docker push username/repository:tag            # Upload tagged image to registry
docker run username/repository:tag                   # Run image from a registry