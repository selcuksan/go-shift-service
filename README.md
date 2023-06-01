# Shift Service

This service that allows users to manage their shift periods and create their shift lists. This service allows users to select which days and times they will be on duty in a given date range and create their shifts. In addition, users can view the shift roster for their shift period, and makeshift change requests, and have the option to delete or update existing shifts. Through notifications, users can be notified of updates related to their shift periods. The Shift Service provides a user-friendly platform that helps users to schedule shifts in an organized manner and manage their roster.

## Content
- [Project structure](#project-structure)

## Project structure

### `config`

Configuration. First, `config.yml` is read, then environment variables overwrite the yaml config if they match.
The config structure is in the `config.go`.
The `env-required: true` tag obliges you to specify a value (either in yaml, or in environment variables).

Reading the config from yaml contradicts the ideology of 12 factors, but in practice, it is more convenient than
reading the entire config from ENV.
It is assumed that default values are in yaml, and security-sensitive variables are defined in ENV.

### `docs`

Swagger documentation. Auto-generated by [swag](https://github.com/swaggo/swag) library.
You don't need to correct anything by yourself.

### `handlers`

The handlers directory contains the main handlers or controllers for the project. These handlers handle incoming requests, perform necessary actions, and return appropriate responses. They encapsulate the business logic and interact with other components of the project, such as services and data repositories.

It is important to note that the project structure described here may not include all the directories and files present in the actual project. The provided overview focuses on the key directories relevant to understanding the structure and organization of the project.

## Major Packages used in this project

- **gin**: Gin is an HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need a smashing performance, get yourself some Gin.
- **mongo go driver**: The Official Golang driver for MongoDB.
- **jwt**: JSON Web Tokens are an open, industry-standard RFC 7519 method for representing claims securely between two parties. Used for Access Token and Refresh Token.
- **viper**: For loading configuration from the `.env` file. Go configuration with fangs. Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile, or Java properties formats.
- **bcrypt**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **testify**: A toolkit with common assertions and mocks that plays nicely with the standard library.
- Check more packages in `go.mod`.

### Public API Request Flow without JWT Authentication Middleware

![Public API Request Flow](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-public-api-request-flow.png?raw=true)

### Private API Request Flow with JWT Authentication Middleware

> JWT Authentication Middleware for Access Token Validation.

![Private API Request Flow](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-private-api-request-flow.png?raw=true)

### How to run this project?

We can run this Go Backend Clean Architecture project with or without Docker. Here, I am providing both ways to run this project.

- Clone this project

```bash
# Move to your workspace
cd your-workspace

# Clone this project into your workspace
git clone https://github.com/Bulut-Bilisimciler/go-shift-service.git

# Move to the project root directory
cd go-shift-service
```

#### Run without Docker

Make sure PostgreSQL is running and accessible with the credentials you provided in the .env file.

Open a terminal or command prompt and navigate to the root directory of your project.

Run the following command to execute the Go program:
- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install `go` if not installed on your machine.
- Install `PostgreSQL` if not installed on your machine.
- Important: Open the `.env` file and modify the values of `DB_HOST`, `DB_USER`, and `DB_PASSWORD` to match your PostgreSQL configuration. Update any other configuration variables if necessary.
- Run `go run main.go`.
- Access API using `http://localhost:9097`

#### Run with Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run `docker-compose up -d`.
- Access API using `http://localhost:9097`
