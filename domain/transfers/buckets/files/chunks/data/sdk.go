package data

import (
	"github.com/xmn-services/rod-network/libs/file"
	"github.com/xmn-services/rod-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	fileService file.Service,
) Service {
	hashAdapter := hash.NewAdapter()
	return createService(hashAdapter, fileService)
}

// NewRepository creates a new repository instance
func NewRepository(fileRepository file.Repository) Repository {
	return createRepository(fileRepository)
}

// Repository represents a data repository
type Repository interface {
	Retrieve(hash hash.Hash) ([]byte, error)
}

// Service represents a data service
type Service interface {
	Save(data []byte) error
	Delete(hash hash.Hash) error
}
