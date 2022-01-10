package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tkircsi/simple-bank/util"
)

func TestPasetoMaker(t *testing.T) {
	key := util.GenerateRandomString(32)
	maker, err := NewPasetoMaker(key)
	require.NoError(t, err)

	user := util.RandomUser()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(user[0], duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, user[0], payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	key := util.GenerateRandomString(32)
	maker, err := NewPasetoMaker(key)
	require.NoError(t, err)

	user := util.RandomUser()

	token, err := maker.CreateToken(user[0], -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
