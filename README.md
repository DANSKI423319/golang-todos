# Go Language - Todo API

This repository aims to contain an example dockerised Go API capturing CRUD functionality.

## Features
- [X] Create new todo items
- [X] Read all todo items
- [X] Read details of a specific todo item
- [X] Update existing todo items
- [ ] Delete todo items

## Requirements

### Goose 🪿

To run the migrations you'll find it helpful to have Goose installed.

Find installation information at the [official documentation](https://pressly.github.io/goose/documentation/cli-commands/):

### Docker Desktop 🐳

You'll need docker on your machine, and the desktop app to go with it.

Find installation information at the [official website](https://www.docker.com/products/docker-desktop/)

## Getting Started
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd golang-todos
   ```

2. Setup the `.env` file based on the `.env.example`

3. Start the docker container:
   ```bash
   docker compose up -d
   ```

4. Use goose
   ```bash
   goose up
   ```

4. Access the API at `http://localhost:8080` using API client of your choice
