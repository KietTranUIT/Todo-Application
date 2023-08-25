package repository

import (
	"database/sql"
	"user-service/conf"
	"user-service/internal/core/port/repository"
)

type database struct {
	db *sql.DB
}

func NewDB(config conf.ConfigDatabase) (repository.Database, error) {
	db, err := newDatabase(config)
	return database{
		db: db,
	}, err
}

func newDatabase(config conf.ConfigDatabase) (*sql.DB, error) {
	db, err := sql.Open(config.Driver, config.GetURL())

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db database) Close() error {
	return db.db.Close()
}

func (db database) GetDB() *sql.DB {
	return db.db
}
