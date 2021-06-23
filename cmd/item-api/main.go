package main

import (
	"net/http"

	"github.com/Reywaltz/backend_xsolla/cmd/item-api/handlers"
	log "github.com/Reywaltz/backend_xsolla/pkg/log"
	"github.com/gorilla/mux"
)

func main() {
	log, err := log.NewLogger()
	if err != nil {
		log.Fatalf("Can't create logger: %s", err.Error())
	}
	log.Infof("Inited Logger")

	userHandlers := handlers.NewItemHandlers(log)

	router := mux.NewRouter()
	router.StrictSlash(true)

	userHandlers.Routes(router)

	http.ListenAndServe(":8000", router)
}
