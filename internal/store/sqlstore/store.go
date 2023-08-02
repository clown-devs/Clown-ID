package sqlstore

import (
	"clown-id/internal/store"
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB

	UserRepository *UserRepository
}

func New(connStr string) (*Store, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil { // database ping
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository == nil {
		s.UserRepository = &UserRepository{
			db: s.db,
		}
	}
	return s.UserRepository
}
