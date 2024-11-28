# Current-Time-With-Go 

#### This project is a simple API built using Go that provides the current time in Toronto, stores it in a MySQL database, and returns it as a JSON response. The application also logs the current time each time the endpoint is accessed.

## Features:
- Provides current time in Toronto in YYYY-MM-DD HH:MM:SS format.
- Logs the current timestamp into a MySQL database.
- Written in Go (Golang) and uses MySQL for data storage.

## Table of Contents
1. Prerequisites
2. Installation
3. Running the Application
4. MySQL Database Setup
5. Endpoints
6. Docker Setup

##  1. Prerequisites
To set up and run this application, you need the following:

+ Go 1.23 or higher
+ MySQL 5.7 or higher
+ Docker (optional, if you prefer containerized deployment)

## 2. Installation
#### Clone the repository
```
bash

git clone https://github.com/your-username/time-api.git
cd time-api
```
#### Install Go dependencies

Run the following command to install the required dependencies defined in the ```go.mod``` file:

```
bash

go mod download
```

## 3. Running the Application

#### Set up the MySQL Database
Before running the application, you need to set up the MySQL database. Instructions are provided below in the MySQL Database Setup section.

#### Run the Go Application
To start the server, run the following command:
```
bash

go run main.go
```
The server will start on port 8080 by default. You can visit http://localhost:8080/current-time to get the current time in Toronto.

## 4. MySQL Database Setup

#### Create the MySQL Database and Table
You need to create the MySQL database and table before starting the application.

1. Connect to MySQL:
```
bash

mysql -u root -p
```
2. Create a Database:

Run the following command to create the time_api database:
```
sql

CREATE DATABASE time_api;
```

3.Create the Table:

Switch to the time_api database and create the time_log table:

```
sql

USE time_api;
CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
);
```
4.Update Database Credentials:

In the main.go file, make sure to replace the following line with your actual MySQL credentials:
```
go

dsn := "root:1234@tcp(localhost:3306)/time_api"
```
Replace root and 1234 with your MySQL username and password.

## 5. Endpoints
```GET /current-time```

+ Description: Returns the current time in Toronto, formatted as YYYY-MM-DD HH:MM:SS.
+ Response: A JSON object with the current time.
```
Example:

json
Copy code
{
  "time": "2024-11-28 12:34:56"
}
```
## 6. Docker Setup
If you prefer to run the application using Docker, you can build and run the container with the following steps.

### Build the Docker Image
Run the following command to build the Docker image:
```
bash

docker build -t time-api .
```

### Run the Docker Container

After building the Docker image, run the container with the following command:
```
bash

docker run -d -p 8080:80 --name time-api time-api
```
The application will now be available on http://localhost:8080.

