package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/configs"
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

	runDBMigration(cfg.MigrationURL, cfg.DBSource)

	defer db.Close()

	println("connected on database", db.Ping())

	webserver := server.NewWebServer(":8086")
	routes.SetupAddressRoutes(webserver, db)

	webserver.Start()

}

func runDBMigration(migrationURL string, dbSource string) {
	time.Sleep(10 * time.Second)
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance: ", err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up: ", err)
	}
	log.Println("db migrate successfully")
}
