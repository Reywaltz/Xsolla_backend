package postgres_test

import (
	"testing"

	"github.com/Reywaltz/backend_xsolla/pkg/postgres"
)

func TestConnToDatabase(t *testing.T) {
	t.Parallel()

	if _, err := postgres.NewDB(); err != nil {
		t.Fatalf("Can't get connection :%v", err)
	}
}
