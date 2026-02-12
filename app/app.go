package app

import (
	"main/db"
	"main/handler"
	"main/repository"
	"main/router"
	"main/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	db.ConnectDatabase()

	walletRepo := repository.NewWalletRepository(db.DB)
	walletService := service.NewWalletService(walletRepo, db.DB)
	walletHandler := handler.NewWalletHandler(walletService)

	r := router.SetupRouter(walletHandler)

	return &App{
		Router: r,
	}
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}