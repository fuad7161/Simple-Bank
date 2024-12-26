// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntrie(ctx context.Context, arg CreateEntrieParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntrie(ctx context.Context, id int64) error
	DeleteTransfer(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetEntrie(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]Account, error)
	ListEntrie(ctx context.Context, arg ListEntrieParams) ([]Entry, error)
	ListTransfer(ctx context.Context, arg ListTransferParams) ([]Transfer, error)
	UpdateAccountBios(ctx context.Context, arg UpdateAccountBiosParams) (Account, error)
	UpdateEntrieBios(ctx context.Context, arg UpdateEntrieBiosParams) (Entry, error)
	UpdateTransferBios(ctx context.Context, arg UpdateTransferBiosParams) (Transfer, error)
}

var _ Querier = (*Queries)(nil)
