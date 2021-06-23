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
	Create(item models.Item) error
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

	tmpName, tmpType := *item.Name, *item.Type
	sku := tmpName[:3] + `-` + tmpType[:3]

	item.SKU = &sku

	if err := i.ItemRepo.Create(item); err != nil {
		i.log.Errorf("Can't insert item: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	rawJSON := []byte(`{"SKU": "` + *item.SKU + `"}`)
	i.log.Infof("Created item: %v", item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(rawJSON)

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
