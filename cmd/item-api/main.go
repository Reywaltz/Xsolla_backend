package main

import (
	"net/http"

	"github.com/Reywaltz/backend_xsolla/cmd/item-api/handlers"
	"github.com/Reywaltz/backend_xsolla/internal/repositories"
	log "github.com/Reywaltz/backend_xsolla/pkg/log"
	"github.com/Reywaltz/backend_xsolla/pkg/postgres"
	"github.com/gorilla/mux"
)

func main() {
	log, err := log.NewLogger()
	if err != nil {
		log.Fatalf("Can't create logger: %s", err.Error())
	}
	log.Infof("Inited Logger")

	db, err := postgres.NewDB()
	if err != nil {
		log.Fatalf("Can't init db: %s", err)
	}
	log.Infof("db: %v", db)

	itemRepo := repositories.NewItemRepository(db)
	log.Infof("Created ItemRepository")

	itemHandlers := handlers.NewItemHandlers(log, itemRepo)
	log.Infof("Inited Item handlers")

	router := mux.NewRouter()
	router.StrictSlash(true)

	itemHandlers.Routes(router)
	log.Infof("Inited router")

	log.Infof("Server is up")
	http.ListenAndServe(":8000", router)
}
