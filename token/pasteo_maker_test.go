package token

import (
	"testing"
	"time"

	"github.com/santhoshvempali/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasteo(t *testing.T) {
	maker, err := NewPasteoMaker(util.RandomString(32))
	require.NoError(t, err)
	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, username, payload.Username)
	require.WithinDuration(t, expiresAt, payload.Expires_at, time.Second)

}
