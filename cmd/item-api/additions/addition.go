package additions

import (
	"errors"
	"net/http"
	"strconv"
)

type Query struct {
	Limit  *string
	Offset *string
}

func HandleURLQueries(r *http.Request) (*Query, error) {
	var queries Query
	query := r.URL.Query()

	if limit := query.Get("limit"); limit != "" {
		value, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		if value < 0 {
			return nil, errors.New("Limit can't be negative")
		}
		queries.Limit = &limit
	}

	if offset := query.Get("offset"); offset != "" {
		value, err := strconv.Atoi(offset)
		if err != nil {
			return nil, err
		}
		if value < 0 {
			return nil, errors.New("Offset can't be negative")
		}

		queries.Offset = &offset
	}

	return &queries, nil
}
