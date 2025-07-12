package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func welcomeHandlar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to your recap era....")
}

// ================== \\

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
			http.Error(w, "Error is scaning", http.StatusInternalServerError)
			return
		}

		menus = append(menus, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}

func menuHandlar(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllMenus(w, r)
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

	fmt.Println("Connected to PostgreSQL on port: http://localhost:5000")

	mux := http.NewServeMux()

	mux.HandleFunc("/", welcomeHandlar)
	mux.HandleFunc("/menus", menuHandlar)

	err = http.ListenAndServe(":5000", mux)

	if err != nil {
		log.Println("Server Error:", err)
	}

}
