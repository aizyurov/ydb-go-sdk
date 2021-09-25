// Code generated by ydbgen; DO NOT EDIT.

package tests

import (
	"strconv"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/resultset"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

var (
	_ = strconv.Itoa
	_ = time.Now
	_ = table.NewQueryParameters
	_ = resultset.Result.Scan
	_ = types.StringValue
)

func (c *ConvAssert) Scan(res resultset.Result) (err error) {
	_ = res.ScanWithDefaults(&c.Int8Int16, &c.Int32Int64, &c.Int16Int8, &c.Uint64Int8, &c.Uint32Uint, &c.Int32Int, &c.Int32ToByte)
	return res.Err()
}