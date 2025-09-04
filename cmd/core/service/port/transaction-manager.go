package port

import "context"

type ITransactionManager interface {
	DoInTransaction(ctx context.Context, txFunc func(ctx context.Context) error) error
}
