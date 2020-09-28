package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "../models"
	utils "../utils"
)

// NewSQLScores ..
func NewSQLScores(db *sql.DB) ScoresInterface {
	return &ScoresRepository{
		Conn: db,
	}
}

// ScoresRepository ..
type ScoresRepository struct {
	Conn *sql.DB
}

func (sr *ScoresRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Scores, error) {
	rows, err := sr.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Scores, 0)
	for rows.Next() {
		data := new(models.Scores)

		err := rows.Scan(&data.ID, &data.Nickname, &data.Score, &data.CreatedAt)
		if err != nil {
			return nil, err
		}
		fmt.Printf("data: %v", data)
		payload = append(payload, data)
	}

	return payload, nil
}

// Fetch ..
func (sr *ScoresRepository) Fetch(ctx context.Context, lastRn int64, limit int64) ([]*models.Scores, error) {
	query := "SELECT id, nickname, score, created_at FROM scores_ordered WHERE rn > $1 ORDER BY score DESC, created_at ASC LIMIT $2"

	return sr.fetch(ctx, query, lastRn, limit)
}

// GetByID ..
func (sr *ScoresRepository) GetByID(ctx context.Context, id int64) (*models.Scores, error) {
	query := "SELECT id, nickname, score, created_at FROM scores where id = $1"

	rows, err := sr.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Scores{}
	if len(rows) > 0 {
		payload = rows[0]
		fmt.Printf("data: %v", payload)
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

// Create ..
func (sr *ScoresRepository) Create(ctx context.Context, p *models.Scores) (string, error) {
	query := "INSERT INTO scores (id, nickname, score, created_at) VALUES ($1, $2, $3, current_timestamp)"
	ID := utils.GenerateUUID()

	stmt, err := sr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return "", err
	}

	_, err = stmt.ExecContext(ctx, ID, p.Nickname, p.Score)
	defer stmt.Close()
	if err != nil {
		return "", err
	}

	return ID, nil
}

// Update ..
func (sr *ScoresRepository) Update(ctx context.Context, p *models.Scores) (*models.Scores, error) {
	query := "UPDATE scores SET nickname = $2, score = $3 WHERE id = $1"

	stmt, err := sr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(ctx, p.ID, p.Nickname, p.Score)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Delete ..
func (sr *ScoresRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM scores Where id = $1"

	stmt, err := sr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
