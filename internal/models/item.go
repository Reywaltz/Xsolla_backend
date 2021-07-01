package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type Item struct {
	SKU  string          `json:"sku"`
	Name string          `json:"name"`
	Type string          `json:"type"`
	Cost decimal.Decimal `json:"cost"`
}

func (i *Item) Bind(r *http.Request) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &i); err != nil {
		return err
	}

	if i.SKU != "" {
		return errors.New("SKU will be generated on create")
	}

	if i.Name == "" {
		return errors.New("Name is required")
	}

	if i.Type == "" {
		return errors.New("Type is required")
	}

	var tmp decimal.Decimal
	if i.Cost == tmp {
		return errors.New("Cost is required")
	}

	_, err = strconv.Atoi(i.Cost.String())
	if err != nil {
		return err
	}

	if i.Cost.IsNegative() {
		return errors.New("Cost can't be negative")
	}

	if len(strings.TrimSpace(i.Name)) < 3 || len(strings.TrimSpace(i.Type)) < 3 {
		return errors.New("Name or type are too short")
	}

	return nil
}
