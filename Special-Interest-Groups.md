# Litmus SIGs 

(as of 10/08/2020)

As the LitmusChaos project & community continue to scale, there has been a great need for area-wise focus and better roadmap/product management around how the components grow, especially with the varied nature of Litmus usage by different personas (Developers, QA, CI/CD & others DevOps functions focused engineers to SREs). Also, the integration points (with other tools, frameworks) exist across several components, which makes it important to manage these better and ensure that the monthly releases are balanced in terms of the progress made across different areas. There also has been growing interest in the community around contributions, with strong preferences for the different areas. 

As maintainers, this prompted us to come up with Special Interest Groups (inspired of course, by CNCF & Kubernetes communities) where community members align themselves around components/areas they are most interested - in terms of contributions or integration and thereby ensure the overall growth of the project. 

By definition, _SIGs oversee and coordinate the interests and needs of end-users and the project in a particular area_.

The SIGs at this point are categorized on the core functions in LitmusChaos, while some SIGs do cut across these in terms of their scope (SIG-Testing, SIG-Documentation, for example). Each SIG has:

- A dedicated charter with well-specified Goals & Non-Goals
- One or more SIG chairs/leads  
- Respective Teams in the Github Org with a set of repositories under its purview
- A meeting cadence that will help in bringing the members together to discuss and share updates on actions furthering 
  the SIG's objective in each monthly release

Some of the planned (& informally operational) SIGs with links to their charters are described below: 

- [SIG-Integrations]()
- [SIG-Observability]()
- [SIG-Deployment]()
- [SIG-CI]()
- [SIG-Testing]()
- [SIG-Documentation]()
- [SIG-Orchestration]()
 
# SIG Details 

This section describes each SIG in detail. 

## SIG-Integrations Charter

### Note

_This SIG is expected to comprise, amongst others, of_:

 _(a) Members who are using Litmus to orchestrate their existing chaos experiments (using Pumba, Chaostoolkit, Powerfulseal, etc., or any other custom chaos logic)_

_(b) Members who are using other abstractions to run litmus experiments, for ex: Argo Workflows_


### Goals

- Contribute requirements and maintain the plugin code supporting the BYOC (Bring-Your-Own-Chaos) model of execution in LitmusChaos (co-maintain this with SIG-Orchestration) 

- Maintain the respective Chaos Experiment business logic & (CRs)/Charts running custom chaoslib on the ChaosHub 

- Establish the patterns in providing the monitoring & reporting hooks consistent with native litmus experiments, for the respective chaos integrations 

- Contribute to the CI/E2E processes around the chaos integrations & plugin code to maintain quality 

- Maintain developer & user documentation, demo artefacts/examples associated with usage of the respective chaos integrations 

- Provide well-defined governance and contribution guidelines for the GitHub repositories under the purview of SIG-Integration

- Support queries in the community (slack/email) over questions and issues around the chaos integrations

### Non-Goals

*Note: Can be picked optionally if the SIG members are interested, as part of the scope/representative of a different SIG*

- Replicate the existing experiment suites with the chaos tool/integration

### Chairs/Leads

- [Sumit Nagal](https://github.com/sumitnagal)
- [Vijay Thomas](https://github.com/vijayto)

### Repositories

- [litmuschaos/litmus-python](https://github.com/litmuschaos/litmus-python)
- [litmuschaos/chaos-workflows](https://github.com/litmuschaos/chaos-workflows)
- [litmuschaos/chaos-runner](https://github.com/litmuschaos/chaos-runner) [co-maintain with SIG-Orchestration]
- [litmuschaos/chaos-charts](https://github.com/litmuschaos/chaos-charts) [co-maintain with SIG-Deployment & SIG-Chaos]

### SIG Meeting Cadence

- Meeting Notes:
- Meeting Cadence:
- Zoom Link:
- SIG-Documentation Meet Recordings: