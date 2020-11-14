# POC Test

## Quick Overview

This repo was created to show a simple implementation of microservices. There are 2 layers, 
namely the API Gateway (GraphQL) and each microservice. Each microservice is aggregated by 
GraphQL via gRPC.

## Pre-Requirement

- Go 1.13+
- Docker & Docker Compose

## Introduction

for time efficiency, I don't separate each service into different repos. 
However, I have made it easy to split it into different repos.

There are 3 microservices and 1 API Gateway. In this simple POC, I will show you how to 
implement the microservice that I often do. Because there was no specific case given, 
I made a case for a seller who was just about to open a shop, and had never created an account at all. 
Here are example of a cases that I make:

- Create a new user to register as a seller
- Opening a new shop
- Adding products
- Edit products
- Remove products
- Get a product list
- Get product details

