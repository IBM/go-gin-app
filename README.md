# WARNING: This repository is no longer maintained :warning:

> This repository will not be updated. The repository will be kept available in read-only mode.
<p align="center">
    <a href="https://cloud.ibm.com">
        <img src="https://cloud.ibm.com/media/docs/developer-appservice/resources/ibm-cloud.svg" height="100" alt="IBM Cloud">
    </a>
</p>


<p align="center">
    <a href="https://cloud.ibm.com">
    <img src="https://img.shields.io/badge/IBM%20Cloud-powered-blue.svg" alt="IBM Cloud">
    </a>
    <img src="https://img.shields.io/badge/platform-go-lightgrey.svg?style=flat" alt="platform">
    <img src="https://img.shields.io/badge/license-Apache2-blue.svg?style=flat" alt="Apache 2">
</p>


# Create and deploy a Go Gin application

> We have applications available for [Node.js Express](https://github.com/IBM/node-express-app), [Go Gin](https://github.com/IBM/go-gin-app), [Python Flask](https://github.com/IBM/python-flask-app), [Python Django](https://github.com/IBM/python-django-app), [Java Spring](https://github.com/IBM/java-spring-app), [Java Liberty](https://github.com/IBM/java-liberty-app), [Swift Kitura](https://github.com/IBM/swift-kitura-app), [Android](https://github.com/IBM/android-app), and [iOS](https://github.com/IBM/ios-app).

In this sample web application, you will create a Go cloud application using Gin. This application contains an opinionated set of files for web serving:

- `public/index.html`
- `public/404.html`
- `public/500.html`

This application also enables a starting place for a Go microservice using Gin. A microservice is an individual component of an application that follows the **microservice architecture** - an architectural style that structures an application as a collection of loosely coupled services, which implement business capabilities. The microservice exposes a RESTful API matching a [Swagger](http://swagger.io) definition.

## Steps

### Deploying to IBM Cloud

[IBM Cloud DevOps](https://www.ibm.com/cloud/devops) services provides toolchains as a set of tool integrations that support development, deployment, and operations tasks inside IBM Cloud.

Also, this application comes with the following capabilities:
- [Swagger UI](http://swagger.io/swagger-ui/) running on: `/explorer`
- An OpenAPI 2.0 definition hosted on: `/swagger/api`
- A Healthcheck: `/health`

### Building locally

To get started building this web application locally, you can either run the application natively or use the [IBM Cloud Developer Tools](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started) for containerization and easy deployment to IBM Cloud.

All of your `go` dependencies are listed in `go.mod`.

#### Native application development

- Install [Go](https://golang.org/dl/)

In order for Go applications to run anywhere including your $GOPATH/src
```
export GO111MODULE=on
```

Fetch and install dependencies listed in go.mod:
```bash
go build ./...
```

To run your application locally:
```bash
go run server.go
```

Your sources will be compiled to your `$GOPATH/bin` directory. Your application will be running at `http://localhost:8080`. You can also verify the state of your locally running application using the Selenium UI test script included in the `scripts` directory.

#### IBM Cloud Developer Tools

Install [IBM Cloud Developer Tools](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started) on your machine by running the following command:
```
curl -sL https://ibm.biz/idt-installer | bash
```

Create an application on IBM Cloud by running:

```bash
ibmcloud dev create
```

This will create and download a starter application with the necessary files needed for local development and deployment.

Your application will be compiled with Docker containers. To compile and run your app, run:

```bash
ibmcloud dev build
ibmcloud dev run
```

This will launch your application locally. When you are ready to deploy to IBM Cloud on Kubernetes, run one of the commands:

```bash
ibmcloud dev deploy -t container // to K8s cluster
```

You can build and debug your app locally with:

```bash
ibmcloud dev build --debug
ibmcloud dev debug
```

## CRA Scanning 

This repository includes a `.cra/.cveignore` file that is used by Code Risk Analyzer (CRA) in IBM Cloud Continuous Delivery. This file helps address vulnerabilities that are found by CRA until a remediation is available, at which point the vulnerabilities will be addressed in the respective package versions. CRA keeps the code in this repository free of known vulnerabilities, and therefore helps make applications that are built on this code more secure. If you are not using CRA, you can safely ignore this file.

## License

This sample application is licensed under the Apache License, Version 2. Separate third-party code objects invoked within this code pattern are licensed by their respective providers pursuant to their own separate licenses. Contributions are subject to the [Developer Certificate of Origin, Version 1.1](https://developercertificate.org/) and the [Apache License, Version 2](https://www.apache.org/licenses/LICENSE-2.0.txt).

[Apache License FAQ](https://www.apache.org/foundation/license-faq.html#WhatDoesItMEAN)

