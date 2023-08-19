package store

import "clown-id/internal/models"

type UserRepository interface {
	Find(string) (*models.User, error) // ID
	FindByEmail(string) (*models.User, error)
	FindByUsername(string) (*models.User, error)
	Create(*models.User) error
}

type TokenRepository interface {
	Find(string) (*models.RefreshToken, error)
	Create(*models.RefreshToken) error
	Delete(string) error
}

type ClientRepository interface {
	AllApps() ([]models.Application, error)
	AllClients() ([]models.Client, error)
}
