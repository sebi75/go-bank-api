# Go Bank API

This is a monorepo for the Go Bank API, which consists of two microservices: the Auth Service and the Banking Service. The Auth Service is responsible for handling user authentication using JWT, while the Banking Service is the main REST API for managing resources.

## Getting Started

These instructions will help you set up and run the Go Bank API on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/dl/) (v1.17 or higher)
- [Docker](https://www.docker.com/) (Optional, for running the services in containers)
- A proper database setup for each service (e.g., PostgreSQL, MySQL, etc.)

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
For example, create a .env file in the auth-service and banking-service
directories with the appropriate configuration settings, such as:

```sh
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourdbuser
DB_PASSWORD=yourdbpassword
DB_NAME=yourdbname
JWT_SECRET=yourjwtsecret
```

Make sure to replace the values with your own settings.

Build each microservice:

```sh
cd auth-service
go build
cd ../banking-service
go build
```

Run each microservice:

# In the auth-service directory

./auth-service

# In the banking-service directory

./banking-service
Alternatively, you can use Docker to run the services in containers. Create a Dockerfile for each microservice and use docker-compose to orchestrate the containers.

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

```

```
