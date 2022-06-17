package value

import (
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/value/exp/allocator"
)

type int32Value int32

func (v int32Value) toYDBType(a *allocator.Allocator) *Ydb.Type {
	typeId := a.TypePrimitive()
	typeId.TypeId = Ydb.Type_INT32

	t := a.Type()
	t.Type = typeId

	return t
}

func (v int32Value) toYDBValue(a *allocator.Allocator) *Ydb.Value {
	vv := a.Int32Value()
	vv.Int32Value = int32(v)

	vvv := a.Value()
	vvv.Value = vv

	return vvv
}

func Int32Value(v int32) int32Value {
	return int32Value(v)
}
