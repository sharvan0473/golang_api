package dbs

import (
	"testing"
)

// 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
func TestTransferTx(t *testing.T) {
	//store := NewStore(testDb)
	//account := createRandomAccount(t)
	//account2 := createRandomAccount(t)
	//
	//n := 5
	//amount := int64(10)
	//
	//errs := make(chan error)
	//results := make(chan TransferTxResult)
	//for i := 0; i < n; i++ {
	//	go func() {
	//		result, err := store.TransferTx(context.Background(), TansferTxParams{
	//			FromAccountId: account.ID,
	//			ToAccountId:   account2.ID,
	//			Amount:        amount,
	//		})
	//		errs <- err
	//		results <- result
	//	}()
	//}
	//for i := 0; i < n; i++ {
	//	err := <-errs
	//	require.NoError(t, err)
	//
	//	result := <-results
	//	require.NotEmpty(t, result)
	//	transfer := result.Transfer
	//	require.NotEmpty(t, transfer)
	//	require.Equal(t, account.ID, transfer.FromAccountID)
	//	require.Equal(t, account2.ID, transfer.ToAccountID)
	//
	//	_, err = store.GetTransfer(context.Background(), transfer.ID)
	//	require.NoError(t, err)
	//
	//	//check entry
	//	fromEntry := result.FromEntry
	//
	//	require.NotEmpty(t, fromEntry)
	//	require.Equal(t, account.ID, fromEntry.AccountID)
	//	require.Equal(t, -amount, fromEntry.Amount)
	//
	//	_, err = store.GetEntry(context.Background(), fromEntry.ID)
	//	require.NoError(t, err)
	//}

}
