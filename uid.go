package uid

import (
	"strconv"
	"math"
	"errors"
)

// Identifier is an interface for setting and getting uids
type Identifier interface {
	Id() Uid
	SetId(uid Uid)
}

// The Uid type is a long
type Uid uint64

var intToCharMap = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B',
	'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3',
	'4', '5', '6', '7', '8', '9', '_', '-'}

// Next 9 bits are type
// Maximum value: 511
var typeOffset uint32 = 64 - 9
var typeMaxValue int64 = maxUnSignedValue(9)

// Next 8 bits are node id (the service node which created this uid)
// Maximum value: 511
var nodeOffset uint32 = typeOffset - 9
var nodeMaxValue int64 = maxUnSignedValue(9)

// Next 32 bits are the timestamp
// Maximum value: signed int32
var timestampOffset uint32 = nodeOffset - 32
var timestampMaxValue int64 = maxSignedValue(32)

// Next 14 bits are the offset (or "microtime") to ensure uniqueness
// Maximum value: 16383
// var uniqueOffset = timestampOffset - 14
var uniqueOffsetMaxValue int64 = maxUnSignedValue(14)

// New creates a new unique identifier taking into account
// the uid's type, the node which created this uid, the timestamp
// of the creation and an offset to ensure uniqueness
func New(uidType int64, Node int64, timestamp int64, offset int64) (Uid, error) {
	var err error

	if int64(timestamp) > timestampMaxValue {
		return 0, errors.New("Timestamp overflow")
	}

	if int64(Node) > nodeMaxValue {
		return 0, errors.New("Node overflow")
	}

	if int64(uidType) > typeMaxValue {
		return 0, errors.New("Type overflow")
	}

	if int64(offset) > uniqueOffsetMaxValue {
		return 0, errors.New("Offset overflow")
	}

	return Uid(((int64(uidType)&typeMaxValue)<<(typeOffset)) |
			((int64(Node)&nodeMaxValue)<<(nodeOffset)) |
			((int64(timestamp)&timestampMaxValue)<<(timestampOffset)) |
			(int64(offset)&uniqueOffsetMaxValue)), err
}

// Type returns the uid's type
func (t Uid) Type() int64 {
	return ((int64(t) >> (typeOffset)) & typeMaxValue)
}

// Node returns the node on which this uid has been created on
func (t Uid) Node() int64 {
	return ((int64(t) >> (nodeOffset)) & nodeMaxValue)
}

// Timestamp returns the timestamp of this uid
func (t Uid) Timestamp() int64 {
	return ((int64(t) >> (timestampOffset)) & timestampMaxValue)
}

// Offset returns the unique offset of this uid
func (t Uid) Offset() int64 {
	return (int64(t) & uniqueOffsetMaxValue)
}

// String: Uid implements the Stringer interface
func (t Uid) String() string {
	return strconv.FormatInt(int64(t), 10)
}

func (t Uid) MarshalText() (text []byte, err error) {
	text = make([]byte,11)
	length := int64(len(intToCharMap))
	id := int64(t)

	for i := 0; i < 11 ; i++ {
		rest := id % length
		text[i] = intToCharMap[rest]
		id = id/length
	}

	return text, err
}

func (t Uid) UnmarshalText(text []byte) (err error) {
	length := int64(len(intToCharMap))
	id := int64(t)

	for i := 0; i < 11 ; i++ {
		rest := id % length
		text[i] = intToCharMap[rest]
		id = id/length
	}

	return err
}

// maxSignedValue returns the maximum value of n signed bits
// Example: maxValue(8) == 127
func maxSignedValue(numberOfBits int32) int64 {
	return int64(math.Pow(2, float64(numberOfBits - 1))) - 1
}

// maxUnSignedValue returns the maximum value of n unsigned bits
// Example: maxValue(8) == 255
func maxUnSignedValue(numberOfBits int32) int64 {
	return int64(math.Pow(2, float64(numberOfBits))) - 1
}
