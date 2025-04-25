# Go Language - Todo API

This repository contains a simple Go API that provides endpoints to browse, read, edit, add, and delete todo items from a database.

## Features
- Browse all todo items
- Read details of a specific todo item
- Edit existing todo items
- Add new todo items
- Delete todo items

## Requirements
The only requirement to use this repository is the `goose` go package

Find installation guide at the [official documentation](https://pressly.github.io/goose/documentation/cli-commands/): 

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

4. Access the API at `http://localhost:8080` using API client of your choice
