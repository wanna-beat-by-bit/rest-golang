# rest-golang
Training project to write REST API server on Golang
Main purpose of the project to represent abilities in usage the following things:
- REST architecture
- Golang knowledge
- Folder arhitecture
- Writing code that is close to real production

Idea of a project - create a backend for task manager application.

# Technology stack
From the side of a language was used following packages:
- GIN framework (For endpoint and hanlders manipulation)
- viper (Config management)
- godotenv (To protect important data such as passwords from db)
- logrus (to log errors)
- migrate (to migrate database schema)

When deal with database, was used following technologies:
- Docker (Was created container with database)
- Postgres (Database to store data)
- Postman (To test server api)

# How to run the project ?
To run a project you need to configure docker at port 8000
with postgres database in it.
Example for a command in order to bring up a docker would be:

docker run --name=task-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

And to migrate data in database, you can must run the command below, while locating in root folder:
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

If you want to clear the database, run this:
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

So, last step is to run the whole application, execute the command below, while locating in root folder:
go run cmd/main.go 
