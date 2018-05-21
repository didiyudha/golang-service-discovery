# golang-service-discovery

This is a simple example to integrate golang services to Consul. This repo cover how golang services register its self to Consul, discover service when a service need to call another service, provide a healthcheck api to be consumed by Consul, and get Key Value (KV) store that we populate in Consul UI and consumed by golang services.

# How to run
- Makesure docker is already installed on your machine
- Change directory to deploy
- *docker-compose-build*
- *docker-compose up*
