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

## **Start your local Kubernetes cluster instance**

From a terminal with administrator access, run:

```bash
minikube start
```

<span style="color:green">**Expected Output**</span>

```bash
😄  minikube v1.12.1 on Ubuntu 20.04
🎉  minikube 1.18.1 is available! Download it: https://github.com/kubernetes/minikube/releases/tag/v1.18.1
💡  To disable this notice, run: 'minikube config set WantUpdateNotification false'

✨  Using the docker driver based on existing profile
👍  Starting control plane node minikube in cluster minikube
🏃  Updating the running docker "minikube" container ...
🐳  Preparing Kubernetes v1.18.3 on Docker 19.03.2 ...
🔎  Verifying Kubernetes components...
🌟  Enabled addons: default-storageclass, storage-provisioner
🏄  Done! kubectl is now configured to use "minikube"
```

Once done, you’d be able to interact with your cluster and run kubectl commands
> *Initially, some services such as the storage-provisioner, may not yet be in a Running state. This is a normal condition during cluster bring-up and will resolve itself momentarily.*

***

# **For Core Backend Development**

## **Control Plane Backend**

Backend components consist of three microservices

1. GraphQL-Server
2. Authentication-Server
3. MongoDB

## **Steps to run the Control Plane Backend**

1. Install mongoDB on Kubernetes <br/>
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

   Or

   ```bash
   docker run -d -p 27017:27017 --name mongo-service litmuschaos/mongo:4.2.8
   ```

2. Run the Authentication Service <br />
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

3. Run the GraphQL Server <br />
   a. Export the following environment variables

   ```sh
   export DB_SERVER=mongodb://localhost:27017
   export JWT_SECRET=litmus-portal@123
   export PORTAL_ENDPOINT=http://localhost:8080
   export AGENT_SCOPE=cluster
   export SELF_AGENT=true
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

# **For Frontend Development**

## **Install Litmus**

### **Through Helm**

* **Add the Litmus Helm repository**
  ```bash
  helm repo add litmuschaos https://litmuschaos.github.io/litmus-helm/
  helm repo list
  ```

* **Create a Litmus namespace**
  
  ```bash
  kubectl create ns litmus
  ```

* **Install Litmus ChaosCenter**
  
  ```bash
  helm install chaos litmuschaos/litmus --namespace=litmus
  ```

> NOTE: To change the chart version to the latest CI for the local development setup, you can navigate to the `charts/litmus-2-0-0-beta/values.yml` and then modify all the type `tag` to have `ci` as the value.

### **Through Kubernetes Manifest**

* **Applying the 2.8 manifest**

  ```bash
  kubectl apply -f https://raw.githubusercontent.com/litmuschaos/litmus/2.8.0/mkdocs/docs/2.8.0/litmus-2.8.0.yaml
  ```

* **Applying the master stable manifest**

  ```bash
  kubectl apply -f https://raw.githubusercontent.com/litmuschaos/litmus/master/litmus-portal/cluster-k8s-manifest.yml
  ```

***

## **Setup the Portal services locally**

To set up and log in to Litmus Portal locally, expand the available services just created in the `litmus` namespace since the server service contains GraphQL and Authentication required for the portal.

```bash
kubectl get svc -n litmus
```

<span style="color:green">**Expected Output**</span>

```bash
NAME                               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                         AGE
litmusportal-auth-server-service   NodePort    10.99.79.189    <none>        9003:32197/TCP,3030:32060/TCP   24s
litmusportal-frontend-service      NodePort    10.110.66.146   <none>        9091:30230/TCP                  24s
litmusportal-server-service        NodePort    10.108.178.98   <none>        9002:30523/TCP,8000:32662/TCP   24s
mongo-headless-service             ClusterIP   None            <none>        27017/TCP                       24s
mongo-service                      ClusterIP   10.107.168.37   <none>        27017/TCP                       24s
```

Since this is your local setup and you won’t have the prediction environment backing you up, you’d need to configure authentication and GraphQL services manually for the application to simulate the ideal behaviour.

## **Changes in `run.sh` file**

In the `run.sh` file inside `litmus-portal` directory

* comments out line 31-34 `Run script only in litmus portal dir` 
* Update the `VERSION` value to `2.8.0` or relevant version, by default its set to `ci`
* Update `SELF_CLUSTER` key to `SELF_AGENT`


## **Forward these services**
cd into the `litmus-portal` folder inside the cloned repo and use the respective setup to boot each service

* **Mongo DB** (Only if backend services are to be run locally)

  ```bash
  kubectl port-forward svc/mongo-service 27017:27017 -n litmus
  ```
  > We’re using 27017 as our local Mongo DB server

* **Authentication**

  ```bash
  kubectl port-forward svc/litmusportal-auth-server-service 3000:9003 -n litmus
  ```
  or
  ```bash
  bash run.sh auth (For local auth backend)
  ```
  > We’re using 3000 as our local and 9003 as our container authentication server

* **GraphQL**

  ```bash
  kubectl port-forward svc/litmusportal-server-service 8080:9002 -n litmus
  ```
  or
  ```bash
  bash run.sh gql (For local GraphQL backend)
  ```
  > We’re using 8080 as our local and 9002 as our container GraphQL server

***

## **Access the Frontend locally**

If you haven’t already cloned the litmus project do so from the `litmuschaos/litmus` repository

```bash
git clone https://github.com/litmuschaos/litmus.git litmus
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