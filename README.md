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
