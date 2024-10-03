# fullstack-go
A full stack web application using net/http, database/sql and html/template.

# Usage
This project uses [Task](https://taskfile.dev/) to run tasks.

Instructions on how to install Task is found on this [link](https://taskfile.dev/installation/).

## Step 1
Copy .env.example to .env.

```sh
cp .env.example .env
```

## Step 2
Set the environment variables in .env according to your setup.

## Step 3
Start the database.

```sh
task db
```

## Step 4
Migrate the database.

Open another terminal and run this command:
```sh
task migrate
```

## Step 5
Run the application with auto reload.

```sh
task dev
```

## Step 6
Open your browser on [localhost:8888](http://localhost:8888).

# Install Tools
The following tools are used in this project:

- [sqlc](https://sqlc.dev/)
- [air](https://github.com/air-verse/air)

To install them:
```sh
task tools
```

# Running Tests
First, setup the test environment:
```sh
task setup_test
```

Then, run the tests:
```sh
task test
```
