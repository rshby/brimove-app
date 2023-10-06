package account

import (
	"context"
	"database/sql"
)

type AccountRepository struct {
	DB *sql.DB
}

// function provider
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db,
	}
}

// method create account
func (a *AccountRepository) Insert(ctx context.Context, entity *Account) (*Account, error) {
	query := "INSERT INTO accounts(username, password) VALUES($1, $2) RETURNING id"

	res := a.DB.QueryRowContext(ctx, query, entity.Username, entity.Password)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var id int
	if err := res.Scan(&id); err != nil {
		return nil, err
	}

	entity.Id = int(id)
	return entity, nil
}

// method get by username
func (a *AccountRepository) GetByEmail(ctx context.Context, email string) (*Account, error) {
	query := "SELECT id, username, password FROM accounts WHERE username=$1"

	row := a.DB.QueryRowContext(ctx, query, email)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var account Account
	if err := row.Scan(&account.Id, &account.Username, &account.Password); err != nil {
		return nil, err
	}

	return &account, nil
}
