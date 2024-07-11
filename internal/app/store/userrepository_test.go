package store_test

import (
	"github.com/futod4m4/http-rest-api/internal/app/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t)
}
