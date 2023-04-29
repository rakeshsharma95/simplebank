package gapi

import (
	"testing"
	"time"

	db "github.com/rdevelop/simplebank/db/sqlc"
	"github.com/rdevelop/simplebank/util"
	"github.com/rdevelop/simplebank/worker"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuretion: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
