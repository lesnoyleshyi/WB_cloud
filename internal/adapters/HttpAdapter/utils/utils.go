package utils

import (
	"WB_cloud/internal/domain/entities"
	"github.com/mailru/easyjson"
	"io"
)

type transferParams struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func GetAccount(r io.ReadCloser) (*entities.Account, error) {
	var account entities.Account

	defer func() { _ = r.Close() }()
	if err := easyjson.UnmarshalFromReader(r, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func GetTransferParams(r io.ReadCloser) (from, to string, amount float64, err error) {
	var params transferParams

	defer func() { _ = r.Close() }()
	if err := easyjson.UnmarshalFromReader(r, &params); err != nil {
		return "", "", 0, err
	}

	return params.From, params.To, params.Amount, nil
}
