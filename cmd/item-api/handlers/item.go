package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Reywaltz/backend_xsolla/cmd/item-api/additions"
	"github.com/Reywaltz/backend_xsolla/internal/models"
	log "github.com/Reywaltz/backend_xsolla/pkg/log"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type ItemRepository interface {
	GetAll(*additions.Query) ([]models.Item, error)
	Create(item models.Item) error
	Delete(item models.Item) (*string, error)
	GetOne(item models.Item) (models.Item, error)
	Update(item models.Item) error
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

func (i *ItemHandlers) getItems(w http.ResponseWriter, r *http.Request) {
	queries, err := additions.HandleURLQueries(r)
	if err != nil {
		i.log.Errorf("Wrong query param: %s", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := i.ItemRepo.GetAll(queries)
	if err != nil {
		i.log.Errorf("Can't get data from DB: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
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
}

func (i *ItemHandlers) getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sku, ok := vars["sku"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var item models.Item
	item.SKU = &sku

	res, err := i.ItemRepo.GetOne(item)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			i.log.Errorf("Can't get item with sku: %s", err)
			w.WriteHeader(http.StatusNotFound)

			return
		} else {
			i.log.Errorf("Can't get item: %s", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
	out, err := json.Marshal(&res)
	if err != nil {
		i.log.Errorf("Can't marshall json", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (i *ItemHandlers) createItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	if err := item.Bind(r); err != nil {
		i.log.Errorf("Can't bind JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tmpName, tmpType := *item.Name, *item.Type
	sku := strings.ToUpper(tmpName[:3] + `-` + tmpType[:3])

	item.SKU = &sku

	if err := i.ItemRepo.Create(item); err != nil {
		i.log.Errorf("Can't insert item: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	rawJSON := []byte(`{"sku": "` + *item.SKU + `"}`)
	i.log.Infof("Created item: %s", item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(rawJSON)
}

func (i *ItemHandlers) deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sku, ok := vars["sku"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var item models.Item
	item.SKU = &sku

	res, err := i.ItemRepo.Delete(item)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			i.log.Errorf("No such item: %s", err)
			w.WriteHeader(http.StatusNotFound)

			return
		} else {
			i.log.Fatalf("Can't delete item: %s", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	i.log.Infof("Item with SKU={%s} is deleted", *res)
	w.WriteHeader(http.StatusNoContent)
}

func (i *ItemHandlers) EditItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sku, ok := vars["sku"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var tmp models.Item
	if err := tmp.Bind(r); err != nil {
		i.log.Errorf("Can't bind JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tmp.SKU = &sku

	if err := i.ItemRepo.Update(tmp); err != nil {
		i.log.Errorf("Can't update item with sku: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (i *ItemHandlers) Routes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/items", i.createItem).Methods("POST")
	subRouter.HandleFunc("/items", i.getItems).Methods("GET")
	subRouter.HandleFunc("/items/{sku}", i.deleteItem).Methods("DELETE")
	subRouter.HandleFunc("/items/{sku}", i.getItem).Methods("GET")
	subRouter.HandleFunc("/items/{sku}", i.EditItem).Methods("PUT")
}
