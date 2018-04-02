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

`go get github.com/samirprakash/boom`

# Usage

### Pre-requisites

`boom` assumes a few things to be in installed in your local build environment or on the remote build box, in order to execute the commands. Check that the below dependencies are satisfied.

* Install [maven](https://maven.apache.org/)
* Setup [maven](https://maven.apache.org/) with a valid `settings.xml` in `./m2` folder
* Install [docker](https://www.docker.com/)

### Available sub-commands
* boom maven 
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
* boom docker
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

### Available maven sub-commands and flags
* boom maven build
  - validates and compiles the code base
  - cleans up the workspace
  - runs unit tests with JaCoCo code coverage plugin
  - packages the compiled code in the output as defined in `pom.xml`
  - runs and uploads sonar static analysis reports

* boom maven validate
  - validates and compiles the code base
  
* boom maven clean
  - cleans up the workspace

* boom maven test [ --integration-tests | -i ] [ --unit-tests | -u ] -h
  - provides option to selectively execute integration tests or unit tests
  - with `--integration-tests` or `-i` flag executes the tests marked with a category `integration-tests`
  - with `--unit-tests` or `-u` flag executes the tests marked with a category `unit-tests`
  - without any flags executes all the tests which do not have a category defined

* boom maven package [ --skip-test | -s ] -h
  - generates a JAR artifact from the code base
  - with `--skip-test` or `-s` flag provides an option to skip the test execution while packaging

* boom maven verify -h
  - verifies the integration test results
  - should run after the code has been packaged and integration tests have been executed

* boom maven deploy [ --repository-id ] -h
  - deploys artifacts to remote repository such as artifactory or nexus
  - with `--repository-id` set to the `repository-id-configured-in-pom`, it would upload the artifact
  - authentication for remote repository needs to be configured beforehand by the user

#### deploy

### Docker options

#### build
#### compose
#### test
#### tag


