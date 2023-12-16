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


## Key Features CRUD Operations: Create, Read, Update, and Delete products by  id, price, and category entities.

-Create:
You can create a new product using a POST request to a specific endpoint. The request body should contain the product information, including name, price, and category. The API validates the provided data and assigns a unique ID to the new product. If successful, the API returns a response code (e.g., 201 Created) and the newly created product data.

-Read:
You can retrieve product information using GET requests to different endpoints. The response body will contain the requested product information (name, price, category, and ID).

-Update:
You can update existing product information using a PUT request to the specific product endpoint (e.g., /products/{id}). The request body should contain the updated values for any of the product attributes (id, name, price, category). The API validates the updated data and updates the corresponding product record in the database.

-Delete:
You can remove a product from the system using a DELETE request to the specific product endpoint. The API verifies the product ID and deletes the corresponding record from the database.

### Scalability and Performance: Utilize Docker for easy deployment and containerization for efficient resource management. Database Management: Utilize a MySQL Docker image for reliable and persistent data storage.
  # Image Used - MySQL
