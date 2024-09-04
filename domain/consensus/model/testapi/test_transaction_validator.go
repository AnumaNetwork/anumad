package testapi

import (
	"github.com/AnumaNetwork/anumad/domain/consensus/model"
	"github.com/AnumaNetwork/anumad/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
