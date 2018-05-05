package ignite

import (
	"time"
)

const (
	// Supported standard types and their type codes are as follows:
	typeByte        = 1
	typeShort       = 2
	typeInt         = 3
	typeLong        = 4
	typeFloat       = 5
	typeDouble      = 6
	typeChar        = 7
	typeBool        = 8
	typeString      = 9
	typeUUID        = 10
	typeDate        = 11
	typeByteArray   = 12
	typeShortArray  = 13
	typeIntArray    = 14
	typeLongArray   = 15
	typeFloatArray  = 16
	typeDoubleArray = 17
	typeCharArray   = 18
	typeBoolArray   = 19
	typeStringArray = 20
	typeUUIDArray   = 21
	typeDateArray   = 22
	// TODO: Object array = 23
	// TODO: Collection = 24
	// TODO: Map = 25
	// TODO: Enum = 28
	// TODO: Enum Array = 29
	// TODO: Decimal = 30
	// TODO: Decimal Array = 31
	typeTimestamp      = 33
	typeTimestampArray = 34
	typeTime           = 36
	typeTimeArray      = 37
	typeNULL           = 101
)

// Char is Apache Ignite "char" type
type Char rune

// Date is Unix time, the number of MILLISECONDS elapsed
// since January 1, 1970 UTC.
type Date int64

// ToDate converts Golang time.Time to Apache Ignite Date
func ToDate(t time.Time) Date {
	t1 := t.UTC()
	t2 := t1.Unix() * 1000
	t2 += int64(t1.Nanosecond()) / int64(time.Millisecond)
	return Date(t2)
}

// Time is Apache Ignite Time type
type Time int64

// ToTime converts Golang time.Time to Apache Ignite Time
func ToTime(t time.Time) Time {
	t1 := t.UTC()
	t2 := time.Date(1970, 1, 1, t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), time.UTC)
	t3 := t2.Unix() * 1000
	t3 += int64(t2.Nanosecond()) / int64(time.Millisecond)
	return Time(t3)
}