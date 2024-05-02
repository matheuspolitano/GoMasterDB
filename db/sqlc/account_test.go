package db

import (
	"context"
	"testing"

	"github.com/matheuspolitano/GoMasterDB/utils"
	"github.com/stretchr/testify/assert"
)

func TestAddAccount(t *testing.T) {
	accountParams := AddAccountParams{
		Username: utils.RandUsername(),
		Ammount:  utils.RandAmount(),
		Currency: utils.RandCurrency(),
	}
	account, err := testQueries.AddAccount(context.Background(), accountParams)
	assert.NoError(t, err)
	assert.NotEmpty(t, account)
	assert.NotZero(t, account.ID)
	assert.Equal(t, account.Username, accountParams.Username)
	err = testQueries.RemoveAccount(context.Background(), account.ID)
	assert.NoError(t, err)
}
