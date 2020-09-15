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
cd litmus/litmus-portal/graphql-server

SERVICE_ADDRESS=localhost:8080 DB_SERVER=localhost:27017 JWT_SECRET=litmus-portal@123 go run server.go
```
let the server run in the background.

2. Authentication Server
```
cd litmus/litmus-portal/authentication

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

### Get CLUSTER_ID & ACCESS_KEY

After setting up the project name, user name, user email & Password:

Press (Ctrl + Shift + I) to inspect, switch to the console tab.

Expand the next state object (any of the mentioned ones) and copy the selectedProjectID

Open localhost:8080 on your browser to access the GQL Playground and type in the below provided query & pass some dummy variables as shown below, with the exception of the “project_id” field which should be the same as the one you’ve just retrieved in the previous step:

```
Query:
mutation RegCluster($data: ClusterInput!){
  userClusterReg(clusterInput: $data)
}

Query variables:
{
  "data": {
    "cluster_name": "Helo",
    "project_id": "<selectedProjectID>",
    "description": "Cluster test",
    "platform_name": "GKE",
    "cluster_type":  "internal"
  }
}
```

You will get the JWT key on the side panel now.
Open a browser window with the URL: localhost:8080/file/<jwt_key>.yaml

### Retrieve the Cluster ID and key

Scroll to the very bottom of this page and identify two fields: “CID” and “KEY”,
Copy the strings in the “value” section of these respectively. The “value” string corresponding to the “CID” field will act as your local cluster ID for minikube from now on, it’s a good idea to save it’s value somewhere safe in your system for later usage.

### Confirm Cluster ID and Access Key pair

Go back to the GraphQL playground at [http://localhost:8080/](http://localhost:8080/) and create a new tab
enter the following query with the respective query variables.

**NOTE: The “cluster_id” & “access_key” fields will have the strings you have retrieved above as “CID” and “KEY”  values respectively.**

```
Query:
mutation ConfirmCluster($data: ClusterIdentity!){
  clusterConfirm(identity: $data){
    isClusterConfirmed
    cluster_id
    newClusterKey
  }
}
Query Variables:
{
  "data": {
    "cluster_id": "<Your CID value>",
    "access_key": "<Your Key Value>"
  }
}
```
copy the “newClusterKey” string value. This will now be used to run your subscriber locally

## Start Subscriber

Create litmus namespace if it’s not created: 

`kubectl create ns litmus`

Change the present directory to the subscriber folder:

`cd litmus/litmus-portal/backend/subscriber`

Enter the following command with your CID = <cluster ID> string in field “cluster_id” and the KEY=<Cluster Key> as the string in the “newClusterKey” field.


` CID=<CLUSTER_ID> KEY=<ACCESS_KEY> GQL_SERVER=http://localhost:8080/query go run subscriber.go -kubeconfig ~/.kube/config `

## Generate workflow events:

On a separate terminal instance to get the workflow event

` argo submit https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/Argo/argowf-native-pod-delete.yaml -n litmus`

The logs should show up in the workflow agent and also in the mongodb. If there’s a subscriber running for workflow events it should also receive these.

Once it is done, you should be able to see your workflow on the local website UI in which we had logged on earlier.