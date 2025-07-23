package repository

import (
	"context"
	"database/sql"
	"fmt"
	"itk/internal/models"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) CreateWallet(w *models.Wallet) (*models.Wallet, error) {
	if err := r.db.QueryRow(
		"INSERT INTO wallets (amount) VALUES ($1) RETURNING id",
		w.Amount).Scan(&w.ID); err != nil {
		return nil, err
	}
	return w, nil
}

// GetWallet возвращает информацию о кошельке по ID
func (r *WalletRepository) GetWallet(walletId string) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	if err := r.db.QueryRow("SELECT * FROM wallets WHERE id = $1", walletId).Scan(&wallet.ID, &wallet.Amount); err != nil {
		return nil, err
	}
	return wallet, nil
}

// Deposit пополняет кошелек на сумму amount
func (r *WalletRepository) Deposit(walletId string, amount int) error {

	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var currentAmount int
	row := tx.QueryRow("SELECT amount FROM wallets WHERE id = $1 FOR UPDATE", walletId)
	if err = row.Scan(&currentAmount); err != nil {
		tx.Rollback()
		return fmt.Errorf("select for update: %w", err)
	}

	newAmount := currentAmount + amount
	_, err = tx.Exec("UPDATE wallets SET amount = $1 WHERE id = $2", newAmount, walletId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update amount: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}

// Withdraw снимает сумму amount с кошелька
func (r *WalletRepository) Withdraw(walletId string, amount int) error {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var currentAmount int
	row := tx.QueryRow("SELECT amount FROM wallets WHERE id = $1 FOR UPDATE", walletId)
	if err = row.Scan(&currentAmount); err != nil {
		tx.Rollback()
		return fmt.Errorf("select for update: %w", err)
	}

	if currentAmount < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient funds")
	}

	newAmount := currentAmount - amount
	_, err = tx.Exec("UPDATE wallets SET amount = $1 WHERE id = $2", newAmount, walletId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update amount: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}
