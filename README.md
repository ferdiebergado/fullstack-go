# fullstack-go
A full stack web application using net/http, database/sql and html/template.

# Requirements
1. Go version 1.22 or higher
2. Docker or Podman

# Usage
This project uses [Task](https://taskfile.dev/) as task-runner/build tool.

Install Task by following the instructions found on this [link](https://taskfile.dev/installation/).

## Step 1
Install development tools.

The following tools are used in this project:

- [migrate](https://github.com/golang-migrate/migrate)
- [sqlc](https://sqlc.dev/)
- [air](https://github.com/air-verse/air)

Run the task below to install them:
```sh
task tools
```

## Step 2
Copy .env.example to .env.

```sh
cp .env.example .env
```

## Step 3
Set the environment variables in .env according to your setup.

If using podman, change CONTAINER variable to podman.

## Step 4
Start the database.

```sh
task db
```

## Step 5
Migrate the database.

Open another terminal and run this command:
```sh
task migrate
```

## Step 6
Run the application with auto reload.

```sh
task dev
```

## Step 7
Open your browser on [localhost:8888](http://localhost:8888).

An api is also available at the /api endpoint.

# Running Tests
First, setup the test environment:
```sh
task setup_test
```
You just need to run this once per terminal session.

Then, run the tests:
```sh
task test
```
