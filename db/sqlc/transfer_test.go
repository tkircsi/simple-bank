package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tkircsi/simple-bank/util"
)

func createRandomTransfer(t *testing.T, fromAccount Account, toAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func deleteRandomTransfer(t *testing.T, transfer Transfer) {
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
}

func TestCreateTransfer(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	_ = createRandomTransfer(t, fromAccount, toAccount)

}

func TestGetTransfer(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, fromAccount, toAccount)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestUpdateTransfer(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, fromAccount, toAccount)

	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestDeleteTransfer(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	transfer := createRandomTransfer(t, fromAccount, toAccount)

	deleteRandomTransfer(t, transfer)
}

func TestListTransfers(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 2,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
