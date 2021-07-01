package models_test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/Reywaltz/backend_xsolla/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestBindItem(t *testing.T) {
	type testCase struct {
		Name        string
		In          []byte
		ExpectedErr bool
	}

	const (
		URL = `/items`
	)

	ItemWithSKU := `{"sku": "SOM-TAG","name": "Videogame #1","type": "Game","cost": "12"}`

	ItemWithoutName := `{"type": "Game","cost": "12"}`

	ItemWithoutType := `{"name": "Videogame #1","cost": "12"}`

	ItemWithoutCost := `{"name": "Videogame #1","type": "Game"}`

	ItemWithCostNegative := `{"name": "Videogame #1","type": "Game", "cost":"-12"}`

	ItemWithCostLiteral := `{"name": "Videogame #1","type": "Game", "cost":"zxc"}`

	InvalidJson := `"name": "Videogame #1","type": "Game"}`

	ItemWithShortName := `{"name": "Vi","type": "Game","cost": "12"}`

	ItemWithShortType := `{"name": "Videogame #1","type": "Ga","cost": "12"}`

	ItemWithWrongCost := `{"name": "Videogame #1","type": "Game","cost": "1z2"}`

	ValidItem := `{"name": "Videogame #1","type": "Game","cost": "12"}`

	Testcases := []testCase{
		{Name: "Item with SKU", In: []byte(ItemWithSKU), ExpectedErr: true},
		{Name: "Item without name", In: []byte(ItemWithoutName), ExpectedErr: true},
		{Name: "Item without type", In: []byte(ItemWithoutType), ExpectedErr: true},
		{Name: "Item without cost", In: []byte(ItemWithoutCost), ExpectedErr: true},
		{Name: "Item with negative cost", In: []byte(ItemWithCostNegative), ExpectedErr: true},
		{Name: "Item with literal cost", In: []byte(ItemWithCostLiteral), ExpectedErr: true},
		{Name: "Invalid JSON", In: []byte(InvalidJson), ExpectedErr: true},
		{Name: "Item with short name", In: []byte(ItemWithShortName), ExpectedErr: true},
		{Name: "Item with short type", In: []byte(ItemWithShortType), ExpectedErr: true},
		{Name: "Item with wrong cost", In: []byte(ItemWithWrongCost), ExpectedErr: true},
		{Name: "Valid item", In: []byte(ValidItem), ExpectedErr: false},
	}
	for _, tc := range Testcases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			reader := bytes.NewBuffer(tc.In)
			r := httptest.NewRequest("GET", URL, reader)
			var tmp models.Item

			err := tmp.Bind(r)
			assert.Equal(t, tc.ExpectedErr, err != nil)
		})
	}
}
