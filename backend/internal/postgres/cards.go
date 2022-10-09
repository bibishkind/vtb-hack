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

func (p *postgres) GetAllCards(ctx context.Context, tx pgx.Tx) ([]*domain.Card, error) {
	q1 := `select * from cards`
	rows, err := tx.Query(ctx, q1)
	if err != nil {
		return nil, err
	}

	var cards []*domain.Card

	defer rows.Close()

	for rows.Next() {
		var card domain.Card
		if err = rows.Scan(&card.Id, &card.Title, &card.Description, &card.Body, &card.Price, &card.Thumbnail); err != nil {
			return nil, err
		}
		cards = append(cards, &card)
	}

	return cards, nil
}

func (p *postgres) DeleteCard(ctx context.Context, tx pgx.Tx, cardId int) error {
	q1 := `delete from cards where id=$1`
	_, err := tx.Exec(ctx, q1, cardId)
	return err
}
