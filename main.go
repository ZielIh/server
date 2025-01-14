package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"server/db"
	"server/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {

	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("DB init error: %v", err)
	}

	http.HandleFunc("/detect", func(w http.ResponseWriter, r *http.Request) {
		handlers.DetectCheat(w, r, database)
	})
	log.Println("Starting server on :80")
	log.Fatal(http.ListenAndServe(":80", nil))

}
