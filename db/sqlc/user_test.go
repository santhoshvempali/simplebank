package db

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/santhoshvempali/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	if err != nil {
		fmt.Print("error in creating hashed password", err)
	}
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.True(t, user.PasswordChangedAt.IsZero())

	require.NotZero(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Username, user2.Username)

}

// func TestUpdateAccount(t *testing.T) {
// 	acc1 := createRandomAccount(t)
// 	arg := UpdateAccountParams{
// 		ID:      acc1.ID,
// 		Balance: util.RandomMoney(),
// 	}
// 	acc2, err := testQueries.UpdateAccount(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, acc2)
// 	require.Equal(t, acc1.ID, acc2.ID)
// }
