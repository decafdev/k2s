# K2S

[![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.com/channels/929003936709509160/1038103432378187776)

## Objectives

1. Eliminate the need for Kubernetes-specific knowledge when deploying API & event-driven services
2. Provide a simple mechanism for cluster admission control, improving the overall security posture

## Key Results

- A first-time user can create their first k2s deployment in under 5 minutes without any prior Kubernetes experience.

## Features

- [ ] support the use of a private registry
- [ ] support API deployments

## Setup and Teardown!

### Get a k8s cluster up and running

1. Install kind `brew install kind`
2. Create a cluster `kind create cluster --config ./kind-config.yml --name local`
3. Start k2s `go run .`

After you deploy your service, you access it via local port 32080. You access the Traefik ingress's dashboard via local port 32088.

### Teardown

If you want to tear your cluster down, `kind delete` doesn't seem to work very reliably, and we don't currently implement anything nicer. So:

First, change your Kubernetes context to something other than `kind-local`. Next, delete the `kind-local` context with `kubectl config delete-context kind-local`. Finally, run:

```
docker stop local-control-plane && docker rm local-control-plane
docker stop local-worker && docker rm local-worker
docker stop local-worker2 && docker rm local-worker2
```

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
