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

1. Browse the [catalog](http://a5d3f939433a1425e8bed9f7cf632ed4-533721864.us-east-2.elb.amazonaws.com/category.html) on the sock-sock website and see if it is working before chaos is applied.

2. Create a new Chaos Workflow from Litmus Workflows tab.
![image](https://user-images.githubusercontent.com/19591814/171174523-0fe9e27e-84c0-4b4e-98ec-17b12168ca8e.png)

3. Choose eks-common as the agent and click next
![image](https://user-images.githubusercontent.com/19591814/171179449-ac405d9b-9c1d-438a-9d4d-4e1c72c93f4d.png)

4. Choose Accenture-ChaosHub in the list of chaos hubs.
![image](https://user-images.githubusercontent.com/19591814/171179533-8b281c0c-eacc-4313-9403-0563d0d03b9c.png)

5. Accenture-ChaosHub has two pre-defined chaos workflows. Choose accenture-demo1
![image](https://user-images.githubusercontent.com/19591814/171179643-9132591c-d6d9-4e2d-ab0b-7bdac41898c9.png)

6. Name the chaos workflow with something that contains your name, so that it can be identified. Many participants are sharing this environment. 
![image](https://user-images.githubusercontent.com/19591814/171179820-e309b8a2-bd54-4a8b-989f-b351cd59be06.png)

7. The experiment is setup for you already. If you wish browse the experiment and probe and learn how it is configured. Do NOT change any settings as the experiment is pre-configured to run against catalog service.
![image](https://user-images.githubusercontent.com/19591814/171180068-f1885113-b669-48ff-8283-33208785cf33.png)

8. Click next on the resilience score screen. There is only one experiment, no need to set weight.
![image](https://user-images.githubusercontent.com/19591814/171180207-db8daecf-720e-4619-a9a5-9a3dd3b4f0cc.png)

9. In the schedule screen, "Schedule now" is selected by default, click Next.
![image](https://user-images.githubusercontent.com/19591814/171180339-04c43760-2048-48ed-b24d-e9b95ddefb8c.png)

10. The final screen is a summary screen. Browse the YAML file to learn how the chaos workflow is constructed. Click Finish button.
![image](https://user-images.githubusercontent.com/19591814/171180513-8bb55ccf-3b84-436f-884e-fae541c68cb7.png)

11. Click Go to Workflow.
![image](https://user-images.githubusercontent.com/19591814/171180603-5cc2a506-da04-44e3-844b-a35091717204.png)

12. You will see that your chaos workflow is running and click on your workflow name to browse the running workflow.
![image](https://user-images.githubusercontent.com/19591814/171180766-905be774-5c18-4e57-a395-a06a5dfb0311.png)

13. The workflow starts with setting up the experiments environment before running the actual experiments. At this point of time your sock-shop catalog must be available. [Browse](http://a5d3f939433a1425e8bed9f7cf632ed4-533721864.us-east-2.elb.amazonaws.com/category.html) and check it is working. 
![image](https://user-images.githubusercontent.com/19591814/171180901-4a083eff-1bd8-4609-9bab-9a03cad25f41.png)

14. After a minute or so, pod-delete will start running. Few seconds into it, the catalog service may not be available, as there is no redundancy to that service. Till the pod is back up, the catalog service should remain unavailable. 
![image](https://user-images.githubusercontent.com/19591814/171181313-aa8f755a-8dba-42d2-9d67-ef17e9948e01.png)

15. Browse the catalog service and see it as unavailable. This means, that you are able to run the pod-delete chaos against the catalog service.
![image](https://user-images.githubusercontent.com/19591814/171181521-ba5a13f1-3b92-4103-80e6-6fb1245c17f1.png)

16. On the workflow browser, click on the pod-delete experiment to see the logs console.
![image](https://user-images.githubusercontent.com/19591814/171181625-0daafd9f-cef7-44bf-ae0d-5e6567469fe3.png)

17. Click on the Chaos Results tab to see the chaos workflow results.
![image](https://user-images.githubusercontent.com/19591814/171181727-59eeb7dd-5456-4fe3-a614-3e620a0c919d.png)

18. The workflow will run for a few minutes. When pod-delete chaos experiment is completed, the results can be seen in the Chaos Results tab. At the end Revent-Chaos step is run, this will run any chaos logs and results to keep the system at it's original state. Note that, after the revert chaos step, logs and results are not available as part of the workflow browser screen.

![image](https://user-images.githubusercontent.com/19591814/171182176-2b080546-858d-42f5-b61c-c239fb7a8721.png)

19. By now, the catalog pod must be back and [service](http://a5d3f939433a1425e8bed9f7cf632ed4-533721864.us-east-2.elb.amazonaws.com/category.html) must be available.

This is the end of stage 2. Feel free to browse various features in the chaos center. The settings section provides various capabilities around teaming and GitOps.


# Become part of the Litmus community
* If you like the Litmus project, you may want to become one of the project stargazers on GitHub. Provide a star [here](https://github.com/litmuschaos/litmus/stargazers) if you like Litmus.
* Join Litmus [slack channel](https://slack.litmuschaos.io) . Litmus community is at #litmus channel in Kubernetes slack workspace. 

***


## Stage 3
* Create a new chaos workflow against the same catalog experiment using the workflow wizard.
In this case, we create a Pod Delete against front-end website or any other micro service in the sock-shop namespace.

_More details are coming soon_





