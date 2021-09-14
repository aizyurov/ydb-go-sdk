// Code generated by ydbgen; DO NOT EDIT.

package tests

import (
	"strconv"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

var (
	_ = strconv.Itoa
	_ = ydb.StringValue
	_ = table.NewQueryParameters
)

func ydbConvI16ToI8(x int16) int8 {
	const (
		bits = 8
		mask = (1 << (bits - 1)) - 1
	)
	var abs uint64
	{
		v := int64(x)
		m := v >> 63
		abs = uint64(v ^ m - m)
	}
	if abs&mask != abs {
		panic(
			"ydbgen: convassert: " + strconv.FormatInt(int64(x), 10) +
				" (type int16) overflows int8",
		)
	}
	return int8(x)
}

func ydbConvI32ToB(x int32) byte {
	if x < 0 {
		panic("ydbgen: convassert: conversion of negative int32 to byte")
	}
	const (
		bits = 8
		mask = (1 << (bits)) - 1
	)
	var abs uint64
	{
		v := int64(x)
		m := v >> 63
		abs = uint64(v ^ m - m)
	}
	if abs&mask != abs {
		panic(
			"ydbgen: convassert: " + strconv.FormatInt(int64(x), 10) +
				" (type int32) overflows byte",
		)
	}
	return byte(x)
}

func ydbConvU64ToI8(x uint64) int8 {
	const (
		bits = 8
		mask = (1 << (bits - 1)) - 1
	)
	abs := uint64(x)
	if abs&mask != abs {
		panic(
			"ydbgen: convassert: " + strconv.FormatUint(uint64(x), 10) +
				" (type uint64) overflows int8",
		)
	}
	return int8(x)
}
