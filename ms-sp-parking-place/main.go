package main

import (
	"database/sql"
	"fmt"

	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/configs"
	//"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/adapters/web/handlers/address"
	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/adapters/web/routes"
	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/adapters/web/server"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	println("connected on database", db.Ping())

	webserver := server.NewWebServer(":8086")
	routes.SetupAddressRoutes(webserver, db)
	//webAddresHandler := address.NewWebAddressHandler(db)
	//webserver.AddHandler("/address/create", webAddresHandler.Create)
	webserver.Start()

	// http.HandleFunc("/address/create", webAddresHandler.Create)

	// http.ListenAndServe(":8086", nil)
}
