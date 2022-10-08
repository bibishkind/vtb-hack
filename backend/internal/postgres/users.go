package postgres

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) CreateUser(ctx context.Context, tx pgx.Tx, user *domain.User) error {
	q1 := `insert into users (username, password, first_name, middle_name, last_name, email, role, wallet)
           values ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := tx.Exec(ctx, q1, user.Username, user.Password, user.FirstName, user.MiddleName, user.LastName, user.Email, user.Role, user.Wallet)
	return err
}
func (p *Postgres) GetUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*domain.User, error) {
	q1 := `select * from users where username=$1`
	row := tx.QueryRow(ctx, q1, username)

	user := new(domain.User)

	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.Role, &user.Wallet); err != nil {
		return nil, err
	}

	return user, nil
}
