package models

// Scores is the type for scores model
type Scores struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Score     int64  `json:"score"`
	CreatedAt string `json:"created_at"`
}
