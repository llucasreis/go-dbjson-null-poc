package postgres

import (
	"database/sql"
	"encoding/json"

	"github.com/llucasreis/go-dbjson-null-poc/domain"
)

type Game struct {
	ID          int64                     `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Metadata    sql.Null[json.RawMessage] `json:"metadata"`
}

func (g *Game) ToDomain() (*domain.Game, error) {
	metadata, err := g.getMetadata()
	if err != nil {
		return nil, err
	}
	return &domain.Game{
		ID:       g.ID,
		Title:    g.Title,
		Metadata: metadata,
	}, nil
}

func (g *Game) getMetadata() (map[string]any, error) {
	var metadata map[string]any
	if g.Metadata.Valid {
		err := json.Unmarshal(g.Metadata.V, &metadata)
		if err != nil {
			return nil, err
		}
	}
	return metadata, nil
}
