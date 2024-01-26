package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate string    `json:"release_date"`
	RunTime     int       `json:"runtime"`
	MPPARating  string    `json:"mpaa_rating"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdateAt    time.Time `json:"-"`
}
