package api

import (
	"testing"
	"time"

	db "github.com/santhoshvempali/simplebank/db/sqlc"
	"github.com/santhoshvempali/simplebank/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		SECRET:          util.RandomString(32),
		ACCESS_DURATION: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}
