package postgres

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"github.com/jackc/pgx/v5"
)

func (p *postgres) CreateCard(ctx context.Context, tx pgx.Tx, card *domain.Card) (int, error) {
	q1 := `insert into cards (title, description, body, price, thumbnail) values ($1, $2, $3, $4, $5) returning id`
	row := tx.QueryRow(ctx, q1, card.Title, card.Description, card.Body, card.Price, card.Thumbnail)

	var cardId int
	if err := row.Scan(&cardId); err != nil {
		return 0, err
	}

	return cardId, nil
}
