package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/tkircsi/simple-bank/db/sqlc"
	"github.com/tkircsi/simple-bank/token"
	"github.com/tkircsi/simple-bank/util"
)

type Server struct {
	config     *util.Config
	tokenMaker token.Maker
	store      db.Store
	router     *gin.Engine
}

func NewServer(config *util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot vreate token maker: %w", err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		store:      store,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", s.createUser)
	router.POST("/users/login", s.loginUser)

	router.POST("/accounts", s.createAccount)
	router.GET("/accounts/:id", s.getAccount)
	router.GET("accounts", s.listAccount)
	router.DELETE("accounts/:id", s.deleteAccount)
	router.PUT("/accounts/:id", s.updateAccount)

	router.POST("/transfers", s.createTransfer)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
