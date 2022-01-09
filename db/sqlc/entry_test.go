package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tkircsi/simple-bank/util"
)

func createRandomEntry(t *testing.T, acc Account) Entry {

	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func deleteRandomEntry(t *testing.T, entry Entry) {
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
}

func TestCreateEntry(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	acc := createRandomAccount(t)
	_ = createRandomEntry(t, acc)
}

func TestGetEntry(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	acc := createRandomAccount(t)
	entry1 := createRandomEntry(t, acc)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestUpdateEntry(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	acc := createRandomAccount(t)
	entry1 := createRandomEntry(t, acc)

	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	acc := createRandomAccount(t)
	entry1 := createRandomEntry(t, acc)
	deleteRandomEntry(t, entry1)
}

func TestListEntries(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	acc := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 2,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
