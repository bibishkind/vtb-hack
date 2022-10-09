package postgres

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"github.com/jackc/pgx/v5"
)

func (p *postgres) CreateUser(ctx context.Context, tx pgx.Tx, user *domain.User) (int, error) {
	q1 := `insert into users (username, password, first_name, middle_name, last_name, email, role, public_key, private_key)
           values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`
	row := tx.QueryRow(ctx, q1, user.Username, user.Password, user.FirstName, user.MiddleName, user.LastName, user.Email, user.Role, user.PublicKey, user.PrivateKey)

	var userId int
	if err := row.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}

func (p *postgres) GetUserByUsername(ctx context.Context, tx pgx.Tx, username string) (*domain.User, error) {
	q1 := `select * from users where username=$1`
	row := tx.QueryRow(ctx, q1, username)

	user := new(domain.User)

	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.Role, &user.PublicKey, &user.PrivateKey); err != nil {
		return nil, err
	}

	return user, nil
}

func (p *postgres) GetUserById(ctx context.Context, tx pgx.Tx, userId int) (*domain.User, error) {
	q1 := `select * from users where id=$1`
	row := tx.QueryRow(ctx, q1, userId)

	user := new(domain.User)

	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.Role, &user.PublicKey, &user.PrivateKey); err != nil {
		return nil, err
	}

	return user, nil
}
