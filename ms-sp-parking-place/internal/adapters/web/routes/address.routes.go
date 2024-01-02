package routes

import (
	"database/sql"

	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/adapters/web/handlers/address"
	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/adapters/web/server"
	_ "github.com/lib/pq"
)

func SetupAddressRoutes(s *server.WebServer, db *sql.DB) {
	webAddresHandler := address.NewWebAddressHandler(db)
	s.AddHandler("/address/create", webAddresHandler.Create)
}
