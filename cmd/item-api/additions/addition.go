package additions

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Query struct {
	Limit   *string
	Offset  *string
	Type    string
	MinCost string
	MaxCost string
}

const (
	DefaultType    = `%%`
	DefaultMinCost = `0`
	DefaultMaxCost = `4294967295`
)

func (q *Query) HandleURLQueries(r *http.Request) error {
	query := r.URL.Query()
	limit, err := handleLimit(query)
	if err != nil {
		return err
	}
	q.Limit = limit

	offset, err := handleOffset(query)
	if err != nil {
		return err
	}
	q.Offset = offset

	itemType := handleType(query)

	q.Type = itemType

	minCost, err := handleMinCost(query)
	if err != nil {
		return err
	}
	q.MinCost = minCost

	maxCost, err := handleMaxCost(query)
	if err != nil {
		return err
	}
	q.MaxCost = maxCost

	return nil
}

func handleLimit(query url.Values) (*string, error) {
	if limit := query.Get("limit"); limit != "" {
		value, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		if value < 0 {
			return nil, errors.New("Limit can't be negative")
		}

		return &limit, nil
	}

	return nil, nil
}

func handleOffset(query url.Values) (*string, error) {
	if offset := query.Get("offset"); offset != "" {
		value, err := strconv.Atoi(offset)
		if err != nil {
			return nil, err
		}
		if value < 0 {
			return nil, errors.New("Offset can't be negative")
		}

		return &offset, nil
	}

	return nil, nil
}

func handleType(query url.Values) string {
	filterType := strings.TrimSpace(query.Get("type"))
	if filterType == "" {
		return DefaultType
	}

	return filterType
}

func handleMinCost(query url.Values) (string, error) {
	if min := query.Get("min"); min != "" {
		value, err := strconv.Atoi(min)
		if err != nil {
			return min, err
		}
		if value < 0 {
			return min, errors.New("Min cost can't be negative")
		}

		return min, nil
	}

	return DefaultMinCost, nil
}

func handleMaxCost(query url.Values) (string, error) {
	if max := query.Get("max"); max != "" {
		value, err := strconv.Atoi(max)
		if err != nil {
			return max, err
		}
		if value < 0 {
			return max, errors.New("Max cost can't be negative")
		}

		return max, nil
	}

	return DefaultMaxCost, nil
}
