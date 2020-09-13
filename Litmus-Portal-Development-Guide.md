# Litmus Portal development Guide

## Frontend
```
cd litmus/litmus-portal/frontend

npm install 

npm start
```

## Backend

### Setup MONGODB
```
sudo mkdir -p /mongodata

sudo docker run -it -v /data/db:/mongodata -p 27017:27017 --name mongodb -d mongo
```
### Start servers

1. GQL Server
```
cd litmus/litmus-portal/backend/graphql-server

SERVICE_ADDRESS=localhost:8080 \
  DB_SERVER=localhost:27017 JWT_SECRET=litmus-portal@123 go run server.go
```
let the server run in the background.

2. Authentication Server
```
cd litmus/litmus-portal/backend/auth

DB_SERVER=mongodb://localhost:27017 JWT_SECRET=litmus-portal@123 ADMIN_USERNAME=admin ADMIN_PASSWORD=litmus go run src/main.go
```
let the server run in the background.

### Workflow agent

```
minikube start 

chaos-workflows$ kubectl create ns argo

kubectl apply -f  https://raw.githubusercontent.com/argoproj/argo/stable/manifests/install.yaml -n argo

curl -sLO  https://github.com/argoproj/argo/releases/download/v2.8.0/argo-linux-amd64

sudo chmod +x argo-linux-amd64

sudo mv ./argo-linux-amd64 /usr/local/bin/argo

kubectl apply -f https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/App/nginx.yaml

kubectl apply -f https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/App/service.yaml

kubectl apply -f https://litmuschaos.github.io/litmus/litmus-operator-v1.7.0.yaml

kubectl apply -f https://litmuschaos.github.io/litmus/litmus-admin-rbac.yaml

kubectl apply -f https://hub.litmuschaos.io/api/chaos/master?file=charts/generic/pod-delete/experiment.yaml -n litmus

kubectl apply -f https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/Argo/argo-access.yaml -n litmus

argo submit https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/Argo/argowf-native-pod-delete.yaml -n litmus

kubectl patch svc argo-server -n argo -p '{"spec": {"type": "NodePort"}}'

```

[More detailed documentation](https://docs.google.com/document/d/1y-RkLVDrOJYM2T4RPGBWG8KoNG9T7GAflTGQw1sCXUE/edit)