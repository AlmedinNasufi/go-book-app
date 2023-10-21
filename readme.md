# Go Books App Api

This project is a simple RESTful API for an book app with GO

## Features

- Built with the Gin framework for efficient HTTP handling.
- Uses GORM as the Object-Relational Mapping library.
- Backed by a Mysql database.
- Comes with a Postman JSON collection for easy API testing.

## Getting Started

### Prerequisites

Ensure you have Go and Mysql installed.

### Database Migration

The project employs GORM's auto-migration feature, ensuring the MySql database schema stays synchronized with the code's model definitions. Upon application start or model changes, the schema automatically updates to match the latest models, adding necessary tables or columns. Note: To prevent data loss, it won't delete existing columns or tables.

## API Testing with Postman

Import the provided Postman collection from the repository to test the API endpoints.

## Auto Compilation

The repository is equipped with CompileDaemon for auto compilation during development. This way, you don't have to run go run main.go every time you make changes.

To initiate auto compilation, use the following command:

```bash
CompileDaemon -command="./go-crud"
```
