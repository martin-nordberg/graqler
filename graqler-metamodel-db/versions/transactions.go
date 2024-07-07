package versions

import "sync/atomic"

// Transaction states.
const (
	InProgress uint64 = iota
	Committed
	RolledBack
)

// Txn represents one transaction in the linear history of transactions.
type Txn struct {
	txnNumber uint64
	state     atomic.Uint64
	prior     atomic.Pointer[Txn]
	// TODO: list of values affected
}

// lastTxn is the most recent transaction in the history of transactions.
var lastTxn atomic.Pointer[Txn]

// InitializeTxnHistory initializes the global transaction history.
func InitializeTxnHistory() {

	// The first transaction in the history is committed but with nothing affected.
	var emptyFirstTxn = &Txn{
		txnNumber: 0,
		state:     atomic.Uint64{},
		prior:     atomic.Pointer[Txn]{},
	}
	emptyFirstTxn.state.Store(Committed)

	lastTxn.Store(emptyFirstTxn)

}

// GetLastCommittedTxn finds the newest transaction in the history that has been committed.
func GetLastCommittedTxn() *Txn {

	var result *Txn

	for result = lastTxn.Load(); result.state.Load() != Committed; result = result.prior.Load() {
	}

	return result

}

// StartTxn creates and starts a new transaction.
func StartTxn() *Txn {

	// Construct a new transaction.
	var result = &Txn{
		txnNumber: 0,
		state:     atomic.Uint64{},
		prior:     atomic.Pointer[Txn]{},
	}
	result.state.Store(InProgress)

	for swapped := false; !swapped; {
		// Increment from the prior transaction.
		var prior = lastTxn.Load()
		result.prior.Store(prior)
		result.txnNumber = prior.txnNumber + 1

		// Attempt to link the new transaction into the global transaction history.
		swapped = lastTxn.CompareAndSwap(prior, result)
	}

	return result

}

// Commit attempts to commit a transaction.
func Commit(self *Txn) error {

	// TODO: commit each version
	// TODO: in a dedicated goroutine
	// TODO: handle commit failure
	self.state.Store(Committed)

	return nil

}
