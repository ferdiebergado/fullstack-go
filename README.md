# fullstack-go

A full stack web application using net/http, database/sql and html/template.

# Requirements

1. Go version 1.22 or higher
2. Docker or Podman
3. [Task](https://taskfile.dev/)
4. [Esbuild](https://esbuild.github.io/)

# Usage

## Step 1

Install development tools.

The following tools are used in this project:

-   [migrate](https://github.com/golang-migrate/migrate) for database migrations
-   [sqlc](https://sqlc.dev/) for sql code generation
-   [air](https://github.com/air-verse/air) for live reloading
-   and more

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

Run database migrations.

```sh
task migrate
```

## Step 6

Bundle the css files.

Open a new terminal and run the css:watch task.

```sh
task css:watch
```

## Step 7

Bundle the javascript files.

Open a new terminal and run the js:watch task.

```sh
task js:watch
```

## Step 7

Start the application.

```sh
task dev
```

## Step 8

Interact with the application.

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
