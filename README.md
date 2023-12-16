## Unique Installation Steps  - 
- Begin by cloning this repository from GitHub 
- Ensure that the Go programming language is installed on your system (https://go.dev/doc/install)
- Verify that Docker is installed on your machine (https://www.docker.com/products/docker-desktop/)
- Navigate to the cloned directory and execute the following command
   > go mod download

## Starting the Project - 
- Initiate Docker Desktop or Docker
- Next, execute the following command to run a Docker container
  > docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
- Run the command `go run main.go` to start the project
- You can test all the APIs using Postman [by downloading and testing the API](/Product%20API%20collection.postman_collection.json) 

# Running Unit Tests - 
- Ensure that the previous program running with `go run main.go` is closed
- Execute `go test ./...` to run unit tests, or alternatively, run `go test ./... -v`


 
