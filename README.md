# POC Test

## Quick Overview

This repo was created to show a simple implementation of microservices. There are 2 layers, 
namely the API Gateway (GraphQL) and Service Layer. Each microservice is aggregated by 
GraphQL via gRPC server to server (simulated in this test).

## Pre-Requirement

- Go 1.14+
- Docker & Docker Compose

## Introduction

for time efficiency, I don't separate each service into different repos. 
However, I have made it easy to split it into different repos.

There are 3 microservices and 1 API Gateway. In this simple POC, I will show you how to 
implement the microservice that I often do. Because there was no specific case given, 
I made a case for a seller who was just about to open a shop, and had never created an account at all. 
Here are example of a cases that I make:

- Login to get access token
- Create a new user
- Opening a new shop
- Adding products
- Edit products
- Remove products
- Get a product list
- Get product details

## How To Run Dev 

### Architecture

I'm using my own style of clean architecture, which have 3 layers:
- Server (serving)
- Use Case (business logic)
- Repo (data)

The structure will be look like this for each service:
```
.
├── app
│   └── grpc (app handler & server)
├── cmd
│   └── grpc-api (main app)
└── internal
    ├── entity 
    ├── repo (data layer)
    └── usecase (business logic)
```

### Easy Run

Run magic command:
```shell script
~$ make dev-all
```
it will spin up all services inside docker, and also it has hot reload
feature (good for dev). 

### Manual Run

Follow this step to run it manually:
1. Add this to your `/etc/hosts`
    ```
    127.0.0.1 postgres account shop product graphql
    ```
   or modify each config file to match the connection with localhost
2. From the root folder, execute this command:
    ```shell script
    ~$ make run-service
    ```
   it will spin up the postgres, redis, etc.
3. Then, run this 3 command separately:
    ```shell script
    ~$ make dev-account
    ~$ make dev-shop
    ~$ make dev-product
    ```
4. Run the graphql:
    ```shell script
    ~$ make dev-gql
    ```
5. Open the playground on http://localhost:8080 then you're ready to test the query.

## What can be improved?

- Adding elasticsearch for listing the product
- Adding redis to optimize the speed
- Improve GRPC Client implementation, so it will automatically connect when detect the target host
is down

## Unit Test

```shell script
~$ make test
```

The unit test is not covered for the following criteria:
- Any Generated file
- Main app
- Model/Entity
- pkg

Some useful info about unit test:
1. [Unit Testing and Why You Should Be Doing It](https://medium.com/better-programming/unit-testing-and-why-you-should-be-doing-it-ab61407c53ce)
2. [13 Tips for Writing Useful Unit Tests](https://medium.com/better-programming/13-tips-for-writing-useful-unit-tests-ca20706b5368)