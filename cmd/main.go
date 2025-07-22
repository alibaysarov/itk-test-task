package main

import (
	"itk/internal"
	"itk/internal/controllers"
	wallet_repository "itk/internal/repositories"

	_ "itk/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ITK Wallet API
// @version 1.0
// @description API for wallet operations
// @BasePath /api/v1

func main() {
	db := internal.InitDB()
	repo := wallet_repository.NewWalletRepository(db)
	walletController := &controllers.WalletController{Repo: repo}

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/wallets/:walletId", walletController.GetWallet)
		v1.POST("/wallets/:walletId/deposit", walletController.Deposit)
		v1.POST("/wallets/:walletId/withdraw", walletController.Withdraw)
	}
	r.Run(":8080")
}
