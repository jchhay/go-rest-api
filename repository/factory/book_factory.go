package factory

import (
	"jchhay/go-rest-api-gin/repository"
	"jchhay/go-rest-api-gin/repository/memory"
	"jchhay/go-rest-api-gin/repository/sqlite"
)

func NewRepositoryFactory(dbType string) repository.BookRepository {
	if dbType == "sqlite" {
		return sqlite.NewBookRepository()
	}
	return memory.NewBookRepository()
}
