# Litmus Portal development Guide

This is a detailed guide towards setting up the litmus development environment locally, if you are only interested in getting the portal up and running, you can skip “Nginx traffic characteristics during a non-chaotic benchmark run” and “Observe the Nginx benchmark results” sub-sections in the Wokflow Agent section


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

DEPLOYER_IMAGE=itmuschaos/litmusportal-self-deployer:ci SUBSCRIBER_IMAGE=litmuschaos/litmusportal-subscriber:ci SERVICE_ADDRESS=localhost:8080 DB_SERVER=localhost:27017 JWT_SECRET=litmus-portal@123 go run server.go
```
let the server run in the background.

2. Authentication Server
```
cd litmus/litmus-portal/authentication

DB_SERVER=mongodb://localhost:27017 JWT_SECRET=litmus-portal@123 ADMIN_USERNAME=admin ADMIN_PASSWORD=litmus go run src/main.go
```
let the server run in the background.

## Workflow agent

Start a minikube instance on your local machine:
```
minikube start 
```

### Install Argo Workflow Infrastructure

- Create an Argo namespace:

```
root@demo:~/chaos-workflows# kubectl create ns argo
namespace/argo created
```

- Create the CRDs, workflow controller deployment with associated RBAC:

```
kubectl apply -f  https://raw.githubusercontent.com/argoproj/argo/stable/manifests/install.yaml -n argo

customresourcedefinition.apiextensions.k8s.io/clusterworkflowtemplates.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/cronworkflows.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/workflows.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/workflowtemplates.argoproj.io created
serviceaccount/argo created
serviceaccount/argo-server created
role.rbac.authorization.k8s.io/argo-role created
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-admin configured
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-edit configured
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-view configured
clusterrole.rbac.authorization.k8s.io/argo-cluster-role configured
clusterrole.rbac.authorization.k8s.io/argo-server-cluster-role configured
rolebinding.rbac.authorization.k8s.io/argo-binding created
clusterrolebinding.rbac.authorization.k8s.io/argo-binding unchanged
clusterrolebinding.rbac.authorization.k8s.io/argo-server-binding unchanged
configmap/workflow-controller-configmap created
service/argo-server created
service/workflow-controller-metrics created
deployment.apps/argo-server created
deployment.apps/workflow-controller created
```

- Install the Argo CLI on the test harness machine (where the kubeconfig is available):

```
root@demo:~# curl -sLO https://github.com/argoproj/argo/releases/download/v2.8.0/argo-linux-amd64

root@demo:~# chmod +x argo-linux-amd64

root@demo:~# mv ./argo-linux-amd64 /usr/local/bin/argo

root@demo:~# argo version
argo: v2.8.0
BuildDate: 2020-05-11T22:55:16Z
GitCommit: 8f696174746ed01b9bf1941ad03da62d312df641
GitTreeState: clean
GitTag: v2.8.0
GoVersion: go1.13.4
Compiler: gc
Platform: linux/amd64
```

### Install a Sample Application: Nginx

- Install a simple multi-replica stateless Nginx deployment with service exposed over nodeport:

```
root@demo:~# kubectl apply -f https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/App/nginx.yaml

deployment.extensions/nginx created
root@demo:~# kubectl apply -f https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/App/service.yaml 
service/nginx created
```

### Install Litmus Infrastructure

- Apply the LitmusChaos Operator manifest:

```
kubectl apply -f https://litmuschaos.github.io/litmus/litmus-operator-v1.7.0.yaml
```

- Install the litmus-admin service account to be used by the chaos-operator while executing the experiment:

```
kubectl apply -f https://litmuschaos.github.io/litmus/litmus-admin-rbac.yaml
```

- Install the Chaos experiment of choice (in this example, we pick a pod-delete experiment):

```
kubectl apply -f https://hub.litmuschaos.io/api/chaos/master?file=charts/generic/pod-delete/experiment.yaml -n litmus
```

### Create the Argo Access ServiceAccount

- Create the service account and associated RBAC, which will be used by the Argo workflow controller to execute the actions specified in the workflow. In our case, this corresponds to the launch of the Nginx benchmark job and creating the chaosengine to trigger the pod-delete chaos action. In our example, we place it in the namespace where the litmus chaos resources reside, i.e., litmus.

```
root@demo:~# kubectl apply -f https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/Argo/argo-access.yaml -n litmus

