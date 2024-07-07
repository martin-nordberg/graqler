package handles

import (
	"sync/atomic"
)

//-------------------------------------------------------------------------------------------------

type Int32Handle struct {
	handleId HandleId
	value    [2]int32
}

//-------------------------------------------------------------------------------------------------

func MakeInt32Handle(value int32) Int32Handle {
	return Int32Handle{
		value:    [2]int32{value, value},
		handleId: getNextInt32HandleId(),
	}
}

//-------------------------------------------------------------------------------------------------

func (h *Int32Handle) HandleId() HandleId {
	return h.handleId
}

//-------------------------------------------------------------------------------------------------

func (h *Int32Handle) Read(blueOrGreen BlueGreen) int32 {
	return h.value[blueOrGreen]
}

//-------------------------------------------------------------------------------------------------

func (h *Int32Handle) Write(blueOrGreen BlueGreen, value int32) {
	h.value[blueOrGreen] = value
}

//=================================================================================================

var nextInt32HandleId atomic.Uint64

//-------------------------------------------------------------------------------------------------

func init() {
	nextInt32HandleId.Store(1)
}

//-------------------------------------------------------------------------------------------------

func getNextInt32HandleId() HandleId {
	for {
		var h = nextInt32HandleId.Load()
		var result = h * two64DivPhi
		if nextInt32HandleId.CompareAndSwap(h, result) {
			return result
		}
	}
}

//-------------------------------------------------------------------------------------------------
