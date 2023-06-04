# Project Documentation

## Approach to "TODOs

Throughout the development of this project, the following "TODOs" have been addressed:

**Implement bucket API endpoint compatible with S3:**

- [ ] Create a route handler to handle requests for creating buckets.
- [ ] Implement the handler to be compatible with s3.
- [ ] Implement the logic to create buckets in the  DBMS.

**Implement object API endpoint compatible with S3:**

- [ ]     Create a route handler to handle requests for uploading objects.
- [ ]     Implement the handler to be compatible with s3.
- [ ]     Implement the logic to upload objects to the corresponding bucket.

**Implement object listing API endpoint compatible with S3:**

- [ ]     Create a route handler to handle requests for listing objects in a bucket.
- [ ]     Handle the request parameters such as marker, max keys, and prefix and implement the handler to be compatible with s3.
- [ ]     Retrieve the list of objects from the DBMS and returned the response.

**Implement object download API endpoint compatible with S3:**

- [ ]     Create a route handler to handle requests for downloading objects.
- [ ]     Handle the request parameters such as range headers.
- [ ]     Retrieve the requested object from the  DBMS  and return the response.

### Project Overview & Structure

This project aims to provide an S3-compatible service using the Golang programming language and the Gin web framework. It allows users to create buckets, upload objects, list objects, and download objects.

#### Project Structure

The project follows a modular structure to promote maintainability and scalability. Here is an overview of the main directories and files:

- main.go --> Entry point of the application.
- handlers/ --> Contains the router, the routes and the handlers for different API endpoints.
- models/ --> Contains the data models.
- repositories/ --> Implements the repository layer for interacting with the database.
- services/ --> Implements the business logic layer to handle various operations.
- database/ -->  Contains the database connection.
- utils/ --> Contains the application utilities.
- .env --> Contains the application configuration.
- docs/ --> Contains the application documentaion.

#### Folder Structure

Here is an overview of the folder structure :

    .
    ├── docs
    ├── inetnal
    │   ├── datastore                   
    │       ├── pg.go
    │   ├── delivery               
    │       ├── handlers
    │       ├── router
    │   ├── domain                  
    │       ├── bucket.go
    │       ├── object.go
    │   ├── repository                   
    │       ├── bucketRepository.go
    │       ├── objectRepository.go
    │   ├── usecase                     
    │       ├── bucketService.go
    │       ├── objectService.go
    │           
    ├── utils
    │   ├── etag.go                     
    │   ├── ranger.go                  
    ├── .env
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── README.md

### Database Design

The database postgres consists of two tables: Buckets and Objects.

- The Buckets table represents different storage buckets and contains the following columns:
        name: Stores the name of the bucket.
        location: Stores the location of the bucket.

- The Objects table represents objects/files stored within the buckets and contains the following columns:
        key: Represents the key of the object.
        size: Stores the size of the object.
        content: Represents the file bytes.
        bucket_name: Acts as a foreign key referencing the name column in the Buckets table.

The design allows for organizing objects within specific buckets, with each object associated with a particular bucket using the foreign key relationship.

### Getting Started

To start playing with this project, follow these steps:

1. Clone the repository: `git clone <repository_url>` .

2. Install the dependencies: `go mod download` .

3. Configure the application settings in the .env file and create your database locally if needed.

4. Start the application: `go run main.go` or use docker

5. The application will be running on the specified port. You can now use the AWS CLI, curl or Postman to interact with the endpoints.

Example usage with curl :

- Creating a bucket: `curl --request PUT http://localhost:8080/cubbit-bucket` .

- Uploading an object: `curl --request PUT http://localhost:8080/cubbit-bucket/folder/logo.png -T file_path` .

- Listing objects: `curl http://localhost:8080/cubbit-bucket` .

- Downloading an object: `curl http://localhost:8080/cubbit-bucket/folder/logo.png --output file_path` .

Note: when using postman, make sure to upload files as binary.
