package account

import (
	"brimoveapp/dto"
	"context"
)

type AccountService struct {
	AccountRepo *AccountRepository
}

func NewAccountService(accRepo *AccountRepository) *AccountService {
	return &AccountService{
		AccountRepo: accRepo,
	}
}

// method insert account
func (a *AccountService) Insert(ctx context.Context, request *dto.InsertAccountRequest) {

}

// method get account by usename
func (a *AccountService) GetByEmail(ctx context.Context, username string) (*dto.InsertAccountResponse, error) {
	// call procedure get by username in repository
	account, err := a.AccountRepo.GetByEmail(ctx, username)
	if err != nil {
		return nil, err
	}

	// create response object
	response := dto.InsertAccountResponse{
		account.Username, account.Password,
	}

	return &response, nil
}
