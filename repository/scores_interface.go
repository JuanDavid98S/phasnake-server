package repository

import (
	"context"

	models "../models"
)

// ScoresInterface is the type for scores interface
type ScoresInterface interface {
	Fetch(ctx context.Context, lastRn int64, limit int64) ([]*models.Scores, error)
	GetByID(ctx context.Context, id string) (*models.Scores, error)
	Create(ctx context.Context, p *models.Scores) (string, error)
	Update(ctx context.Context, p *models.Scores) (*models.Scores, error)
	Delete(ctx context.Context, id string) (bool, error)
}
