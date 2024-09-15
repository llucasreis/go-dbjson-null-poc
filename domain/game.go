package domain

type Game struct {
	ID          int64
	Title       string
	Description string
	Metadata    map[string]any
}
