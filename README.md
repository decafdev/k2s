# K2S

[![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)]([https://https://discord.com/](https://discord.com/channels/929003936709509160/1038103432378187776))

## Objectives

1. Eliminate the need for kubernetes specific knowledge when deploying api & event driven services
2. Provide a simple mechanism for cluster admission control improving the overall security posture

## Key Results

- A first time user can create their first k2s deployment in under 5 minutes without any prior kubernetes experience.

## Features

- [ ] support the use of a private registry
- [ ] support api deployments

## Configuration Options

- `PRIVATE_REGISTRY_URL: https://my-registry.domain.io`
- `PRIVATE_REGISTRY_USER: service-user-name`
- `PRIVATE_REGISTRY_PASS: service-user-password`
- `TRAEFIK_VERSION: 0.2.5`
- `TRAEFIK_REPLICAS: 3`

## Endpoints

create a new immutable deployment of the specified type

> POST:/deployments

list all deployments

> GET:/deployments

List all deployments for a specific service

> GET:/deployments/:name

Get deployment details for a specific version

> GET:/deployments/:name/:version

