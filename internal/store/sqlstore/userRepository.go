package sqlstore

import (
	"clown-id/internal/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

// Returns only public fields (without password)
func (r *UserRepository) Find(id string) (*models.User, error) {
	return r.findByField("id", id)
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	return r.findByField("email", email)
}

func (r *UserRepository) findByField(fieldName, fieldValue string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(fmt.Sprintf(`SELECT u.id, u.username, u.email FROM users u WHERE u.%s = $1 `, fieldName), fieldValue).Scan(
		&user.ID,
		&user.Username,
		&user.Email)

	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}
