package handles

import (
	"sync/atomic"
	"time"
)

//-------------------------------------------------------------------------------------------------

type TimeHandle struct {
	handleId HandleId
	value    [2]time.Time
}

//-------------------------------------------------------------------------------------------------

func MakeTimeHandle(value time.Time) TimeHandle {
	return TimeHandle{
		value:    [2]time.Time{value, value},
		handleId: getNextTimeHandleId(),
	}
}

//-------------------------------------------------------------------------------------------------

func (h *TimeHandle) HandleId() HandleId {
	return h.handleId
}

//-------------------------------------------------------------------------------------------------

func (h *TimeHandle) Read(blueOrGreen BlueGreen) time.Time {
	return h.value[blueOrGreen]
}

//-------------------------------------------------------------------------------------------------

func (h *TimeHandle) Write(blueOrGreen BlueGreen, value time.Time) {
	h.value[blueOrGreen] = value
}

//=================================================================================================

var nextTimeHandleId atomic.Uint64

//-------------------------------------------------------------------------------------------------

func init() {
	nextTimeHandleId.Store(1)
}

//-------------------------------------------------------------------------------------------------

func getNextTimeHandleId() HandleId {
	for {
		var h = nextTimeHandleId.Load()
		var result = h * two64DivPhi
		if nextTimeHandleId.CompareAndSwap(h, result) {
			return result
		}
	}
}

//-------------------------------------------------------------------------------------------------
