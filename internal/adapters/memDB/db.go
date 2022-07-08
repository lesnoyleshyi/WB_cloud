package memDB

import (
	"WB_cloud/internal/domain/entities"
	domainErrors "WB_cloud/internal/domain/errors"
	"context"
	"sync"
)

type MemDb struct {
	storage map[string]*entities.Account
	mu      *sync.RWMutex
}

func New() MemDb {
	storage := make(map[string]*entities.Account)
	mu := sync.RWMutex{}

	return MemDb{
		storage: storage,
		mu:      &mu,
	}
}

func (d MemDb) Connect(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (d MemDb) Close(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (d MemDb) Create(ctx context.Context, account entities.Account) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.storage[account.Id]; ok {
		return domainErrors.ErrAccountAlreadyExists
	}

	d.storage[account.Id] = &account

	return nil
}

func (d MemDb) GetBalance(ctx context.Context, account entities.Account) (int, error) {
	var ok bool
	var acc *entities.Account

	d.mu.RLock()
	defer d.mu.RUnlock()

	if acc, ok = d.storage[account.Id]; !ok {
		return 0, domainErrors.ErrNoSuchAccount
	}

	return acc.Balance, nil
}

func (d MemDb) Transfer(ctx context.Context, from entities.Account, to entities.Account, amount int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.storage[from.Id]; !ok {
		return domainErrors.ErrNoSuchAccount
	}
	if _, ok := d.storage[to.Id]; !ok {
		return domainErrors.ErrNoSuchAccount
	}
	if from.Balance < amount {
		return domainErrors.ErrNoEnoughMoney

	}

	d.storage[from.Id].Balance -= amount
	d.storage[to.Id].Balance += amount

	return nil
}
