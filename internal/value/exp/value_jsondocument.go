package value

import (
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/value/exp/allocator"
)

type jsondocumentValue string

func (v jsondocumentValue) toYDBType(a *allocator.Allocator) *Ydb.Type {
	typeId := a.TypePrimitive()
	typeId.TypeId = Ydb.Type_JSON_DOCUMENT

	t := a.Type()
	t.Type = typeId

	return t
}

func (v jsondocumentValue) toYDBValue(a *allocator.Allocator) *Ydb.Value {
	vv := a.TextValue()
	vv.TextValue = string(v)

	vvv := a.Value()
	vvv.Value = vv

	return vvv
}

func JSONDocumentValue(v string) jsondocumentValue {
	return jsondocumentValue(v)
}
