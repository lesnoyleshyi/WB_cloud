package ports

import (
	"WB_cloud/internal/domain/entities"
	"context"
	"github.com/shopspring/decimal"
)

type AccountStorage interface {
	Create(ctx context.Context, account entities.Account) error
	GetBalance(ctx context.Context, account entities.Account) (balance decimal.Decimal, err error)
	Transfer(ctx context.Context, from entities.Account, to entities.Account, amount decimal.Decimal) error
}
