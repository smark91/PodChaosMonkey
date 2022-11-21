# PodChaosMonkey

### Introduction

This golang program delete a random pod from a specified namespace every specified timespan.
This is achieved using the Kubernetes API internally through a ServiceAccount 

### ENV

- DELETE_DELAY: timeout between one iteration and the others in seconds
- DELETE_NAMESPACE: name of namespace where the deletion happen

### How to run test
`docker buildx build --platform linux/amd64,linux/arm64/v8 -t <org_name>/pod-chaos-monkey:<pod_version> --target test .`

### How to build multiple arch (arm64/amd64)
`docker buildx build --platform linux/amd64,linux/arm64/v8 --push -t <org_name>/pod-chaos-monkey:<pod_version> .`