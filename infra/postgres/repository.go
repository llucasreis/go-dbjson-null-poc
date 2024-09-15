package postgres

import (
	"database/sql"
	"log"

	"github.com/llucasreis/go-dbjson-null-poc/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByID(id int64) (*domain.Game, error) {
	query := `select id, title, metadata from games where id = $1`

	row := r.db.QueryRow(query, id)

	var model Game
	err := row.Scan(&model.ID, &model.Title, &model.Metadata)
	if err != nil {
		log.Println("[ERROR] finding game by id: ", err.Error())
		return nil, err
	}
	return model.ToDomain()
}
