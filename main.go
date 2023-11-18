package main

import (
	"database/sql"
	"fmt"
	"github.com/banaaron/tekken-backend/database"
	"github.com/banaaron/tekken-backend/handlers"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func init() {
	// initialize the database connection
	database.InitDb()
	fmt.Println("http://localhost:8888/api/character")
}

func main() {
	var err error
	githubURL := "https://github.com/aarontbarratt/tekken-backend#tekken-backend"

	server := http.NewServeMux()
	server.Handle("/api/help", http.RedirectHandler(githubURL, http.StatusSeeOther))
	server.HandleFunc("/api/teapot", handlers.HandleTeapot)
	server.HandleFunc("/api/character", handlers.HandleCharacter)
	err = http.ListenAndServe(":8888", server)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Db *sql.DB) {
		err := Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.Db)
}
