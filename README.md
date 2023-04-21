# Go Bank API

This is a monorepo for the Go Bank API, which consists of two microservices: the Auth Service and the Banking Service. The Auth Service is responsible for handling user authentication using JWT, while the Banking Service is the main REST API for managing resources.

## Getting Started

These instructions will help you set up and run the Go Bank API on your local machine for development and testing purposes.

### Prerequisites

- Go (v1.17 or higher) - Optional, for local development without Docker
- Docker (for running the services in containers)
- Docker Compose (for orchestrating the containers)

### Installing

1. Clone the repository:

ssh:

```sh
git clone git@github.com:sebi75/go-bank-api.git
cd go-bank-api
```

https:

```sh
git clone https://github.com/sebi75/go-bank-api.git
cd go-bank-api
```

Set up the environment variables for each microservice.
Rename the .env.example file to .env in the root directory of each microservice, and set the values for the environment variables. For example, in the auth-service/.env file, you should have something like this:

```sh
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourdbuser
DB_PASSWORD=yourdbpassword
DB_NAME=yourdbname
JWT_SECRET=yourjwtsecret
```

Make sure to replace the values with your own.

Build and run the microservices using Docker Compose. In the root directory of the project, run the following command:

```sh
docker-compose up --build
```

Docker Compose will build the images and start the containers for both services.

Using the Monorepo in Visual Studio Code
To properly work with the monorepo in Visual Studio Code, create a go-bank-api.code-workspace file in the root directory of the repository, and include the following configuration:

```json
{
	"folders": [
		{
			"path": "auth-service"
		},
		{
			"path": "banking-service"
		}
		{
			"path": "banking-lib"
		}
	]
}
```

Now, open the go-bank-api.code-workspace file in Visual Studio Code, and you should have a properly configured workspace for the monorepo.

Usage
Once the services are up and running, you can interact with the API using tools like Postman or curl. Below are some example endpoints:

Auth Service:

POST /auth/register: Register a new user
POST /auth/login: Authenticate a user and generate a JWT token
Banking Service:

GET /customers: Get a list of customers
POST /customers: Create a new customer
PUT /customers/{id}: Update a customer
DELETE /customers/{id}: Delete a customer
Refer to the API documentation for a complete list of endpoints and request/response formats.
