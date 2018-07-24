# paprika-api
simple golang api

# Running
1. Create `.env` using `.env.example`
2. `go get ./`
3. `go run main.go`

# Building
1. Create `.env` using `.env.example`
2. `go get ./`
3. `go build -o paprika-api ./`
4. `./paprika-api`

# Testing
1. Create `.env` using `.env.example`
2. `go get ./`
3. `go test ./test/*`
  * Tests are modularized to test individual components of the api, i.e. server, controllers, db, etc.

# Local Kubernetes
Debug prereqs:
1. [Minikube](https://kubernetes.io/docs/setup/minikube/)
2. [Skaffold](https://github.com/GoogleContainerTools/skaffold#installation)
3. [Kustomize](https://github.com/kubernetes-sigs/kustomize/blob/master/INSTALL.md)
4. [Modd](https://github.com/cortesi/modd#install)
5. [Squash](https://github.com/solo-io/squash/tree/master/docs/install#installing-squash)
6. [Squash VS Code Extension](https://marketplace.visualstudio.com/items?itemName=ilevine.squash)

Debug steps:
1. `minikube start`
2. `kubectl proxy` (if deubgging with Squash)
  * Add `"vs-squash.squash-server-url": "http://localhost:8001/api/v1/namespaces/squash/services/squash-server:http-squash-api/proxy/api/v2"` to VSCode settings
3. `skaffold dev`
4. `modd`

API is now available at `http://192.168.99.100:31000/`. Any changes made to any `.go` files in the directory will trigger Modd to rebuild the go app which then triggers Skaffold to build, push, and desply app.

By using Squash, you can connect to the pod and debug the API.
