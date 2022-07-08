package app

import (
	"WB_cloud/internal/adapters/memDB"
	ports "WB_cloud/internal/ports/output"
	"context"
)

type Storage interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	ports.AccountStorage
}

func NewStorage(storageType string) Storage {
	switch storageType {
	//case "postgres":
	//	return postgres.New()
	//case "mongo":
	//	return myMongo.New()
	case "in-mem":
		return memDB.New()
	default:
		return memDB.New()
	}
}
