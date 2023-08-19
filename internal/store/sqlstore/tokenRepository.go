package sqlstore

import (
	"clown-id/internal/models"
	"database/sql"
	"fmt"
)

type TokenRepository struct {
	db *sql.DB
}

func (r *TokenRepository) Find(tokenValue string) (*models.RefreshToken, error) {
	token := &models.RefreshToken{}
	err := r.db.QueryRow(`SELECT t.Token, t.app_id, t.client_id, t.expires_at, t.user_id  FROM refresh_tokens t WHERE t.token = $1 `, tokenValue).Scan(
		&token.Token,
		&token.AppId,
		&token.ClientId,
		&token.ExpiresAt,
		&token.UserId,
	)

	if err != nil {
		return nil, err
	}

	return token, err
}

// Can be only one token with same user_id, app_id and client_id
func (r *TokenRepository) Create(token *models.RefreshToken) error {
	var existing_token string
	isExists := r.db.QueryRow("SELECT token FROM refresh_tokens WHERE app_id = $1 AND client_id = $2 AND user_id = $3", token.AppId, token.ClientId, token.UserId).Scan(&existing_token)
	if isExists == nil {
		err := r.Delete(existing_token)
		if err != nil {
			return err
		}

	}
	_, err := r.db.Exec("INSERT INTO refresh_tokens(token, app_id, client_id, expires_at, user_id) VALUES ($1, $2, $3, $4, $5)",
		&token.Token,
		&token.AppId,
		&token.ClientId,
		&token.ExpiresAt,
		&token.UserId,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *TokenRepository) Exists(tokenValue string) bool {
	var found bool
	if err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM refresh_tokens WHERE token = $1)", tokenValue).Scan(&found); err != nil {
		return false
	}
	return found
}

func (r *TokenRepository) Delete(tokenValue string) error {
	if !r.Exists(tokenValue) {
		return fmt.Errorf("cannot delete token, this token doesn't exists: %s", tokenValue)
	}
	_, err := r.db.Exec("DELETE FROM refresh_tokens WHERE token = $1", tokenValue)
	return err
}
