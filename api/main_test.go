package api

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "github.com/tkircsi/simple-bank/db/sqlc"
	"github.com/tkircsi/simple-bank/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("cannot read configuration:", err)
	}
	config.AccessTokenDuration = time.Minute

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
