package repository

import (
	"itk/internal/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dataBaseUrl string
)

func TestMain(t *testing.M) {
	dataBaseUrl = os.Getenv("DATABASE_TEST_URL")
	if dataBaseUrl == "" {
		dataBaseUrl = "postgres://user:password@localhost:54321/test_db?sslmode=disable"
	}
	os.Exit(t.Run())
}

func TestWalletRepository_Create(t *testing.T) {
	repo, teardown := TestStore(t, dataBaseUrl)
	defer teardown("wallets")
	w, err := repo.CreateWallet(&models.Wallet{
		Amount: 150,
	})
	assert.NoError(t, err)
	assert.NotNil(t, w)
}

func TestWalletRepository_Get(t *testing.T) {
	repo, teardown := TestStore(t, dataBaseUrl)
	defer teardown("wallets")
	w, err := repo.CreateWallet(&models.Wallet{Amount: 150})
	if err != nil {
		t.Fatal(err)
	}
	fetchedWallet, err := repo.GetWallet(w.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedWallet)
}
