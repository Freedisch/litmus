LitmusChaos Releases

## 1.6.0

Refer: https://github.com/litmuschaos/litmus/milestone/20

### Objectives

- (New Feature) An alpha version of the Litmus UI/Portal to schedule chaos experiments and workflows on your Kubernetes Cluster
  * [ ] One fully functional pre-defined chaos workflow on the charthub
  * [ ] A detailed product design specification doc 
  * [ ] Figma designs for the portal on the repo
  * [ ] A basic UI that is able to execute the aforementioned chaos workflow
  * [ ] Basic authentication login of the portal
  * [ ] A well-defined CI/CD pipeline (build/test/release/deploy) mechanism 
  * [ ] e2e suite with ~50% coverage for "current" capabilities
  * [ ] Basic usage and developer documentation

- (Enhancement) Improvements to the chaos scheduler 
  * [ ] Improved OpenAPI based schema validation for the .spec.schedule 
  * [ ] Support of randomized injection of chaos in "repeat" mode

- (Enhancement) Improvements to the validating webhook admission controller
  * [ ] Validation for ConfigMap and Secret volume dependencies on the cluster
  * [ ] Validation of .spec.schedule of the ChaosSchedule (call out mutually fields, time ranges)

- (Enhancement) Improvements to the chaos experiments / chaoslib
  * [ ] Add experiment execution summary event in chaosresult 
  * [ ] Add support network-chaos on containerd & crio runtimes

- (Enhancement) Improvements to the chaos infrastructure (operator, runner) 
  * [ ] Add check gates and bulk chaos injections at namespace level
  * [ ] Add overrides for experiment image & imagePullPolicy from ChaosEngine
  
- (Enhancement) Improvements to CI infrastructure
  * [ ] Add Dockerfile linter in Travis CI
  * [ ] K8s manifest validation for charts & operator manifests
  
## 1.5.1

Refer: https://github.com/litmuschaos/litmus/milestone/21

### Objectives

- (New Feature) Improvements to chaos infrastructure
  * [ ] Add support for securitycontext definition of experiment pods
  * [ ] Add support for pod scheduling policies of experiment pods

- (New Feature) New chaos experiments 
  * [ ] Add go-based chaoslib for pod-memory-hog experiment

- (Tech Preview/PoC) Improvements to chaoslib
  * [ ] Litmus native chaoslib for network chaos experiments (nsenter-based)

- (Enhancement) Improvements to the validating webhook admission controller
  * [ ] Validate app labels'propagation into podTemplateSpec of AUT (application-under-test)

- (Enhancement) Improvements to CI infrastructure
  * [ ] Extend /run-e2e-tests capability to chaos operator, chaos-runner & litmus-go repos

- (Enhancement) Improvements to Documentation
  * [ ] Upgrades in LitmusChaos
  * [ ] Litmus uninstallation
  
  