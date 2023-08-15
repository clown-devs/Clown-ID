package sqlstore

import (
	"clown-id/internal/models"
	"database/sql"
)

type TokenRepository struct {
	db *sql.DB
}

func (r *TokenRepository) Find(tokenValue string) (*models.RefreshToken, error) {
	token := &models.RefreshToken{}
	err := r.db.QueryRow(`SELECT t.Token, t.app_id, t.client_id, t.expires_at  FROM refresh_tokens t WHERE t.token = $1 `, tokenValue).Scan(
		&token.Token,
		&token.AppId,
		&token.ClientId,
		&token.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return token, err
}

// Can be only one token with same app_id and client_id
func (r *TokenRepository) Create(token *models.RefreshToken) error {

	var existing_token string
	isExists := r.db.QueryRow("SELECT token FROM refresh_tokens WHERE app_id = $1 AND client_id = $2", token.AppId, token.ClientId).Scan(&existing_token)
	if isExists == nil {
		err := r.Delete(existing_token)
		if err != nil {
			return err
		}
	}

	_, err := r.db.Exec("INSERT INTO refresh_tokens(token, app_id, client_id, expires_at) VALUES ($1, $2, $3, $4)",
		&token.Token,
		&token.AppId,
		&token.ClientId,
		&token.ExpiresAt,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *TokenRepository) Delete(tokenValue string) error {
	_, err := r.db.Exec("DELETE FROM refresh_tokens WHERE token = $1", tokenValue)
	return err
}
