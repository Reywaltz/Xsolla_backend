package repositories

import (
	"context"
	"errors"

	"github.com/Reywaltz/backend_xsolla/cmd/item-api/additions"
	"github.com/Reywaltz/backend_xsolla/internal/models"
	"github.com/Reywaltz/backend_xsolla/pkg/postgres"
)

const (
	itemFields  = `sku, name, type, cost`
	GetAllquery = `SELECT ` + itemFields + ` from item where type ILIKE $1 || '%' limit $2 offset $3`
	GetOneQuery = `SELECT ` + itemFields + ` from item WHERE sku=$1`
	InsertQuery = `INSERT INTO item ( ` + itemFields + `) 
	VALUES ($1, $2, $3, $4) returning sku`
	DeleteQuery = `DELETE FROM item WHERE sku=$1 RETURNING SKU`
	UpdateQuery = `UPDATE item set name=$1, type=$2, cost=$3 WHERE sku=$4`
)

type ItemRepo struct {
	DB *postgres.DB
}

func NewItemRepository(db *postgres.DB) *ItemRepo {
	return &ItemRepo{
		DB: db,
	}
}

func (i *ItemRepo) GetAll(queries *additions.Query) ([]models.Item, error) {
	rows, err := i.DB.Conn.Query(context.Background(),
		GetAllquery,
		queries.Type,
		queries.Limit,
		queries.Offset)
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

func (i *ItemRepo) GetOne(item models.Item) (models.Item, error) {
	var tmp models.Item

	row := i.DB.Conn.QueryRow(context.Background(), GetOneQuery, *item.SKU)

	if err := row.Scan(&tmp.SKU, &tmp.Name, &tmp.Type, &tmp.Cost); err != nil {
		return models.Item{}, err
	}

	return tmp, nil
}

func (i *ItemRepo) Create(item models.Item) error {
	_, err := i.DB.Conn.Exec(context.Background(),
		InsertQuery,
		item.SKU,
		item.Name,
		item.Type,
		item.Cost,
	)
	if err != nil {
		return err
	}

	return nil
}

func (i *ItemRepo) Delete(item models.Item) (*string, error) {
	row := i.DB.Conn.QueryRow(context.Background(),
		DeleteQuery, item.SKU)

	var sku *string

	if err := row.Scan(&sku); err != nil {
		return nil, err
	}

	return sku, nil
}

func (i *ItemRepo) Update(item models.Item) error {
	commandTag, err := i.DB.Conn.Exec(context.Background(),
		UpdateQuery,
		item.Name,
		item.Type,
		item.Cost,
		item.SKU)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("No item with such sku")
	}

	return nil
}
