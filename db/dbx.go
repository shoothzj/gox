package db

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Driver   Driver
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type Dbx struct {
	driver Driver
	user   string
	db     *sql.DB
}

func (d *Dbx) Driver() Driver {
	return d.driver
}

func (d *Dbx) User() string {
	return d.user
}

func (d *Dbx) Db() *sql.DB {
	return d.db
}

func (d *Dbx) Ping() error {
	return d.db.Ping()
}

func (d *Dbx) Close() error {
	return d.db.Close()
}

func NewDbx(config *Config) (*Dbx, error) {
	var dsn string
	switch config.Driver {
	case DriverMySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.DbName)
	case DriverSqlite:
		dsn = ":memory:"
	default:
		return nil, fmt.Errorf("unsupported Driver: %s", config.Driver)
	}

	db, err := sql.Open(string(config.Driver), dsn)
	if err != nil {
		return nil, err
	}

	return &Dbx{
		driver: config.Driver,
		user:   config.User,
		db:     db,
	}, nil
}

func NewDbxFromDsn(driver Driver, dsn string, user string) (*Dbx, error) {
	db, err := sql.Open(string(driver), dsn)
	if err != nil {
		return nil, err
	}

	return &Dbx{
		driver: driver,
		user:   user,
		db:     db,
	}, nil
}

func NewDbxFromDb(db *sql.DB, driver Driver, user string) *Dbx {
	return &Dbx{
		driver: driver,
		user:   user,
		db:     db,
	}
}
