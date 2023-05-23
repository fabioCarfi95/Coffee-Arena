## Build application binary.
.PHONY: build
build:
   go build -race -ldflags "-s -w" -o bin/main main.go
 
## Build application binary and run it.
.PHONY: run
run: build
   bin/main
 
## Send a dummy request to the local server.
.PHONY: local-test
local-test:
   curl --request GET 0.0.0.0:10000/
 
## -----------------------------------------------------------------------------------
 
## Build, tag and push application image to registry then clean up.
.PHONY: push
push:
   docker build -t you/sport:latest -f ./build/docker/dev/Dockerfile .
   docker push you/sport:latest
   docker rmi you/sport:latest
   docker system prune --volumes --force
 
## Apply secret and deploy application to kubernetes cluster.
.PHONY: deploy
deploy: secret
   kubectl apply -f build/deploy/k8s/secret.yaml
   kubectl apply -f build/deploy/k8s/deployment.yaml
   kubectl apply -f build/deploy/k8s/service.yaml
 
## Send a dummy request to the exposed pod on kubernetes.
.PHONY: k8s-test
k8s-test:
   curl --request GET $(shell minikube service sport-service --url)/api/v1/team
#kubectl exec coffee-arena-deployment-75466cff9c-79mx9 -- printenv