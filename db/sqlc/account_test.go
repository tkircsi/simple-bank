package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/tkircsi/simple-bank/util"
)

func createRandomAccount(t *testing.T) Account {
	u := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    u.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func deleteRandomAccount(t *testing.T, acc Account) {
	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)
}

func TestCreateAndDeleteAccount(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	acc := createRandomAccount(t)
	deleteRandomAccount(t, acc)

}

func TestGetAccount(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()
	// create account
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	account1 := createRandomAccount(t)

	deleteRandomAccount(t, account1)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListAccounts(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()
	// create 5 random account
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 2,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	// require.GreaterOrEqual(t, arg.Limit, int32(len(accounts)))
	require.Len(t, accounts, 5)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}

}

func TestListAccountsByOwner(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()
	// create 5 random account
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsByOwnerParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccountsByOwner(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
		require.Equal(t, lastAccount.Owner, acc.Owner)
	}
}
