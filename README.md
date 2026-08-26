<a href="https://docker.com">
    <img src="https://raw.githubusercontent.com/kreuzwerker/terraform-provider-docker/master/assets/docker-logo.png" alt="Docker logo" title="Docker" align="right" height="100" />
</a>
<a href="https://terraform.io">
    <img src="https://raw.githubusercontent.com/kreuzwerker/terraform-provider-docker/master/assets/terraform-logo.png" alt="Terraform logo" title="Terraform" align="right" height="100" />
</a>

# Terraform Provider for Docker

[![Release](https://img.shields.io/github/v/release/gocommunity-io/terraform-provider-dockerd)](https://github.com/gocommunity-io/terraform-provider-dockerd/releases)
[![Installs](https://img.shields.io/badge/dynamic/json?logo=terraform&label=installs&query=$.data.attributes.downloads&url=https%3A%2F%2Fregistry.terraform.io%2Fv2%2Fproviders%2F713)](https://registry.terraform.io/providers/gocommunity-io/dockerd)
[![Registry](https://img.shields.io/badge/registry-doc%40latest-lightgrey?logo=terraform)](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/gocommunity-io/terraform-provider-dockerd/blob/main/LICENSE)  
[![Acc Tests](https://github.com/gocommunity-io/terraform-provider-dockerd/actions/workflows/acc-test.yaml/badge.svg?branch=main)](https://github.com/gocommunity-io/terraform-provider-dockerd/actions/workflows/acc-test.yaml)
[![golangci-lint](https://github.com/gocommunity-io/terraform-provider-dockerd/actions/workflows/golangci-lint.yaml/badge.svg?branch=main)](https://github.com/gocommunity-io/terraform-provider-dockerd/actions/workflows/golangci-lint.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gocommunity-io/terraform-provider-dockerd)](https://goreportcard.com/report/github.com/gocommunity-io/terraform-provider-dockerd)

Sponsored by [Coder](https://coder.com/)

## What You Can Manage With This Provider

This provider covers more than basic Docker images and containers. With Terraform, you can manage:

* Compose applications with [`docker_compose`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/compose)
* Image builds and registry workflows with [`docker_image`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/image), [`docker_registry_image`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/registry_image), and [`docker_tag`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/tag)
* Buildx builders for advanced multi-platform builds with [`docker_buildx_builder`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/buildx_builder)
* Swarm services with [`docker_service`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/service)
* Runtime resources such as [`docker_container`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/container), [`docker_network`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/network), and [`docker_volume`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/volume)
* Supporting platform objects like [`docker_config`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/config), [`docker_secret`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/secret), and [`docker_plugin`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/resources/plugin)
* Operational actions such as [`docker_exec`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/actions/exec), [`docker_image_import`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/actions/image_import), [`docker_image_load`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/actions/image_load), [`docker_image_save`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/actions/image_save), [`docker_container_export`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/actions/container_export) and [`docker_system_prune`](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs/actions/system_prune), 

Available data sources include images, image tags and manifests, containers, networks, plugins, and container logs. See the full [provider documentation](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs) for the complete resource and data source list.

## Documentation

The documentation for the provider is available on the [Terraform Registry](https://registry.terraform.io/providers/gocommunity-io/dockerd/latest/docs).
You need at least Terraform `1.1.5` to use this provider.

Migration guides:
* Do you want to migrate from `v3.x` to `v4.x`? Please read the [V3 - V4 migration guide](docs/v3_v4_migration.md)
* Do you want to migrate from `v2.x` to `v3.x`? Please read the [V2 - V3 migration guide](docs/v2_v3_migration.md)

## Example usage

Take a look at the examples in the [documentation](https://registry.terraform.io/providers/gocommunity-io/dockerd/4.5.0/docs) of the registry
or use the following example:


```hcl
# Set the required provider and versions
terraform {
  required_providers {
    # We recommend pinning to the specific version of the Docker Provider you're using
    # since new versions are released frequently
    docker = {
      source  = "gocommunity-io/dockerd"
      # or if you want to pull from opentfu
      source = "registry.opentofu.org/gocommunity-io/dockerd"
      version = "4.5.0"
    }
  }
}

# Configure the docker provider
provider "docker" {
}

# Create a docker image resource
# -> docker pull nginx:latest
resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = true
}

# Create a docker container resource
# -> same as 'docker run --name nginx -p8080:80 -d nginx:latest'
resource "docker_container" "nginx" {
  name    = "nginx"
  image   = docker_image.nginx.image_id

  ports {
    external = 8080
    internal = 80
  }
}

# Or create a service resource
# -> same as 'docker service create -d -p 8081:80 --name nginx-service --replicas 2 nginx:latest'
resource "docker_service" "nginx_service" {
  name = "nginx-service"
  task_spec {
    container_spec {
      image = docker_image.nginx.repo_digest
    }
  }

  mode {
    replicated {
      replicas = 2
    }
  }

  endpoint_spec {
    ports {
      published_port = 8081
      target_port    = 80
    }
  }
}
```

## Building The Provider

[Go](https://golang.org/doc/install) 1.18.x (to build the provider plugin)


```sh
$ git clone git@github.com:gocommunity-io/terraform-provider-dockerd
$ make build
```

## Contributing

The Terraform Docker Provider is the work of many of contributors. We appreciate your help!

To contribute, please read the contribution guidelines: [Contributing to Terraform - Docker Provider](CONTRIBUTING.md)

## License

The Terraform Provider Docker is available to everyone under the terms of the Mozilla Public License Version 2.0. [Take a look the LICENSE file](LICENSE).

## Sponsors

[![Coder](https://avatars.githubusercontent.com/u/95932066?s=100&v=2)](https://coder.com/)