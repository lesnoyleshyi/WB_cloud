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
	errCh := make(chan error)

	fn := func(entities.Account) {
		d.mu.Lock()
		defer d.mu.Unlock()

		if _, ok := d.storage[account.Id]; ok {
			errCh <- domainErrors.ErrAccountAlreadyExists
		}

		d.storage[account.Id] = &account

		errCh <- nil
	}

	go fn(account)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		return err
	}
}

func (d MemDb) GetBalance(ctx context.Context, account entities.Account) (int, error) {
	type res struct {
		balance int
		err     error
	}
	ch := make(chan res)

	fn := func() {
		var ok bool
		var acc *entities.Account

		d.mu.RLock()
		defer d.mu.RUnlock()

		if acc, ok = d.storage[account.Id]; !ok {
			ch <- res{balance: 0, err: domainErrors.ErrNoSuchAccount}
			return
		}

		ch <- res{balance: acc.Balance, err: nil}
	}

	go fn()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case ret := <-ch:
		return ret.balance, ret.err
	}
}

func (d MemDb) Transfer(ctx context.Context, from entities.Account, to entities.Account, amount int) error {
	errCh := make(chan error)

	fn := func() {
		d.mu.Lock()
		defer d.mu.Unlock()

		if _, ok := d.storage[from.Id]; !ok {
			errCh <- domainErrors.ErrNoSuchAccount
		}
		if _, ok := d.storage[to.Id]; !ok {
			errCh <- domainErrors.ErrNoSuchAccount
		}

		if from.Balance < amount {
			errCh <- domainErrors.ErrNoEnoughMoney
		}

		d.storage[from.Id].Balance -= amount
		d.storage[to.Id].Balance += amount

		errCh <- nil
	}

	go fn()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		return err
	}
}
