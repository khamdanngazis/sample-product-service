
# Go Application

This is a Go application with migrations and seeders. The application includes a simple service and repository layer for interacting with a database, along with a migration and seeder system.

## Table of Contents
- [Go Application](#go-application)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Application](#running-the-application)
  - [Running Migrations and Seeders](#running-migrations-and-seeders)
  - [Docker Setup](#docker-setup)
  - [Contributing](#contributing)
  - [License](#license)

---

## Prerequisites

Before you begin, ensure that you have the following tools installed:

- [Go](https://golang.org/dl/) (1.20 or higher)
- [Docker](https://www.docker.com/get-started) (optional, for containerized environments)
- [Docker Compose](https://docs.docker.com/compose/install/) (optional, for managing multi-container environments)

---

## Installation

To install and set up the project, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/yourusername/yourapp.git
cd yourapp
```

2. Install Go dependencies:

```bash
go mod download
```

---
## Configuration
The application configuration is located in the **`cmd/config/config.yaml`** file. This file contains the settings required to configure various aspects of the application (such as database connections, API keys, etc.).

Make sure to update the configuration file with the necessary values before running the application.

Example **`config.yaml`**:
```yaml
database:
  main:
    host: "localhost"
    port: "5432"
    user: "postgres"
    password: "12345abc"
    dbname: "market_places"
    sslmode: "disable"
    timezone: "Asia/Jakarta"
    encoding: "UTF8"
    debug: false
redis:
  host: "localhost"
  port: "6379"
  password: ""

appport: ':8001'

```
---
## Running the Application

To run the application locally, use the following command:

```bash
cd cmd
go run main.go
```

This will start the Go application on the default port (8080).

If you are running inside a Docker container, you can follow the instructions in the Docker section.

---

## Running Migrations and Seeders

1. **Migrations**: The migration logic is implemented inside the `cmd/migration/migration.go` file. To run the migrations, use the following command:

```bash
go run cmd/migration/migration.go
```

This will apply any database migrations defined in the migration logic.

2. **Seeders**: If you have seed data to populate your database, make sure the seeder logic is implemented in a corresponding file. You can run the seeder by executing:

```bash
go run cmd/seeder/seeder.go
```

This will seed the database with initial data (if configured).

---

## Docker Setup

To run the application in a Docker container:

1. **Build and start the Docker container**:

```bash
docker-compose up --build
```

This will build and run your Go application containerized, mapping port 8001 to the host.

2. **Stopping the Docker container**:

```bash
docker-compose down
```

---

## Contributing

Feel free to fork this project, make improvements, or submit pull requests.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
