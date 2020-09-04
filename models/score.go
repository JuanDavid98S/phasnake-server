package models

type Score struct {
	ID     int64 `json:"id"`
	User   User  `json:"user"`
	Points int64 `json:"points"`
}
