package errors

import "errors"

var (
	ErrAccountAlreadyExists = errors.New("account with such id is already exists")
	ErrNoSuchAccount        = errors.New("no account with such id is in database")
	ErrNoEnoughMoney        = errors.New("no enough money")
)
