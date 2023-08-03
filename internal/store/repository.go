package store

import "clown-id/internal/models"

type UserRepository interface {
	Find(string) (*models.User, error) // ID
	FindByEmail(string) (*models.User, error)
	Create(*models.User) error
}