serviceaccount/argo-chaos created
clusterrole.rbac.authorization.k8s.io/chaos-cluster-role created
clusterrolebinding.rbac.authorization.k8s.io/chaos-cluster-role-binding created

```

### Nginx traffic characteristics during a non-chaotic benchmark run

Before proceeding with the chaos workflows, let us first look at how the benchmark run performs under normal circumstances & what are the properties of note.

To achieve this:

- Let us run a simple Kubernetes job that internally executes an apache-bench test on the Nginx service with a standard input of 10000000 requests over a 300s period.

```
root@demo:~# kubectl create -f https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/App/nginx-bench.yaml

job.batch/nginx-bench-c9m42 created
```

- Observe the output post the 5 min duration & note the failed request count. Usually, it is 0, i.e., there was no disruption in Nginx traffic.

```
root@demo:~# kubectl logs -f nginx-bench-zq689-6mnrm

2020/06/23 01:42:29 Running: ab -r -c10 -t300 -n 10000000 http://nginx.default.svc.cluster.local:80/
2020/06/23 01:47:35 This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking nginx.default.svc.cluster.local (be patient)
Finished 808584 requests


Server Software:        nginx/1.19.0
Server Hostname:        nginx.default.svc.cluster.local
Server Port:            80

Document Path:          /
Document Length:        612 bytes

