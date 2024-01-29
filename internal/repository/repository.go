package repository

import (
	"database/sql"

	"github.com/crhistiant7/movies-back/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
