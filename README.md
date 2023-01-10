# K2S

[![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.com/channels/929003936709509160/1038103432378187776)

K2s enables staggeringly simple and opinionated Kubernetes deployments.

## Objectives

1. Eliminate the need for Kubernetes-specific specific knowledge when deploying API & event-driven services
2. Provide a simple mechanism for cluster admission control, improving the overall security posture

## Key Results

- A first-time user can create their first k2s deployment in under 5 minutes without any prior Kubernetes experience.

## Features

- [ ] support the use of a private registry
- [ ] support API deployments

## Setup and teardown

### Get a k8s cluster up and running

1. Install KIND `brew install kind`
2. Create cluster `kind create cluster --config ./kind-config.yml -n local`
3. Start k2s `go run . start` (`go run . -h` for help)

After you deploy your service, you access it via local port 32080. Traefik is the ingress. You access Traefik's dashboard via local port 32088.

### Teardown

If you want to delete the k2s-operator resources (better commands to come):

```bash
kubectl delete all --all -n k2s-operator
kubectl delete namespace k2s-operator
```

If you want to tear your cluster down, run `kind delete cluster -n local`.

## Configuration Options

- `PRIVATE_REGISTRY_URL: https://my-registry.domain.io`
- `PRIVATE_REGISTRY_USER: service-user-name`
- `PRIVATE_REGISTRY_PASS: service-user-password`
- `TRAEFIK_VERSION: 0.2.5`
- `TRAEFIK_REPLICAS: 3`

## Endpoints

Create a new immutable deployment of the specified type

> POST:/deployments

List all deployments

> GET:/deployments

List all deployments for a specific service

> GET:/deployments/:name

Get deployment details for a specific version

> GET:/deployments/:name/:version
