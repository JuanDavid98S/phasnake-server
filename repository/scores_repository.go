package repository

import (
	"context"
	"database/sql"

	models "../models"
	utils "../utils"
)

// ScoresRepository is the type for scores repository
type ScoresRepository struct {
	Conn *sql.DB
}

// NewSQLScores is the scores repository constructor
func NewSQLScores(db *sql.DB) ScoresInterface {
	return &ScoresRepository{
		Conn: db,
	}
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
		payload = append(payload, data)
	}

	return payload, nil
}

// Fetch is a method to get top scores from the database
func (sr *ScoresRepository) Fetch(ctx context.Context, lastRn int64, limit int64) ([]*models.Scores, error) {
	query := "SELECT id, nickname, score, created_at FROM scores_ordered WHERE rn > $1 ORDER BY score DESC, created_at ASC LIMIT $2"

	return sr.fetch(ctx, query, lastRn, limit)
}

// GetByID is a method to get a single element from database
func (sr *ScoresRepository) GetByID(ctx context.Context, id string) (*models.Scores, error) {
	query := "SELECT id, nickname, score, created_at FROM scores where id = $1"

	rows, err := sr.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Scores{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

// Create is a method to insert a new score in database
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

// Update is a method to update a row in database
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

// Delete is a method to delete a row in database
func (sr *ScoresRepository) Delete(ctx context.Context, id string) (bool, error) {
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
