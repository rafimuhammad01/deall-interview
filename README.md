# DEAL Backend Engineer Technical Interview Test

## Requirements
1. Create Rest API CRUD User and User Login using NodeJS (or Golang).
2. If you're using NodeJS, we suggest you use ExpressJS. You can use databases from anywhere, but MongoDB is recommended.
3. Login with username and password to access API CRUD (token, but refresh token would be a plus).
4. Make two users with roles: 1 Admin, and 1 User.
5. The Admin has access to all API CRUD, while the User only gets access to the user's data (Read).
6. Architecture Microservices implemented using Kubernetes with Docker container deploy in a VPS (1 node with some pods inside). If you don't have a VPS, then you'll need to:
7. Preparing the YML for running the application in containerized mode and ready for Kubernetes deployment.
8. Deploy the application locally and take a screenshot for the attachment.
9. Upload source code by using GitHub Repository with the script of YML Kubernetes.
10. API documentation (Postman or Swagger) should be made available to the API rest server.
11. Make an architecture  diagram that shows the flow of API CRUD and Login.
12. Attach the Admin credential in the Readme.


## Descriptions
This repository contains 3 services:
1. API Gateway service
2. User Service 
3. Auth Service

### API Gateway
API Gateway Service is the orchestrator of all the service. It also has logic to authorize user

### User Service
User service will be responsible for 6 endpoints:
1. Get All user (admin only)
2. Get User By ID (User and Admin)
3. Create User (admin only)
4. Update user (admin only)
5. Delete User (admin only)
6. Create Admin User (No Auth Needed)

### Auth Service
Auth service will be responsible for 2 endpoints:
1. Login
2. Refresh Token

## Technologies
1. Go as programming language
2. MongoDB
3. Redis

## How to run locally with docker 
1. create .env file by copying from .env.example in 
   1. API Gateway
   2. User Service
   3. Auth Service
2. Up all the services using docker
    ```
    docker compose up -d 
    ```