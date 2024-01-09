package factory

import (
	"jchhay/go-rest-api-gin/internal/repository"
	"jchhay/go-rest-api-gin/internal/repository/memory"
	"jchhay/go-rest-api-gin/internal/repository/sqlite"
)

func NewRepositoryFactory(dbType string) repository.BookRepository {
	if dbType == "sqlite" {
		return sqlite.NewBookRepository(dbType)
	}
	return memory.NewBookRepository()
}
