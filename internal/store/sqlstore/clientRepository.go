package sqlstore

import (
	"clown-id/internal/models"
	"database/sql"
)

type ClientRepository struct {
	db *sql.DB
}

func (r *ClientRepository) AllApps() ([]models.Application, error) {
	apps := make([]models.Application, 0, 10)

	rows, err := r.db.Query("SELECT id, name FROM apps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		app := models.Application{}
		err := rows.Scan(&app.ID, &app.Name)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return apps, nil
}

func (r *ClientRepository) AllClients() ([]models.Client, error) {
	clients := make([]models.Client, 0, 10)

	rows, err := r.db.Query("SELECT id, name FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		client := models.Client{}
		err := rows.Scan(&client.ID, &client.Name)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}
