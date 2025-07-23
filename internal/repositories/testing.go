package repository

import (
	"fmt"
	"itk/internal"
	"strings"
	"testing"
)

func TestStore(t *testing.T, dataBaseUrl string) (*WalletRepository, func(...string)) {
	t.Helper()
	db := internal.InitDB(dataBaseUrl)

	repo := NewWalletRepository(db)

	return repo, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		db.Close()
	}
}
