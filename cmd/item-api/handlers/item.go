package handlers

import (
	"net/http"

	"github.com/Reywaltz/backend_xsolla/internal/models"
	log "github.com/Reywaltz/backend_xsolla/pkg/log"
	"github.com/gorilla/mux"
)

type ItemRepository interface {
	GetAll() []models.Item
	CreateItem(item models.Item) error
}

type ItemHandlers struct {
	log      log.Logger
	ItemRepo ItemRepository
}

func NewItemHandlers(log log.Logger) *ItemHandlers {
	return &ItemHandlers{
		log: log,
	}
}

func (i *ItemHandlers) createItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	if err := item.Bind(r); err != nil {
		i.log.Errorf("Can't bind JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	i.log.Infof("Incomed item: %v", item)
	w.WriteHeader(http.StatusOK)

	return
}

func (i *ItemHandlers) Routes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/items", i.createItem).Methods("POST")
}
