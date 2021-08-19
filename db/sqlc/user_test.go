package db

import (
	"context"
	"github.com/JekaTka/auth-service/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Nickname: util.RandomNickname(),
		Email:    util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Nickname, user.Nickname)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T)  {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	expectedUser := createRandomUser(t)
	user, err := testQueries.GetUser(context.Background(), expectedUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, expectedUser.ID, user.ID)
	require.Equal(t, expectedUser.Email, user.Email)
	require.Equal(t, expectedUser.Nickname, user.Nickname)
	require.WithinDuration(t, expectedUser.CreatedAt.Time, user.CreatedAt.Time, time.Second)
}