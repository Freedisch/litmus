# Chaosthon practical training resources

This page helps the participants of the Chaosthon in getting to know more about LitmusChaos platform and gives some practical training in creating and executing chaos workflows or scenarios.

## Stage1

**Create your own Chaos Control Plane**

* Sign up at https://litmuschaos.cloud

![image](https://user-images.githubusercontent.com/19591814/171148024-de6ba744-77af-49ce-889d-1e964663cadc.png)

* Browse and learn around managing chaos workflows. Read the documentation at docs.litmuschaos.io


**Join the ChaosThon project**

* Once you signup at litmuschaos.cloud, you will be invited to join a project. The invitation will be from umasankar.mukkara@harness.io. The invitation is visible on the top right of the Settings --> Teaming page.
![image](https://user-images.githubusercontent.com/19591814/171151503-5800d2a3-fc48-4a71-84f5-4da482e72fdc.png)

* Click or scroll down on the teaming page to find the invitation. You must be invited as an editor. Accept the invitation.
![image](https://user-images.githubusercontent.com/19591814/171163779-a6b6c6c0-1e12-45bc-994d-b52e90413e07.png)
 
* Once the invitation is accepted, you will have access to the ChaosThon project on your project bar on your top left screen of the chaos center. Click on the project and select "Umasankar90499's project"
![image](https://user-images.githubusercontent.com/19591814/171165922-a6aedcaa-1ee0-44ab-80bd-1e87d54b2b33.png)


***
## Stage 2

It is now assumed that you have access to "Umasankar90499's project" on litmuschaos.cloud. When you are in this project, you will be doing actions as "Umasankar". It is important to keep the environment usable by other participants of Chaosthon. Do not delete the workflows that are not created by you.
When you browse the project, you see a fully configured target environment. You will see chaos workflows that ran against an agent called "eks-common". You will be running a demo chaos workflow against an application called sock-shop catalog. 

### Sample application details

The sample is called sock-shop, which is installed in a namespace called "sock-shop" in a eks cluster, which is connected to your project through the eks-common agents. The details of the agents can be seen at the ChaosAgents section.

* Sockshop application is setup at http://a5d3f939433a1425e8bed9f7cf632ed4-533721864.us-east-2.elb.amazonaws.com/category.html . This is a e-commerce application where various types of socks are sold. 
* Catalog is a micro service that is accessible on this application
![image](https://user-images.githubusercontent.com/19591814/171152601-5796f26d-90d1-4de4-ba67-71a67f2fe86a.png)

### Creating a pod-delete chaos on catalog service

You can create a new chaos workflow that deletes pod belonging to the catalog service. A pre-defined chaos workflow is made available for you in a ChaosHub called "Accenture-ChaosHub". Follow the below steps.

1. Browse the catalog on the sock-sock website and see if it is working before chaos is applied.
2. Create a new Chaos Workflow from Litmus Workflows tab.

* Browse the catalog application when chaos is in play

* Observe the results 

## Stage 3
* Create a new chaos workflow against the same catalog experiment
In this case, we create a Pod Delete against front-end website or any other micro service in the sock-shop namespace.





