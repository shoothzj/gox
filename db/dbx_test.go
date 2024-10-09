package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDbxSqliteFromConfig(t *testing.T) {
	dbx, err := NewDbx(&Config{
		Driver: DriverSqlite,
	})
	require.NoError(t, err)
	_ = dbx.Close()
}

func TestNewDbxSqliteFromDsn(t *testing.T) {
	dbx, err := NewDbxFromDsn(DriverSqlite, ":memory:", "")
	require.NoError(t, err)
	_ = dbx.Close()
}
