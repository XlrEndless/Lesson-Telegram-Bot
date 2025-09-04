package persist

import (
	"TgBot/cmd/core/service/port"
	"context"
	"gorm.io/gorm"
)

type OrmTransactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) port.ITransactionManager {
	return &OrmTransactionManager{db}
}

func (manager *OrmTransactionManager) DoInTransaction(ctx context.Context, txFunc func(ctx context.Context) error) error {
	value := ctx.Value("db")
	var err error
	if value == nil {
		err = manager.db.Transaction(func(tx *gorm.DB) error {
			ctx := context.WithValue(context.Background(), "db", tx)
			err := txFunc(ctx)
			return err
		})
	} else {
		err = txFunc(ctx)
	}
	return err
}
