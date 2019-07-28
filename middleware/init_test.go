package nps

import (
	"testing"
)

func TestInit(t *testing.T) {
	tables := []string{}
	DB.Select(&tables, "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname='public';")
	t.Logf("Number of tables found in test: %d", len(tables))

	if DB.DriverName() != "postgres" {
		t.Errorf("Wrong driver")
	}
}
