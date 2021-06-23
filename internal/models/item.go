package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	numeric "github.com/jackc/pgtype/ext/shopspring-numeric"
)

type Item struct {
	SKU  *string         `json:"sku"`
	Name *string         `json:"name"`
	Type *string         `json:"type"`
	Cost numeric.Numeric `json:"cost"`
}

func (i *Item) Bind(r *http.Request) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &i); err != nil {
		return err
	}

	if i.SKU != nil {
		return errors.New("SKU will be generated on create")
	}

	if i.Name == nil {
		return errors.New("Name is required")
	}

	if i.Type == nil {
		return errors.New("Type is required")
	}

	return nil
}
