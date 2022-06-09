# Litmus Developer Guide

> NOTE: This setup has been written on Ubuntu 20.04 LTS as the setup environment, Based on your version of Linux Distribution, some commands may vary, please find the alternative commands accordingly in case there is some syntactical difference that was missed in this guide.

## **Pre-requisites**

Here are a few things you need to make sure are already present in your system before running Litmus locally

* Kubernetes 1.17 or later
* Helm3 or Kubectl
* Node and npm
* Docker
* Golang
* Local Kubernetes Cluster (via minikube, k3s or kind)

***

# **For Core Backend Development**

## **Control Plane Backend**

Backend components consist of three microservices

1. GraphQL-Server
2. Authentication-Server
3. MongoDB

## **Steps to run the Control Plane Backend**

### 1. Install mongoDB on Kubernetes <br/>
   a. Save the following file as mongo.yaml

```yaml
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mongo
  labels:
    app: mongo
spec:
  selector:
    matchLabels:
      component: database
  serviceName: mongo-headless-service
  replicas: 1
  template:
    metadata:
      labels:
        component: database
    spec:
      automountServiceAccountToken: false
      containers:
        - name: mongo
          image: litmuschaos/mongo:4.2.8
          securityContext:
            allowPrivilegeEscalation: false
          args: ["--ipv6"]
          ports:
            - containerPort: 27017
          imagePullPolicy: Always
          volumeMounts:
            - name: mongo-persistent-storage
              mountPath: /data/db
          resources:
            requests:
              memory: "550Mi"
              cpu: "225m"
              ephemeral-storage: "1Gi"
            limits:
              memory: "1Gi"
              cpu: "750m"
              ephemeral-storage: "3Gi"
          env:
            - name: MONGO_INITDB_ROOT_USERNAME
              valueFrom:
                secretKeyRef:
                  name: litmus-portal-admin-secret
                  key: DB_USER
            - name: MONGO_INITDB_ROOT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: litmus-portal-admin-secret
                  key: DB_PASSWORD
  volumeClaimTemplates:
    - metadata:
        name: mongo-persistent-storage
      spec:
        accessModes:
          - ReadWriteOnce
        resources:
          requests:
            storage: 20Gi
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: mongo
  name: mongo-service
spec:
  ports:
    - port: 27017
      targetPort: 27017
  selector:
    component: database
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: mongo
  name: mongo-headless-service
spec:
  clusterIP: None
  ports:
    - port: 27017
      targetPort: 27017
  selector:
    component: database
---
apiVersion: v1
kind: Secret
metadata:
  name: litmus-portal-admin-secret
stringData:
  DB_USER: "admin"
  DB_PASSWORD: "1234"
```

   b. Apply the manifest and port-forward mongo-service
   
   ```bash
   kubectl apply -f mongo.yaml -n default
   kubectl port-forward svc/mongo-service 27017:27017 
   ```

   ### Or, Install mongo on Docker

   a) Apply the following docker command to run the container

   ```bash
   docker run -d -p 27017:27017 --name mongo-service litmuschaos/mongo:4.2.8
   ```
   
   b) Create a mongo user and password eg(default): admin/1234 

   ```bash
   docker exec -ti mongo-service bash
   mongo
   use admin
   db.createUser({ user: "admin" , pwd: "1234", roles: ["userAdminAnyDatabase", "dbAdminAnyDatabase", "readWriteAnyDatabase"]})
   ```


### 2. Run the Authentication Service <br />
   > NOTE: Make sure to run backend services before the frontend. If you haven’t already cloned the litmus project do so from the `litmuschaos/litmus` repository
```bash
git clone https://github.com/litmuschaos/litmus.git litmus --depth 1
```

   a. Export the following environment variables

   ```sh
   export DB_SERVER=mongodb://localhost:27017
   export JWT_SECRET=litmus-portal@123
   export ADMIN_USERNAME=admin
   export ADMIN_PASSWORD=litmus
   export DB_USER=admin
   export DB_PASSWORD=1234
   ```

   b. Run the go application
   
   ```bash
   cd litmus-portal/authentication
   go run api/main.go
   ```

### 3. Run the GraphQL Server <br />
   a. Export the following environment variables

   ```sh
   export DB_SERVER=mongodb://localhost:27017
   export JWT_SECRET=litmus-portal@123
   export PORTAL_ENDPOINT=http://localhost:8080
   export AGENT_SCOPE=cluster
   export SELF_AGENT=false
   export AGENT_NAMESPACE=litmus
   export LITMUS_PORTAL_NAMESPACE=litmus
   export PORTAL_SCOPE=namespace
   export SUBSCRIBER_IMAGE=litmuschaos/litmusportal-subscriber:ci
   export EVENT_TRACKER_IMAGE=litmuschaos/litmusportal-event-tracker:ci
   export CONTAINER_RUNTIME_EXECUTOR=k8sapi
   export ARGO_WORKFLOW_CONTROLLER_IMAGE=argoproj/workflow-controller:v3.2.9
   export ARGO_WORKFLOW_CONTROLLER_IMAGE=argoproj/workflow-controller:v3.2.9
   export ARGO_WORKFLOW_EXECUTOR_IMAGE=argoproj/argoexec:v3.2.9
   export LITMUS_CHAOS_OPERATOR_IMAGE=litmuschaos/chaos-operator:2.7.0
   export LITMUS_CHAOS_RUNNER_IMAGE=litmuschaos/chaos-runner:2.7.0
   export LITMUS_CHAOS_EXPORTER_IMAGE=litmuschaos/chaos-exporter:2.7.0
   export ADMIN_USERNAME=admin
   export ADMIN_PASSWORD=litmus
   export DB_USER=admin
   export DB_PASSWORD=1234
   export VERSION=ci
   export HUB_BRANCH_NAME=v2.7.x
   export AGENT_DEPLOYMENTS="[\"app=chaos-exporter\", \"name=chaos-operator\", \"app=event-tracker\", \"app=workflow-controller\"]"
   export LITMUS_CORE_VERSION="2.7.0"
   ```

   b. Run the go application

   ```bash
   cd litmus-portal/graphql-server
   go run server.go
   ```

## **Steps to run the agent plane backend**

Use [litmusctl](https://github.com/litmuschaos/litmusctl) on the same box/local cluster and connect a ns agent

***

## **Run Frontend locally**

> NOTE: Make sure to run backend services before the frontend.

If you haven’t already cloned the litmus project do so from the `litmuschaos/litmus` repository

```bash
git clone https://github.com/litmuschaos/litmus.git litmus --depth 1
```

Once cloned, navigate to the frontend directory inside the `litmus-portal` folder

```bash
cd litmus/litmus-portal/frontend
```

When you are inside the frontend directory, install all the dependencies and run the project locally.

```bash
npm i & npm start
```

> It’ll prompt you to start the development server at port `3001` or any other port than 3000 since it is already being used by the auth server, simply type `y` and the portal would be up and running in `localhost:3001` or `localhost:<PORT>`

> NOTE: For the local development setup to work correctly, you will have to enable a CORS extension on your browser; you can find one here or use one of your own preference.

Once you are able to see the Login Screen of Litmus use the following default credentials

```bash
Username: admin
Password: litmus
```
<img src="https://i.ibb.co/yhgYLm1/login-75d67e34bdfa757d7647811731e2637a.png" />