package stm

import (
	"graqler/stm/stm/handles"
	"sync/atomic"
	"time"
)

type ReadTransactionContext struct {
	colorToRead atomic.Int64

	// toBeReadInt32 is a blue/green pair of HandleMaps of int32 values to be read instead of global object handles.
	toBeReadInt32 [2]handles.Int32HandleMap

	toBeReadString [2]handles.StringHandleMap

	toBeReadTime [2]handles.TimeHandleMap
}

func MakeReadTransactionContext(colorToRead int64) *ReadTransactionContext {

	result := &ReadTransactionContext{
		colorToRead:    atomic.Int64{},
		toBeReadInt32:  [2]HandleMap[int32]{},
		toBeReadString: [2]HandleMap[string]{},
		toBeReadTime:   [2]HandleMap[time.Time]{},
	}

	result.colorToRead.Store(colorToRead)
	result.toBeReadInt32[colorToRead] = MakeHandleMap[int32](16)
	result.toBeReadString[colorToRead] = MakeHandleMap[string](16)
	result.toBeReadTime[colorToRead] = MakeHandleMap[time.Time](16)

	return result
}

func (ctx *ReadTransactionContext) ReadInt32(h handles.Int32Handle) int32 {
	c := ctx.colorToRead.Load()
	result, found := ctx.toBeReadInt32[c].Get(h.HandleId())
	if !found {
		result = h.Read(c)
	}
	return result
}

func (ctx *ReadTransactionContext) ReadString(h handles.StringHandle) string {
	c := ctx.colorToRead.Load()
	result, found := ctx.toBeReadString[c].Get(h.HandleId())
	if !found {
		result = h.Read(c)
	}
	return result
}

func (ctx *ReadTransactionContext) ReadTime(h handles.TimeHandle) time.Time {
	c := ctx.colorToRead.Load()
	result, found := ctx.toBeReadTime[c].Get(h.HandleId())
	if !found {
		result = h.Read(c)
	}
	return result
}
