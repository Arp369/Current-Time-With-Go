package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Struct for JSON response
type TimeResponse struct {
	Time string `json:"time"`
}

// Global variable for DB connection
var db *sql.DB

// Initialize MySQL database connection
func initDB() {
	var err error
	// Replace "user:password@tcp(localhost:3306)/time_api" with your actual MySQL credentials
	dsn := "root:1234@tcp(localhost:3306)/time_api"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
}

// Function to log the current time to the MySQL database
func logTimeToDB(currentTime time.Time) {
	_, err := db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
	if err != nil {
		log.Println("Error inserting time into database: ", err)
	}
}

// API handler to return the current time in Toronto
func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the current time in Toronto timezone
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, "Error loading Toronto timezone", http.StatusInternalServerError)
		return
	}
	currentTime := time.Now().In(location)

	// Log the current time to the database
	logTimeToDB(currentTime)

	// Prepare the response in JSON format
	response := TimeResponse{
		Time: currentTime.Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Initialize the database connection
	initDB()
	defer db.Close()

	// Set up the /current-time route
	http.HandleFunc("/current-time", currentTimeHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
