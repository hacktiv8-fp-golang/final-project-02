# Final Project 02

## Installation
1. Clone the repo
```sh
  git clone https://github.com/hacktiv8-fp-golang/final-project-02.git
```
2. Navigate to the root directory
```sh
  cd final-project-02
```
3. Create a .env file in the root directory
```sh
  touch .env
  # or
  echo > .env
```
4. Add environment variable
```sh
  DB_HOST=Database host (example: localhost)
  DB_USER=Database username (example: postgres)
  DB_PASSWORD=Database user's password (example: postgres)
  DB_PORT=Database port (example: 5432)
  DB_NAME=Name of the database to be used (example: "database name")
  DB_SSLMODE=Database SSL mode (example: disable)
```
5. Run the project
```sh
  go run main.go
```