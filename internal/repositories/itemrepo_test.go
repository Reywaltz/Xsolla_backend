package repositories_test

import (
	"testing"

	"github.com/Reywaltz/backend_xsolla/cmd/item-api/additions"
	"github.com/Reywaltz/backend_xsolla/cmd/item-api/handlers"
	"github.com/Reywaltz/backend_xsolla/internal/models"
	"github.com/Reywaltz/backend_xsolla/internal/repositories"
	"github.com/Reywaltz/backend_xsolla/pkg/log"
	"github.com/pashagolub/pgxmock"
	"github.com/shopspring/decimal"
)

type BadQuery struct {
	MaxCost string
}

func initMockItemHandler() (pgxmock.PgxConnIface, *handlers.ItemHandlers, error) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		return nil, nil, err
	}

	repo := repositories.NewItemRepository(mock)
	log, _ := log.NewLogger()
	itemHandler := handlers.NewItemHandlers(log, repo)

	return mock, itemHandler, nil
}

func TestGetAllItemsQueryGood(t *testing.T) {
	t.Parallel()

	mock, itemHandlers, err := initMockItemHandler()
	if err != nil {
		t.Fatalf("Can't init Mock handler: %v", err)
	}

	rows := pgxmock.NewRows([]string{"sku", "name", "type", "cost"}).
		AddRow("POR-GAM", "Portal 2", "Game", decimal.Decimal{}).
		AddRow("DOT-SUB", "Dota", "Subscription", decimal.Decimal{})

	query := &additions.Query{
		Limit:   nil,
		Offset:  nil,
		Type:    additions.DefaultType,
		MinCost: additions.DefaultMinCost,
		MaxCost: additions.DefaultMaxCost,
	}

	mock.ExpectQuery(repositories.GetAllquery).
		WithArgs(query.Type, query.MinCost, query.MaxCost, query.Limit, query.Offset).
		WillReturnRows(rows)

	_, err = itemHandlers.ItemRepo.GetAll(query)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not met: %s", err)
	}
}

func TestGetAllItemsQueryBad(t *testing.T) {
	t.Parallel()

	mock, itemHandlers, err := initMockItemHandler()
	if err != nil {
		t.Fatalf("Can't init Mock handler: %v", err)
	}

	query := &additions.Query{
		Limit:   nil,
		Offset:  nil,
		Type:    additions.DefaultType,
		MinCost: "SomeString",
		MaxCost: additions.DefaultMaxCost,
	}

	mock.ExpectQuery(repositories.GetAllquery).
		WithArgs(query.Type, query.MinCost, query.MaxCost, query.Limit, query.Offset).
		WillReturnError(err)

	_, err = itemHandlers.ItemRepo.GetAll(query)
	if err != nil {
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not met: %s", err)
	}
}

func TestGetOneQuery(t *testing.T) {
	t.Parallel()
	mock, itemHandlers, err := initMockItemHandler()
	if err != nil {
		t.Fatalf("Can't init Mock handler: %v", err)
	}

	rows := pgxmock.NewRows([]string{"sku", "name", "type", "cost"}).AddRow("POR-GAM", "Portal 2", "Game", decimal.Decimal{})

	mock.ExpectQuery(repositories.GetAllquery).WithArgs("POR-GAM").WillReturnRows(rows)

	item := models.Item{
		SKU: "POR-GAM",
	}

	_, err = itemHandlers.ItemRepo.GetOne(item)
	if err != nil {
		t.Errorf("Actual result differs from expected: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not met: %s", err)
	}
}

func TestCreateQuery(t *testing.T) {
	t.Parallel()
	mock, itemHandlers, err := initMockItemHandler()
	if err != nil {
		t.Fatalf("Can't init Mock handler: %v", err)
	}

	res := pgxmock.NewResult("INSERT", 1)

	validItem := models.Item{
		SKU:  "MET-GAM",
		Name: "MetalGearSolid",
		Type: "Game",
		Cost: decimal.Decimal{},
	}
	mock.ExpectExec("INSERT INTO item").
		WithArgs(validItem.SKU, validItem.Name, validItem.Type, validItem.Cost).
		WillReturnResult(res)

	err = itemHandlers.ItemRepo.Create(validItem)
	if err != nil {
		t.Errorf("Actual result differs from expected: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not met: %s", err)
	}
}

func TestDeleteQuery(t *testing.T) {
	t.Parallel()
	mock, itemHandlers, err := initMockItemHandler()
	if err != nil {
		t.Fatalf("Can't init Mock handler: %v", err)
	}

	row := pgxmock.NewRows([]string{"sku"}).AddRow("MET-GAM")

	validItem := models.Item{
		SKU: "MET-GAM",
	}

	mock.ExpectQuery("DELETE FROM item").
		WithArgs(validItem.SKU).
		WillReturnRows(row)

	_, err = itemHandlers.ItemRepo.Delete(validItem)
	if err != nil {
		t.Errorf("Actual result differs from expected: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not met: %s", err)
	}
}

func TestUpdateQuery(t *testing.T) {
	t.Parallel()
	mock, itemHandlers, err := initMockItemHandler()
	if err != nil {
		t.Fatalf("Can't init Mock handler: %v", err)
	}

	validItem := models.Item{
		SKU:  "TEST-GAM",
		Name: "Test",
		Type: "Game",
		Cost: decimal.Decimal{},
	}

	res := pgxmock.NewResult("UPDATE", 1)

	mock.ExpectExec("UPDATE item set").
		WithArgs(validItem.Name, validItem.Type, validItem.Cost, validItem.SKU).
		WillReturnResult(res)

	err = itemHandlers.ItemRepo.Update(validItem)
	if err != nil {
		t.Errorf("Actual result differs from expected: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not met: %s", err)
	}
}
