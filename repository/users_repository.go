package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "../models"
)

// NewSQLUserRepository retunrs implement of post repository interface
func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &roachUserRepository{
		Conn: db,
	}
}

type roachUserRepository struct {
	Conn *sql.DB
}

func (m *roachUserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(&data.ID, &data.Nick)
		if err != nil {
			return nil, err
		}
		fmt.Printf("data: %v", data)
		payload = append(payload, data)
	}

	return payload, nil
}

func (m *roachUserRepository) getNewID() int64 {
	lastID := int64(0)
	m.Conn.QueryRow("SELECT id FROM users ORDER BY id DESC LIMIT 1").Scan(&lastID)
	return lastID + 1
}

func (m *roachUserRepository) Fetch(ctx context.Context, num int64) ([]*models.User, error) {
	query := "SELECT id, nick FROM users LIMIT $1"

	return m.fetch(ctx, query, num)
}

func (m *roachUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := "SELECT id, nick FROM users where id = $1"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.User{}
	if len(rows) > 0 {
		payload = rows[0]
		fmt.Printf("data: %v", payload)
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *roachUserRepository) Create(ctx context.Context, p *models.User) (int64, error) {
	query := "INSERT INTO users (id, nick) VALUES ($1, $2)"
	lastID := m.getNewID()

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	_, err = stmt.ExecContext(ctx, lastID, p.Nick)
	defer stmt.Close()
	if err != nil {
		return -1, err
	}

	return lastID, nil
}

func (m *roachUserRepository) Update(ctx context.Context, p *models.User) (*models.User, error) {
	query := "UPDATE users SET nick = $1 WHERE id = $2"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(ctx, p.Nick, p.ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *roachUserRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM users Where id = $1"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
