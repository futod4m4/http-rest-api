package sqlstore

import (
	"database/sql"
	"github.com/futod4m4/http-rest-api/internal/app/store"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User store.User().Create()
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
