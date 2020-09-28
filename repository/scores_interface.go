package repository

import (
	"context"

	models "../models"
)

// ScoresInterface explain...
type ScoresInterface interface {
	Fetch(ctx context.Context, lastRn int64, limit int64) ([]*models.Scores, error)
	GetByID(ctx context.Context, id int64) (*models.Scores, error)
	Create(ctx context.Context, p *models.Scores) (string, error)
	Update(ctx context.Context, p *models.Scores) (*models.Scores, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
