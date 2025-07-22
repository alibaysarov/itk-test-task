package controllers

import (
	wallet_repository "itk/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletController struct {
	Repo *wallet_repository.WalletRepository
}

type WalletOperationRequest struct {
	WalletId      uuid.UUID `json:"walletId" binding:"required"`
	OperationType string    `json:"operationType" binding:"required,oneof=DEPOSIT WITHDRAW"`
	Amount        int       `json:"amount" binding:"required,gt=0"`
}
type DepositRequest struct {
	Amount int `json:"amount" binding:"required,gt=0"`
}
type WithdrawRequest struct {
	Amount int `json:"amount" binding:"required,gt=0"`
}

// GetWallet godoc
// @Summary Получить кошелек
// @Description Получить информацию о кошельке по ID
// @Tags wallets
// @Param walletId path string true "ID кошелька"
// @Success 200 {object} models.Wallet
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /wallets/{walletId} [get]
func (wc *WalletController) GetWallet(c *gin.Context) {
	walletIdStr := c.Param("walletId")
	walletId, err := uuid.Parse(walletIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid walletId"})
		return
	}

	wallet, err := wc.Repo.GetWallet(walletId.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

// Deposit godoc
// @Summary Пополнить кошелек
// @Description Пополнить кошелек на сумму
// @Tags wallets
// @Param walletId path string true "ID кошелька"
// @Param input body DepositRequest true "Сумма для пополнения"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /wallets/{walletId}/deposit [post]
func (wc *WalletController) Deposit(c *gin.Context) {
	walletIdStr := c.Param("walletId")
	walletId, err := uuid.Parse(walletIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid walletId"})
		return
	}
	var req DepositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wc.Repo.Deposit(walletId.String(), req.Amount)
	c.JSON(http.StatusOK, gin.H{"status": "deposit successful"})
}

// Withdraw godoc
// @Summary Снять средства
// @Description Снять средства с кошелька
// @Tags wallets
// @Param walletId path string true "ID кошелька"
// @Param input body WithdrawRequest true "Сумма для снятия"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /wallets/{walletId}/withdraw [post]
func (wc *WalletController) Withdraw(c *gin.Context) {
	walletIdStr := c.Param("walletId")
	walletId, err := uuid.Parse(walletIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid walletId"})
		return
	}
	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wc.Repo.Withdraw(walletId.String(), req.Amount)
	c.JSON(http.StatusOK, gin.H{"status": "withdraw successful"})
}
