// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccount(ctx context.Context, accountID int64) error
	DeleteEntry(ctx context.Context, entryID int64) error
	DeleteTransfer(ctx context.Context, transferID int64) error
	GetAccount(ctx context.Context, accountID int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, accountID int64) (Account, error)
	GetEntry(ctx context.Context, entryID int64) (Entry, error)
	GetTransfer(ctx context.Context, transferID int64) (Transfer, error)
	ListAсcounts(ctx context.Context, arg ListAсcountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error)
	UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error)
}

var _ Querier = (*Queries)(nil)