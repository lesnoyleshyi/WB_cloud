package ports

import (
	"WB_cloud/internal/domain/entities"
	"context"
)

type AccountStorage interface {
	Create(ctx context.Context, account entities.Account) error
	GetBalance(ctx context.Context, account entities.Account) (balance int, err error)
	Transfer(ctx context.Context, from entities.Account, to entities.Account, amount int) error
}
