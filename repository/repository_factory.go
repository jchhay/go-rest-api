package repository

type RepositoryFactory interface {
	NewBookRepository() BookRepository
}