Concurrency Level:      10
Time taken for tests:   300.001 seconds
Complete requests:      808584
Failed requests:        0
Total transferred:      683259395 bytes
HTML transferred:       494857692 bytes
Requests per second:    2695.27 [#/sec] (mean)
Time per request:       3.710 [ms] (mean)
Time per request:       0.371 [ms] (mean, across all concurrent requests)
Transfer rate:          2224.14 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.7      0      25
Processing:     0    3   2.0      3      28
Waiting:        0    3   1.9      2      28
Total:          0    4   2.2      3      33
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      5
  80%      5
  90%      7
  95%      8
  98%      9
  99%     11
 100%     33 (longest request)
```

In the next step, we shall execute a chaos workflow that runs the same benchmark job while a random pod-delete (Nginx replica failure) occurs and observe the degradation in the attributes we have noted: failed_requests.

### Create the Chaos Workflow

Applying the workflow manifest performs the following actions in parallel:

- Starts an Nginx benchmark job for the specified duration (300s)

- Triggers a random pod-kill of the Nginx replica by creating the chaosengine CR. Cleans up after chaos.

```
root@demo:~# argo submit https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/Argo/argowf-native-pod-delete.yaml -n litmus
Name:                argowf-chaos-sl2cn
Namespace:           litmus
ServiceAccount:      argo-chaos
Status:              Pending
Created:             Fri May 15 15:31:45 +0000 (now)
Parameters:
  appNamespace:      default
  adminModeNamespace: litmus
  appLabel:          nginx
```

### Visualize the Chaos Workflow

You can visualize the progress of the chaos workflow via the Argo UI. Convert the argo-server service to type NodePort & view the dashboard at: `https://<node-ip>:<nodeport>`

```
root@demo:~# kubectl patch svc argo-server -n argo -p '{"spec": {"type": "NodePort"}}'
service/argo-server patched
```
![Argo visualisation](https://res.cloudinary.com/practicaldev/image/fetch/s--MkzlakSi--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://user-images.githubusercontent.com/21166217/82098260-38738b00-9722-11ea-81b4-b3c466a60080.png)

### Observe the Nginx benchmark results

Observing the Nginx benchmark results over 300s with a single random pod kill shows an increased count of failed requests.

```
root@demo:~# kubectl logs -f nginx-bench-7pnvv

2020/06/23 07:00:34 Running: ab -r -c10 -t300 -n 10000000 http://nginx.default.svc.cluster.local:80/
2020/06/23 07:05:37 This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking nginx.default.svc.cluster.local (be patient)
Finished 802719 requests


Server Software:        nginx/1.19.0
Server Hostname:        nginx.default.svc.cluster.local
Server Port:            80

Document Path:          /
Document Length:        612 bytes

Concurrency Level:      10
Time taken for tests:   300.000 seconds
Complete requests:      802719
Failed requests:        866
   (Connect: 0, Receive: 289, Length: 289, Exceptions: 288)
Total transferred:      678053350 bytes
HTML transferred:       491087160 bytes
Requests per second:    2675.73 [#/sec] (mean)
Time per request:       3.737 [ms] (mean)
Time per request:       0.374 [ms] (mean, across all concurrent requests)
Transfer rate:          2207.20 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0  11.3      0    3044
Processing:     0    3  57.2      3   16198
Waiting:        0    3  54.2      2   16198
Total:          0    4  58.3      3   16199

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      4
  80%      5
  90%      6
  95%      7
  98%      9
  99%     11
 100%  16199 (longest request)
```

Further iterations of these tests with increased pod-kill instances over the benchmark period or an increased kill count (i.e., number of replicas killed at a time) can give more insights about the behavior of the service, in turn leading us to the mitigation procedures.

**Note**: To test with different variables, edit the ChaosEngine spec in the workflow manifest before re-submission.

### Get CLUSTER_ID & ACCESS_KEY

After setting up the project name, user name, user email & Password:

Press (Ctrl + Shift + I) to inspect, and switch to the console tab.

Expand the next state object (any of the mentioned ones) and copy the selectedProjectID value

Open localhost:8080 on your browser to access the GQL Playground and type in the below provided query & pass some dummy variables as shown below, with the exception of the “project_id” field which should be the same as the one you’ve just retrieved in the previous step

Below you can see a reference implementation of the GQL Query passing, and the actual query below it:

![GQL_1](https://user-images.githubusercontent.com/40641427/93286585-3b2bdd00-f7f5-11ea-9f1a-af36b3e837f4.png)

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
Open a browser window with the URL: `localhost:8080/file/<jwt_key>.yaml`

### Retrieve the Cluster ID and key

Scroll to the very bottom of this page and identify two fields: “CID” and “KEY”,
Copy the strings in the “value” section of these respectively. The “value” string corresponding to the “CID” field will act as your local cluster ID for minikube from now on, it’s a good idea to save it’s value somewhere safe in your system for later usage.

You can see the fields being referred above at the very bottom of the below provided screenshot as a reference:
![C_ID_and_AccessKey](https://user-images.githubusercontent.com/40641427/93287087-44697980-f7f6-11ea-9f4b-1e43469a6b5f.png)


### Confirm Cluster ID and Access Key pair

Go back to the GraphQL playground at [http://localhost:8080/](http://localhost:8080/) and create a new tab
enter the following query with the respective query variables.

**NOTE: The “cluster_id” & “access_key” fields will have the strings you have retrieved above as “CID” and “KEY”  values respectively.**

Reference Screenshot with the query have been provided below:

![GQL_2](https://user-images.githubusercontent.com/40641427/93288671-10905300-f7fa-11ea-9ec7-954b1f7bb6e4.png)

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

`cd /litmus/litmus-portal/cluster-agents/subscriber`

Enter the following command with your CID = <cluster ID> string in field “cluster_id” and the KEY=<Cluster Key> as the string in the “newClusterKey” field.


` CID=<CLUSTER_ID> KEY=<ACCESS_KEY> GQL_SERVER=http://localhost:8080/query go run subscriber.go -kubeconfig ~/.kube/config `

## Generate workflow events:

On a separate terminal instance to get the workflow event

` argo submit https://raw.githubusercontent.com/litmuschaos/chaos-workflows/master/Argo/argowf-native-pod-delete.yaml -n litmus`

The logs should show up in the workflow agent and also in the mongodb. If there’s a subscriber running for workflow events it should also receive these.

Once it is done, you should be able to see your workflow on the local website UI in which we had logged on earlier.