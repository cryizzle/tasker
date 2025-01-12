# Tasker

## Tech Stack

- **Frontend**: React, TypeScript
- **Backend**: Go
- **Database**: MySQL
- **Containerization**: Docker
- **Version Management**: nvm (Node Version Manager)

## Prerequisites

Before setting up the project locally, ensure you have the following installed:

- Docker
- [Node.js and nvm](https://github.com/nvm-sh/nvm)
- [Go](https://golang.org/doc/install)
- MySQL

## Setup Guide

### 1. Clone the Repository

```bash
git clone https://github.com/cryizzle/tasker.git
cd tasker
```

### 2. Start the Backend Server

To build and run the backend services using Docker:

```bash
make server_start
```

Note: By default, the database is started with a dump which has some data already populated. To start with empty tables, change the following volumes in `docker_compose.yml`:
```
    volumes:
      - ./db/tasker_db.sql:/docker-entrypoint-initdb.d/tasker_db.sql  # Uncomment this for fresh DB with no data
      # - ./db/tasker_backup.sql:/docker-entrypoint-initdb.d/tasker_backup.sql
```

To view live logs of the server:

```bash
make server_log
```

To restart the server after making changes to the backend code:

```bash
make server_restart
```

### 3. Run Backend Tests

Execute the following command to run backend unit tests:

```bash
make server_test
```

### 4. Connect to the Database

To access the MySQL database locally for reading and querying:

```bash
make db_connect
```

### 5. Start the Frontend Client

Ensure you have nvm installed and the appropriate Node.js version set:

```bash
nvm use
```

To start the React frontend development server:

```bash
make client_start
```

To build the frontend application:

```bash
make client_build
```
Executables are built in the `tasker_client/out` for MACOS and Windows
