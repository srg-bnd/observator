// Services for server
package services

import (
	"github.com/srg-bnd/observator/internal/storage"
)

type Service struct {
	storage storage.Repositories
}

func NewService(storage storage.Repositories) *Service {
	return &Service{
		storage: storage,
	}
}
