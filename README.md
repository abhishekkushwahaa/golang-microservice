# Microservices with Golang

This repository contains the code for the microservices with Golang Ecommerce project. The project is divided into multiple microservices, each responsible for a specific part of the business.

## Microservices

1. **Product Service**: Responsible for managing products.
2. **User Service**: Responsible for managing users.
3. **Cart Service**: Responsible for managing the cart.
4. **Order Service**: Responsible for managing orders.
5. **Payment Service**: Responsible for managing payments.
6. **Shipping Service**: Responsible for managing shipping.

## Architecture

The project is built using the microservices architecture. Each service is a separate codebase and is deployed separately. The services communicate with each other over HTTP.

## Tech Stack

1. **Golang**: All the services are built using Golang.
2. **MySQL**: MySQL is used as the database.
3. **Docker**: Docker is used to containerize all the services
4. **Postman**: Postman is used for API testing.
5. **Gorilla Mux**: Gorilla Mux is used as the router for all the services.
6. **Gorm**: Gorm is used as the ORM for the project.
7. **RabbitMQ**: RabbitMQ is used as the message broker. It is used for communication between services.

## How to run

1. Clone the repository
2. Run `docker-compose up` to start all the services
3. Use Postman to test the APIs
4. Run `docker-compose down` to stop all the services
5. Run `docker-compose down -v` to stop all the services and remove the volumes
6. Run `docker-compose up --build` to rebuild the images and start the services
7. Run `docker-compose up -d` to start the services in the background
