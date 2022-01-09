package db

import (
	"context"
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tkircsi/simple-bank/util"
)

func createRandomUser(t *testing.T) User {
	ruser := util.RandomUser()

	arg := CreateUserParams{
		Username:       ruser[0],
		FullName:       ruser[1],
		Email:          ruser[2],
		HashedPassword: ruser[3],
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.CreatedAt)

	return user
}

func deleteRandomUser(t *testing.T, user User) {
	err := testQueries.DeleteUser(context.Background(), user.Username)
	require.NoError(t, err)
}

func TestCreateAndDeleteUser(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	user := createRandomUser(t)
	deleteRandomUser(t, user)
}

func TestGetUser(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		Username: user1.Username,
		FullName: "Dr. " + user1.FullName,
		Email:    "dr." + user1.Email,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, "dr."+user1.Email, user2.Email)
	require.Equal(t, "Dr. "+user1.FullName, user2.FullName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserPassword(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	user1 := createRandomUser(t)

	newpasswod := fmt.Sprintf("%x", sha256.Sum256([]byte("newpassword")))
	arg := UpdateUserPasswordParams{
		Username:          user1.Username,
		HashedPassword:    newpasswod,
		PasswordChangedAt: time.Now(),
	}

	user2, err := testQueries.UpdateUserPassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, newpasswod, user2.HashedPassword)

	require.WithinDuration(t, arg.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestListUsers(t *testing.T) {
	defer func() {
		err := testQueries.CleanUpDB(context.Background())
		require.NoError(t, err)
	}()

	var rusers []User
	for i := 0; i < 10; i++ {
		rusers = append(rusers, createRandomUser(t))
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 2,
	}
	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
