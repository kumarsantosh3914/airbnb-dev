package db

// facilitates dependency injection for repository
type Storage struct {
	ReviewRepository ReviewRepository
}

func NewStorage() *Storage {
	return &Storage{
		ReviewRepository: &ReviewRepositoryImpl{},
	}
}
