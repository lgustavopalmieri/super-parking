package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/configs"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("error to loading config: ", err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		fmt.Println("error to connecting database:", err)
		os.Exit(1)
	}
	println("connected on database", db.Ping())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8086", nil)
}
