# Litmus Developer Guide

> NOTE: This setup has been written on Ubuntu 20.04 LTS as the setup environment, Based on your version of Linux Distribution, some commands may vary, please find the alternative commands accordingly in case there is some syntactical difference that was missed in this guide.

## **Pre-requisites**

Here are a few things you need to make sure are already present in your system before running Litmus locally

* Kubernetes 1.15 or higher
* Helm3 or Kubectl
* Node and npm
* Docker
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

## **Install Litmus**

### **Through Helm**

* **Create a Litmus namespace**
  
  ```bash
  kubectl create ns litmus
  ```

* **Add the Litmus Helm Chart**

  ```bash
  git clone https://github.com/litmuschaos/litmus-helm
  cd litmus-helm
  ```

* **Install Litmus**
  
  ```bash
  helm install litmuschaos --namespace litmus ./charts/litmus-2-0-0-beta/
  ```

> NOTE: To change the chart version to the latest CI for the local development setup, you can navigate to the `charts/litmus-2-0-0-beta/values.yml` and then modify all the type `tag` to have `ci` as the value.

### **Through Kubernetes Manifest**

* **Applying the latest beta manifest**

  ```bash
  kubectl apply -f https://litmuschaos.github.io/litmus/2.0.0-Beta/litmus-2.0.0-Beta.yaml
  ```

* **Applying the master stable manifest**

  ```bash
  kubectl apply -f https://raw.githubusercontent.com/litmuschaos/litmus/master/litmus-portal/cluster-k8s-manifest.yml
  ```

***

## **Setup the Portal services locally**

To set up and log in to Litmus Portal locally, expand the available services just created and look for the `litmusportal-server-service` service in the `litmus` namespace s since the server service contains GraphQL and Authentication required for the portal.

```bash
kubectl get svc -n litmus
```

<span style="color:green">**Expected Output**</span>

```bash
NAME                            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                         AGE
chaos-litmus-portal-mongo       ClusterIP   10.104.107.117   <none>        27017/TCP                       2m
litmusportal-frontend-service   NodePort    10.101.81.70     <none>        9091:30385/TCP                  2m
litmusportal-server-service     NodePort    10.108.151.79    <none>        9002:32456/TCP,9003:31160/TCP   2m
```

Since this is your local setup and you won’t have the prediction environment backing you up, you’d need to configure authentication and GraphQL services manually for the application to simulate the ideal behaviour.

## **To enable**

* **Authentication**

  ```bash
  kubectl port-forward svc/litmusportal-server-service 3000:9003 -n litmus
  ```
  > We’re using 3000 as our local authentication server

* **GraphQL**

  ```bash
  kubectl port-forward svc/litmusportal-server-service 8080:9002 -n litmus
  ```
  > We’re using 8080 as our local GraphQL server

***

## **Access the Portal locally**

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
<img src="https://litmusdocs-beta.netlify.app/assets/images/login-53d18e01dbbc518c5e0fdd8ca5fb9500.png" />