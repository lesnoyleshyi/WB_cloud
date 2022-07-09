package usecases

import (
	"WB_cloud/internal/domain/entities"
	"WB_cloud/internal/ports/output"
	"context"
	"github.com/shopspring/decimal"
)

type AccountService struct {
	storage ports.AccountStorage
}

func NewAccountService(s ports.AccountStorage) AccountService {
	return AccountService{storage: s}
}

func (a AccountService) Create(ctx context.Context, account entities.Account) error {
	return a.storage.Create(ctx, account)
}

func (a AccountService) GetBalance(ctx context.Context, account entities.Account) (balance decimal.Decimal, err error) {
	return a.storage.GetBalance(ctx, account)
}

func (a AccountService) Transfer(ctx context.Context, from entities.Account, to entities.Account, amount decimal.Decimal) error {
	return a.storage.Transfer(ctx, from, to, amount)
}
