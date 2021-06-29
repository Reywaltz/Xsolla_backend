package models_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Reywaltz/backend_xsolla/internal/models"
)

func TestWrongBindWithSKU(t *testing.T) {
	t.Parallel()

	var testItem models.Item
	jsonWithSKU := `{
		"sku": "SOM-TAG",
		"name": "Videogame #1",
		"type": "Game",
		"cost": "12"
	}`

	reader := strings.NewReader(jsonWithSKU)

	r := httptest.NewRequest("GET", "/items", reader)

	if err := testItem.Bind(r); err == nil {
		t.Error("Item binded with SKU from incomed JSON")
	}
}

func TestWrongBindWithoutNameInRequest(t *testing.T) {
	t.Parallel()

	var testItem models.Item
	JSONwithoutName := `{
		"type": "Game",
		"cost": "12"
	}`

	reader := strings.NewReader(JSONwithoutName)

	r := httptest.NewRequest("GET", "/items", reader)

	if err := testItem.Bind(r); err == nil {
		t.Error("Item binded without name from incomed JSON")
	}
}

func TestWrongBindWithoutTypeInRequest(t *testing.T) {
	t.Parallel()

	var testItem models.Item
	JSONWithoutType := `{
		"name": "Videogame #1",
		"cost": "12"
	}`

	reader := strings.NewReader(JSONWithoutType)

	r := httptest.NewRequest("GET", "/items", reader)

	if err := testItem.Bind(r); err == nil {
		t.Error("Item binded without type from incomed JSON")
	}
}

func TestWrongJSONInRequest(t *testing.T) {
	t.Parallel()

	var testItem models.Item
	InvalidJSON := `{
		"name": "Videogame #1",
		"cost": "12"
	`

	reader := strings.NewReader(InvalidJSON)

	r := httptest.NewRequest("GET", "/items", reader)

	if err := testItem.Bind(r); err == nil {
		t.Error("Binded with invalid JSON")
	}
}

func TestNameTrim(t *testing.T) {
	t.Parallel()

	var testItem models.Item
	jsonWithShortName := `{
		"name": "Vi",
		"type": "Game",
		"cost": "12"
	}`

	reader := strings.NewReader(jsonWithShortName)

	r := httptest.NewRequest("GET", "/items", reader)

	if err := testItem.Bind(r); err == nil {
		t.Error("Item binded with short name from incomed JSON")
	}
}

func TestSKUTrim(t *testing.T) {
	t.Parallel()

	var testItem models.Item
	JSONwithShortName := `{
		"name": "Vi",
		"type": "Game",
		"cost": "12"
	}`

	reader := strings.NewReader(JSONwithShortName)

	r := httptest.NewRequest("GET", "/items", reader)

	if err := testItem.Bind(r); err == nil {
		t.Error("Item binded with short name from incomed JSON")
	}
}

func TestValidJSON(t *testing.T) {
	t.Parallel()

	var testItem models.Item
	JSONwithShortName := `{
		"name": "Mega New Game",
		"type": "Game",
		"cost": "12"
	}`

	reader := strings.NewReader(JSONwithShortName)

	r := httptest.NewRequest("GET", "/items", reader)

	if err := testItem.Bind(r); err != nil {
		t.Errorf("Can't bind JSON: %v", err)
	}
}
