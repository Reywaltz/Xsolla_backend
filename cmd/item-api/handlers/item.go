package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Reywaltz/backend_xsolla/internal/models"
	log "github.com/Reywaltz/backend_xsolla/pkg/log"
	"github.com/gorilla/mux"
)

type ItemRepository interface {
	GetAll() ([]models.Item, error)
	// createItem(item models.Item) error
}

type ItemHandlers struct {
	log      log.Logger
	ItemRepo ItemRepository
}

func NewItemHandlers(log log.Logger, itemRepo ItemRepository) *ItemHandlers {
	return &ItemHandlers{
		log:      log,
		ItemRepo: itemRepo,
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

func (i *ItemHandlers) getItems(w http.ResponseWriter, r *http.Request) {
	res, err := i.ItemRepo.GetAll()
	if err != nil {
		i.log.Errorf("Can't get data from DB: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	data, err := json.Marshal(&res)
	if err != nil {
		i.log.Errorf("Can't marshall data: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	return
}

func (i *ItemHandlers) Routes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/items", i.createItem).Methods("POST")
	subRouter.HandleFunc("/items", i.getItems).Methods("GET")

}
