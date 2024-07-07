package handles

import (
	"sync/atomic"
)

//-------------------------------------------------------------------------------------------------

type BoolHandle struct {
	handleId HandleId
	value    [2]bool
}

//-------------------------------------------------------------------------------------------------

func MakeBoolHandle(value bool) BoolHandle {
	return BoolHandle{
		value:    [2]bool{value, value},
		handleId: getNextBoolHandleId(),
	}
}

//-------------------------------------------------------------------------------------------------

func (h *BoolHandle) HandleId() HandleId {
	return h.handleId
}

//-------------------------------------------------------------------------------------------------

func (h *BoolHandle) Read(blueOrGreen BlueGreen) bool {
	return h.value[blueOrGreen]
}

//-------------------------------------------------------------------------------------------------

func (h *BoolHandle) Write(blueOrGreen BlueGreen, value bool) {
	h.value[blueOrGreen] = value
}

//=================================================================================================

var nextBoolHandleId atomic.Uint64

//-------------------------------------------------------------------------------------------------

func init() {
	nextBoolHandleId.Store(1)
}

//-------------------------------------------------------------------------------------------------

func getNextBoolHandleId() HandleId {
	for {
		var h = nextBoolHandleId.Load()
		var result = h * two64DivPhi
		if nextBoolHandleId.CompareAndSwap(h, result) {
			return result
		}
	}
}

//-------------------------------------------------------------------------------------------------
