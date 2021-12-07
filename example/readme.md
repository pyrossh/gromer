# Example

This example demonstrates gromer with a simple web app.

It uses postgres as the database and sqlc to generate queries and models. It uses dbmate for migrations. 

# Requirements

1. go >= 1.16
2. docker >= 20

# Running

```sh
make setup
make start-db
make migrate
make run
```