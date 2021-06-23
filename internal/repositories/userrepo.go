package repositories

import (
	"github.com/Reywaltz/backend_xsolla/internal/models"
	"github.com/Reywaltz/backend_xsolla/pkg/postgres"
)

const (
	itemFields  = `SKU, name, type, cost`
	GetAllquery = `SELECT ` + itemFields + ` from item`
)

type ItemRepo struct {
	DB *postgres.DB
}

func NewItemRepository(db *postgres.DB) *ItemRepo {
	return &ItemRepo{
		DB: db,
	}
}

func (i *ItemRepo) GetAll() ([]models.Item, error) {
	rows, err := i.DB.Conn.Query(GetAllquery)
	if err != nil {
		return nil, err
	}

	out := make([]models.Item, 0)

	for rows.Next() {
		var tmpItem models.Item

		if err := rows.Scan(&tmpItem.SKU, &tmpItem.Name, &tmpItem.Type, &tmpItem.Cost); err != nil {
			return out, err
		}

		out = append(out, tmpItem)
	}

	return out, nil
}
