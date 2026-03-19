package http

import (
	"net/http"

	appAuth "github.com/alves/age-of-genesis/internal/application/auth"
	appPayment "github.com/alves/age-of-genesis/internal/application/payment"
	"github.com/alves/age-of-genesis/internal/config"
	"github.com/alves/age-of-genesis/internal/infrastructure/pagseguro"
	"github.com/alves/age-of-genesis/internal/infrastructure/persistence/mysql"
	"github.com/alves/age-of-genesis/internal/infrastructure/token"
	"github.com/alves/age-of-genesis/internal/interfaces/http/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    config.Config
	router *gin.Engine
}

func NewServer(cfg config.Config) (*Server, error) {
	db, err := mysql.NewConnection(cfg.MySQLDSN)
	if err != nil {
		return nil, err
	}

	userRepo := mysql.NewUserRepository(db)
	txRepo := mysql.NewTransactionRepository(db)

	jwtSvc := token.NewJWTService(cfg.JWTSecret, cfg.JWTExpiresInMinutes)
	pagSeguroClient := pagseguro.NewClient(cfg.PagSeguroBaseURL, cfg.PagSeguroToken)

	loginSvc := appAuth.NewLoginService(userRepo, jwtSvc)
	chargeSvc := appPayment.NewChargeService(pagSeguroClient, txRepo)

	authHandler := handler.NewAuthHandler(loginSvc)
	paymentHandler := handler.NewPaymentHandler(chargeSvc)

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api")
	api.POST("/login", authHandler.Login)
	api.POST("/payments/credit-card", paymentHandler.ChargeCreditCard)

	return &Server{
		cfg:    cfg,
		router: router,
	}, nil
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.cfg.AppPort)
}
