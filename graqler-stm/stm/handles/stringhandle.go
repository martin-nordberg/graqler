package handles

import (
	"sync/atomic"
)

//-------------------------------------------------------------------------------------------------

type StringHandle struct {
	handleId HandleId
	value    [2]string
}

//-------------------------------------------------------------------------------------------------

func MakeStringHandle(value string) StringHandle {
	return StringHandle{
		value:    [2]string{value, value},
		handleId: getNextStringHandleId(),
	}
}

//-------------------------------------------------------------------------------------------------

func (h *StringHandle) HandleId() HandleId {
	return h.handleId
}

//-------------------------------------------------------------------------------------------------

func (h *StringHandle) Read(blueOrGreen BlueGreen) string {
	return h.value[blueOrGreen]
}

//-------------------------------------------------------------------------------------------------

func (h *StringHandle) Write(blueOrGreen BlueGreen, value string) {
	h.value[blueOrGreen] = value
}

//=================================================================================================

var nextStringHandleId atomic.Uint64

//-------------------------------------------------------------------------------------------------

func init() {
	nextStringHandleId.Store(1)
}

//-------------------------------------------------------------------------------------------------

func getNextStringHandleId() HandleId {
	for {
		h := nextStringHandleId.Load()
		result := h * two64DivPhi
		if nextStringHandleId.CompareAndSwap(h, result) {
			return result
		}
	}
}

//-------------------------------------------------------------------------------------------------
