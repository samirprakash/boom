![boom logo from http://pngtree.com/](./assets/img/b444f6901c17b1d5cf0791c4356e5f90.png)

# Table of Contents

- [Overview](#overview)
- [Usage](#Usage)
  * [maven](#maven)
  * [docker](#docker)

# Overview

`boom` is a simple but opinionated command line interface providing options to emulate CI steps - either on your local development environment or on a cloud based CI provider. Currently it handles [maven](https://maven.apache.org/) based projects and supporting [docker](https://www.docker.com/) commands to build, test, package a code base, generating [docker](https://www.docker.com/) images which are ready to be deployed to any cloud provider - [GKE](https://cloud.google.com/kubernetes-engine/), [Azure](https://portal.azure.com/), [AWS](https://aws.amazon.com/) or private cloud hosted in your organization - using an orchestration provider like [kubernetes](https://kubernetes.io/) of your choice.

`boom` is built using the [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper) package and [golang](https://golang.org/)

# Installation

* Run `go get github.com/samirprakash/boom` to install the latest version of this application

`or`

* Clone `github.com/samirprakash/boom` to your `$GOPATH`
* CD to the cloned repo
* Run `go install ./cmd/boom`

# Usage

### Pre-requisites

`boom` assumes a few things to be in installed in your local build environment or on the remote build box, in order to execute the commands. Check that the below dependencies are satisfied.

* Install [maven](https://maven.apache.org/)
* Setup [maven](https://maven.apache.org/) with a valid `settings.xml` in `./m2` folder
* Install [docker](https://www.docker.com/)
* Do a `docker login` with your credentials

### Available sub-commands
* __boom maven__
  - execute maven based commands
  - use `boom maven -h` to check options and flags
  ```  
  Usage:
    boom maven [command]

  Available Commands:
    build       All in One - Validate, Compile, Clean, Unit test, Package, Code Coverage and Sonar
    clean       Cleans up your workspace
    deploy      Copies generated package to artifactory or nexus
    package     Packages your compiled code in a distributable format
    test        Executes unit tests facilitated with code coverage
    validate    Performs a validation and checks for compilation issues
    verify      Runs quality checks on integration test results

  Flags:
    -h, --help   help for maven
  ```

  * __boom maven build__
    - validates and compiles the code base
    - cleans up the workspace
    - runs unit tests with JaCoCo code coverage plugin
    - packages the compiled code in the output as defined in `pom.xml`
    - runs and uploads sonar static analysis reports

  * __boom maven validate__
    - validates and compiles the code base
    
  * __boom maven clean__
    - cleans up the workspace

  * __boom maven test__
    - provides option to selectively execute integration tests or unit tests
    - with `--integration-tests` or `-i` flag executes the tests marked with a category `integration-tests`
    - with `--unit-tests` or `-u` flag executes the tests marked with a category `unit-tests`
    - without any flags executes all the tests which do not have a category defined

  * __boom maven package__
    - generates a JAR artifact from the code base
    - with `--skip-test` or `-s` flag provides an option to skip the test execution while packaging

  * __boom maven verify__
    - verifies the integration test results
    - should run after the code has been packaged and integration tests have been executed

  * __boom maven deploy__
    - deploys artifacts to remote repository such as artifactory or nexus
    - with flag `--repository-id` set to the string value of  `repository-id-configured-in-pom` it uploads the artifact
    - authentication for remote repository needs to be configured beforehand by the user

* __boom docker__
  - execute docker commands
  - use `boom docker -h` to check options and flags
  ```
  Usage:
  boom docker [command]

  Available Commands:
    build       Build docker images and push to a remote repository
    compose     Create docker compose environment based on the docker-compose.yaml in the code base
    tag         tag and push images to docker registry
    test        run collection of tests using newman command line runner

  Flags:
    -h, --help   help for docker
  ```

  * __boom docker build__
    - builds a docker image based on a `Dockerfile`
    - as a default, needs to be executed from the folder which contains your `Dockerfile`
    - pushes the built image to a remote docker registry, only if the current branch is `master` or `develop`
    - with `--image-tag` or `-i` flag set to the value of the tag for the image will tag the image accordingly
    - with `--app-type` or `-t` flag set to the type of service i.e. -  client/service/demo - would append this value to the tag

  * __boom docker compose__
    - configures any pre-requisites required for spinning up the docker containers
    - `--repo-name` or `-r` can be used to clone the `config-source` repo locally by providing the repository name as an entry to this flag
    - spins up a docker compose environment based on the provided docker compose file as a entry for the flag `--compose-file` or `-f`
    - checks the health status of the docker containers based on the ports that have been exposed which can be provided to the command as a comma seperated string i.e. `9090,9091,9092......9099` to the flag `--helthcheck-ports` or `-p`

  * __boom docker test__
    - executes newman tests on the docker compose environment
    - with flag `--network-bridge` or `-n` set to the default network for this compose environment. This value can be found out using `docker network ls`
    - with the flag `--test-collection` or `-c` set to the test collection
    - with the flag `--environment-spec` or `-e` set to the container environment

  * __boom docker tag__
    - tags and pushes the images generated by the docker compose environment to a remote docker registry
    - with flag `--current-image` or `-i` set to the current image which needs to be tagged as per remote docker registry requirements
    - with flag `--new-image` or `-n` set to the new tag that we want to the tag the image with

