package utils

import (
	"WB_cloud/internal/domain/entities"
	"github.com/mailru/easyjson"
	"github.com/shopspring/decimal"
	"io"
)

type transferParams struct {
	From   string          `json:"from"`
	To     string          `json:"to"`
	Amount decimal.Decimal `json:"amount"`
}

func GetAccount(r io.ReadCloser) (*entities.Account, error) {
	var account entities.Account

	defer func() { _ = r.Close() }()
	if err := easyjson.UnmarshalFromReader(r, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func GetTransferParams(r io.ReadCloser) (from, to string, amount decimal.Decimal, err error) {
	var params transferParams

	defer func() { _ = r.Close() }()
	if err := easyjson.UnmarshalFromReader(r, &params); err != nil {
		return "", "", decimal.Decimal{}, err
	}

	return params.From, params.To, params.Amount, nil
}
