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

-----------------

## SIG-Integrations 

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

------------------


## SIG-Observability

### Note

_This SIG is expected to comprise, amongst others, of_:

 _(a) Members who are interested in improving the Litmus framework to enable observation & visualization of chaos via metrics, events,  logs, notifications & alerts )_

_(b) Members who are interested in developing plugins for integrating Litmus with other standard observability frameworks/stacks (includes agents, exporters, dashboards, etc.,_

### Goals

- Contribute requirements, use-cases to improve the notion of experiment success/failure. For Ex: (a) HTTP/Command probes that users can employ in a declarative manner to gather more decision points before calling out experiment verdicts (b) Allow for embedding contextual data into ChaosResult etc., 

- Contribute requirements, use-cases & code to improve the monitoring hooks in Litmus - such as events and metrics, i.e., the additions that need to be made to the existing set. Involves proposing and contributing enhancements to the chaos exporter logic & also evaluating, adopting (fork? ) & enhancing (if-applicable) existing OSS tools to achieve improved monitoring.

- Recommend and maintain (revisit periodically) the artifacts (Kubernetes manifests) for standard observability stacks (For ex: Elasticsearch-Fluentd-Kibana / Loki-Promtail-Grafana) with configs setup for litmus. 

- Discuss and contribute requirements around how Litmus can add value in terms of application-specific monitoring. For Ex: Newer off-the-shelf grafana dashboards that overlay litmus metrics/events against the app metrics. Say, for standard stateful apps like Kafka, Cassandra or other applications listed on the ChaosHub. 

- Contribute to the CI/E2E processes around the monitoring integrations & plugin code to maintain quality 

- Maintain developer & user documentation, demo artifacts/examples associated with usage of the respective integrations 

- Provide well-defined governance and contribution guidelines for the GitHub repositories under the purview of SIG-Observability

- Support queries in the community (slack/email) over questions and issues around chaos observability 

### Non-Goals

- Replace human decision making around experiments (the objective here is to provide the right observability aids to help the user injecting chaos in making better calls) 

### Chairs/Leads

- TBD
- TBD

### Repositories

- [App Monitoring Dashboards](https://github.com/litmuschaos/litmus/tree/master/demo/sample-applications)
- [litmuschaos/chaos-observability](https://github.com/litmuschaos/chaos-observability)
- [litmuschaos/chaos-exporter](https://github.com/litmuschaos/chaos-exporter)

### SIG Meeting Cadence

- Meeting Notes:
- Meeting Cadence:
- Zoom Link:
- SIG-Documentation Meet Recordings

----------------------

## SIG-Deployment 

### Goals

- Facilitate an easy & flexible deployment experience for LitmusChaos via standard frameworks (manifests, helm charts) 
- Enable effective lifecycle management (including seamless upgrades) for Litmus framework components (Operators, CR/CRDs) 
- Identify and maintain the standard application delivery entry points for Litmus manifests 

  - LitmusChaos Helm Repo
  - Helm Hub
  - JFrog ChartCenter
  - OperatorHub
  - CNCF Artifact Hub
--

- Revisit the deployment artefacts in every release and review/approve enhancements to the deployment bundles by the other (functionality) SIGs as part of the release

- Contribute documentation in the form of READMEs & usage guides in the litmus-docs around deployment methods. 

- Provide well-defined governance and contribution guidelines for the GitHub repositories under the purview of SIG-Deployment

- Support queries in the community (slack/email) over questions around deployment practices and issues. 

### Non-Goals 

*Note: Can be picked optionally if the SIG members are interested, as part of the scope/representative of a different SIG* 

- Submit PRs to add new product capabilities into the deployment bundles

### Chairs/Leads:

- [Maria Kotlyarevskaya](https://github.com/Jasstkn)
- [Karthik Satchitanand](https://github.com/ksatchit)

### Repositories

- [litmuschaos/litmus](https://github.com/litmuschaos/litmus/tree/master/docs)
- [litmuschaos/litmus-helm](https://github.com/litmuschaos/litmus-helm)
- [litmuschaos/chaos-charts](https://github.com/litmuschaos/chaos-charts) [co-maintain with SIG-Chaos & SIG-Integrations]

### SIG Meeting Cadence

- Meeting Notes: 
- Meeting Cadence: 
- Zoom Link: 
- SIG-Deployment Meet Recordings: 

-----------------------

## SIG-CI

### Note

_This SIG is expected to comprise, amongst others, of_:

_(a) Members who are using litmus experiments as part of their CI pipelines, for ex: via Gitlab Remote Templates, Github Actions, Jenkins Job Templates, etc.,_

### Goals

- Contribute requirements & drive implementation around hooks needed in litmus chaos infrastructure to execute experiments as CI jobs for the respective CI platforms (interface with the SIG-Chaos & SIG-Orchestration to achieve this)

- Help in the creation & maintenance of reusable job templates for usage in respective CI platforms, with a well-defined release cadence 

- Implement and maintain the pipeline artfeacts & public cluster infrastructure (where applicable) for on-demand execution of the supported CI frameworks

- Maintain developer & user documentation along with examples associated with the chaos CI job templates 

- Provide well-defined governance and contribution guidelines for the GitHub repositories under the purview of SIG-CI

- Support queries in the community (slack/email) over questions and issues around the chaos integrations

### Non-Goals

*Note: Can be picked optionally if the SIG members are interested, as part of the scope/representative of a different SIG* 

- Implement experiment or chaos infrastructure changes necessary to support execution as CI jobs 

### Chairs/Leads

- [Udit Gaurav](https://github.com/uditgaurav)

### Repositories

- [Gitlab Remote Templates](https://github.com/mayadata-io/gitlab-remote-templates)
- [Github Chaos Actions](https://github.com/mayadata-io/github-chaos-actions)
- [Chaos CI Lib](https://github.com/mayadata-io/chaos-ci-lib)

### SIG Meeting Cadence

- Meeting Notes:
- Meeting Cadence:
- Zoom Link:
- SIG-Documentation Meet Recordings:

------------------

## SIG-Testing 

### Goals

- Create cross-reference docs around test strategy for the chaos management, orchestration and experiment business logic

- Create detailed test plans around litmuschaos functionality with a focus on both positive & failure paths. Maintain test/scenario coverage trackers with additions in every release. Interface with SIG-Chaos & SIG-Orchestration to review and drive the implementation of this plan. 

- Establish practices around the automation of identified test scenarios using BDD/e2e test frameworks such as Ginkgo-Gomega, D-Operators & maintain the test libraries. 

- Maintain the e2e Gitlab pipeline artefacts and setup/monitor automated execution of these pipelines, with reporting mechanisms to update history on the [e2e portal](https://litmuschaos.github.io/litmus-e2e)

-  Identify and extend the e2e suite for interoperability on supported chaos platforms & maintain the compatibility matrix in terms of 

    - Runtime 
    - Kubernetes Platforms (On-premise, Clouds) 
    - Operating Systems
 
- Implement & maintain hooks (interface with bots where necessary) into CI processes for on-demand execution of the e2e tests/suite via GitHub Actions with functionality into report unto PRs

- Maintain developer & user documentation associated with usage (standalone/manual execution) of the respective e2e suites

- Provide well-defined governance and contribution guidelines for the GitHub repositories under the purview of SIG-Testing

- Support queries in the community (slack/email) over questions around e2e status of release candidates, flaky-tests, and user execution of e2e suites. 

 ### Non-Goals

*Note: Can be picked optionally if the SIG members are interested, as part of the scope/representative of a different SIG*

- Write e2e tests for functionality added as part of the release (this is expected to be added by the respective developers / SIG pushing this functionality) 

### Chairs/Leads

- [Udit Gaurav](https://github.com/uditgaurav)
- [Shubham Chaudhary](https://github.com/ispeakc0de)

### Repositories

- [litmuschaos/litmus-e2e](https://github.com/litmuschaos/litmus-e2e)

### SIG Meeting Cadence

Meeting Notes:
Meeting Cadence:
Zoom Link:
SIG-Documentation Meet Recordings:


----------------------

## SIG-Documentation

### Goals

- Set the standards for feature (user guides as well as the developer) documentation as well as the process to enforce these standards

- Provide clear paths for docs contribution & maintain the GitHub repositories under the purview of SIG-Documentation

- Coordinate documentation updates during monthly project release with representatives from other SIGs 

- Curate detailed Release Information, including known issues, breaking/backward-incompatible changes 

- Review & enhance the litmus demo resources & onboarding aids 

- Support queries in the community (slack/email) over questions around docs issues and improvements 

- Maintain the documentation infrastructure, i.e., website, wiki (including UX) with respective DevOps/SRE focus: 
  - Docusaurus 
  - Deployment artefacts (docker-compose, K8s deployment) 
  - Preview tools (netlify)
  - Code embed tools (embedmd)
  - Travis CI with static checks 
  - Automated deploy of docs changes (flux) 

### Non-Goals

*Note: Can be picked optionally if the SIG members are interested, as part of the scope/representative of a different SIG*

- Add feature documentation (this is the responsibility of the other (functionality) SIGs

- Maintain design docs on the project wiki page

- Add contributor docs for repos/components other than docs

- Create Release Notes for project releases (this sig will focus on curating it from the other SIGs/Maintainers) 

### Chairs/Leads

- [Divya Mohan](https://github.com/divya-mohan0209)
- [Jayesh Kumar](https://github.com/k8s-dev)

### Repositories

- [litmuschaos/litmus](https://github.com/litmuschaos/litmus/wiki)
- [litmuschaos/litmus-docs](https://github.com/litmuschaos/litmus-docs)

### SIG Meeting Cadence

- Meeting Notes:
- Meeting Cadence:
- Zoom Link:
- SIG-Documentation Meet Recordings:

--------------------

## SIG-Orchestration

### Goals

- Contribute requirements, use-cases & code to enhance the core orchestration infrastructure in Litmus. Areas include: 

  - Simplification/lowering entry-level barriers for chaos 
  - Adding to criteria or filters for chaos
  - Static & runtime validation of user inputs 
  - State management of chaos resources
  - Improvements to the schema of chaos CRs 
  - Scheduling policies for chaos 
  - Increased Chaos targets (resources & Kubernetes platforms) support 
  - Ensuring Scalability of the orchestration components (operator, runner, scheduler, admission-controller, experiments)  

- Improving the design/code patterns and framework used for orchestration operations (reconciliation, resource management). Best practices and quality control (BCH, Go-Report, etc.,) 

- Improved logging, creation of metrics indicating the health of the chaos infrastructure

- Support for operation in air-gapped environments 

- Improvements to the CI infrastructure with increased unit test coverage, e2e/BDD-based integration tests

- Maintain developer & user documentation, deploy artifacts/examples associated with the individual orchestration pieces. 

- Provide well-defined governance and contribution guidelines for the GitHub repositories under the purview of SIG-Orchestration

- Support queries in the community (slack/email) over questions and issues around chaos orchestration 

### Non-Goals

- TBA

### Chairs/Leads

- @rahulchheda 
- @ksatchit 
- @chandankumar4 

### Repositories 

- [litmuschaos/chaos-operator](https://github.com/litmuschaos/chaos-operator)
- [litmuschaos/chaos-runner](https://github.com/litmuschaos/chaos-runner)
- [litmuschaos/chaos-scheduler](https://github.com/litmuschaos/chaos-scheduler)
- [litmuschaos/admission-controller](https://github.com/litmuschaos/admission-controller)

### SIG Meeting Cadence

- Meeting Notes:
- Meeting Cadence:
- Zoom Link:
- SIG-Documentation Meet Recordings:





