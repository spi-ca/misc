package types

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spi-ca/misc"
	"time"

	"github.com/pkg/errors"
)

type JSONTime time.Time

func (dttm JSONTime) MarshalJSON() ([]byte, error) {
	stream := misc.JSONCodec.BorrowStream(nil)
	defer misc.JSONCodec.ReturnStream(stream)
	stream.WriteInt64(time.Time(dttm).UTC().Unix())
	return append([]byte(nil), stream.Buffer()...), stream.Error
}

func (dttm *JSONTime) UnmarshalJSON(b []byte) error {
	iterator := misc.JSONCodec.BorrowIterator(b)
	defer misc.JSONCodec.ReturnIterator(iterator)
	valueType := iterator.WhatIsNext()
	switch valueType {
	case jsoniter.NumberValue:
		*dttm = JSONTime(time.Unix(iterator.ReadInt64(), 0))
		return nil
	default:
		return errors.Errorf("%d is wrong type", valueType)
	}
}
