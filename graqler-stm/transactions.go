package stm

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
