LitmusChaos Releases

## 1.6.0

Refer: https://github.com/litmuschaos/litmus/milestone/20

### Objectives

- An alpha version of the Litmus UI/Portal to schedule chaos experiments and workflows on your Kubernetes Cluster
  - A fully functional pre-defined chaos workflow on the charthub
  - A detailed product design specification doc 
  - Figma designs for the portal on the repo
  - A basic UI that is able to execute the aforementioned chaos workflow
  - Basic authentication login of the portal
  - A well-defined CI/CD pipeline (build/test/release/deploy) mechanism 
  - e2e suite with ~50% coverage
  - Basic usage and developer documentation

- Improvements to the chaos scheduler 
  - Improved OpenAPI based schema validation for the .spec.schedule 
  - Support of randomized injection of chaos in "repeat" mode