package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Simple struct for holding user information
type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

var db *sql.DB

func contactHandler(w http.ResponseWriter, r *http.Request) {
	// Providing contact info with a structured JSON response
	contactInfo := map[string]string{
		"email":   "example@email.com",
		"phone":   "+1234567890",
		"website": "https://example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactInfo)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name, age, location FROM users")
	if err != nil {
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age, &p.Location)
		if err != nil {
			http.Error(w, "Failed to scan data", http.StatusInternalServerError)
			return
		}
		users = append(users, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// =================== //

func getGreetings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, message FROM greetings")
	if err != nil {
		http.Error(w, "Failed to fetch greetings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Greeting struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	}

	var greetings []Greeting
	for rows.Next() {
		var g Greeting
		err := rows.Scan(&g.ID, &g.Message)
		if err != nil {
			http.Error(w, "Error scanning greeting", http.StatusInternalServerError)
			return
		}
		greetings = append(greetings, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(greetings)
}

func createGreeting(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Message string `json:"message"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Message == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO greetings (message) VALUES ($1)", input.Message)
	if err != nil {
		http.Error(w, "Failed to insert greeting", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Greeting created")
}

func updateGreeting(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.ID == 0 || input.Message == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("UPDATE greetings SET message=$1 WHERE id=$2", input.Message, input.ID)
	if err != nil {
		http.Error(w, "Failed to update greeting", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Greeting not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Greeting updated")
}

func deleteGreeting(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.ID == 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM greetings WHERE id=$1", input.ID)
	if err != nil {
		http.Error(w, "Failed to delete greeting", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Greeting not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Greeting deleted")
}

// ======================= //

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getGreetings(w, r)
	case "POST":
		createGreeting(w, r)
	case "PUT":
		updateGreeting(w, r)
	case "DELETE":
		deleteGreeting(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ======================= //

type Menus struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

func getAllMenus(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT id, item FROM menus")

	if err != nil {
		http.Error(w, "Failed to fetch menus", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var menus []Menus
	for rows.Next() {
		var m Menus
		err := rows.Scan(&m.ID, &m.Item)
		if err != nil {
			http.Error(w, "Error scaning menus",
				http.StatusInternalServerError)
			return
		}
		menus = append(menus, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}

func creatMenu(w http.ResponseWriter, r *http.Request) {
	var menu struct {
		Item string `json:"item"`
	}

	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil || menu.Item == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO menus (item) VALUES ($1)",
		menu.Item)

	if err != nil {
		http.Error(w, "Failed to insert menus item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Menu item is created")
}

func menuHandlar(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllMenus(w, r)
	case "POST":
		creatMenu(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	var err error

	connStr := "user=postgres password= dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("Connected to PostgreSQL on port: http://localhost:3000")

	router := http.NewServeMux()

	// Register routes
	router.HandleFunc("/contact", contactHandler)
	router.HandleFunc("/hello", helloHandler)
	router.HandleFunc("/users", getUsersHandler)
	router.HandleFunc("/menus", menuHandlar)

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println("Server error:", err)
	}
}
