![boom logo from http://pngtree.com/](./img/b444f6901c17b1d5cf0791c4356e5f90.png)

# Table of Contents

- [Overview](#overview)
- [Usage](#Usage)
  * [maven](#maven)
  * [docker](#docker)

# Overview

`boom` is a simple but opinionated command line interface providing options to emulate CI steps - either on your local development environment or on a cloud based CI provider. Currently it handles [maven](https://maven.apache.org/) based projects and supporting [docker](https://www.docker.com/) commands to build, test, package a code base, generating [docker](https://www.docker.com/) images which are ready to be deployed to any cloud provider - [GKE](https://cloud.google.com/kubernetes-engine/), [Azure](https://portal.azure.com/), [AWS](https://aws.amazon.com/) or private cloud hosted in your organization - using an orchestration provider like [kubernetes](https://kubernetes.io/) of your choice.

`boom` is built using the [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper) package and [golang](https://golang.org/)

# Installation

Run `go get` to install the latest version of this application

`go get github.com/samirprakash/go-boom`

# Usage

### Pre-requisites

`boom` assumes a few things to be in installed in your local build environment or on the remote build box, in order to execute the commands. Check that the below dependencies are satisfied.

* Install [maven](https://maven.apache.org/)
* Setup [maven](https://maven.apache.org/) with a valid `settings.xml` in `./m2` folder
* Install [docker](https://www.docker.com/)

### Maven options

#### build
#### validate
#### clean
#### test
#### package
#### verify
#### deploy

### Docker options

#### build
#### compose
#### test
#### tag


