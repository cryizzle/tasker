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

## Prioritised Featureset

### 1. Memberships

Introducing user roles such as OWNER and MEMBER allows for more flexible permission management within a Todo List. This feature sets the foundation for implementing finer-grain access controls in the future, such as adding VIEWER roles who can only view tasks without editing capabilities.

### 2. Todo History

Given the collaborative nature of Tasker, tracking changes made to Todo Lists ensures accountability and clarity. By logging who made changes and when, teams can better manage task responsibilities and resolve conflicts.

### 3. SSE Updates for Todo List

Server-Sent Events (SSE) provide real-time updates to users, enabling seamless collaboration without the need for frequent polling or manual page refreshes. This improves user experience by maintaining up-to-date task views effortlessly.

## Deprioritised Featureset

### 1. Pagination

While pagination can improve performance when handling a large number of tasks, the initial development focuses on collaboration and task management features. Pagination will be prioritized later as the application scales.

### 2. User Authentication

A simple layer to identify which user is associated with each action has been implemented. However, more comprehensive user authentication mechanisms have been deprioritized to focus on the collaborative aspects of the application first. Expanding authentication features will be addressed in future iterations.
