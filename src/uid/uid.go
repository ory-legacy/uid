package uid

import (
	"strconv"
	"math"
	"errors"
	"fmt"
)

type Identifier interface {
	Id() Uid
	SetId(uid Uid)
}

// The Uid type is a long
type Uid int64

// Highest bit is 0 for serialization algorithm
var serializationOffset uint32 = 64 - 1

// Next 9 bits are type
// Maximum value: 511
var typeOffset uint32 = serializationOffset - 9

// Next 8 bits are node id (the service node which created this uid)
// Maximum value: 255
var nodeOffset uint32 = typeOffset - 6

// Next 32 bits are the timestamp
// Maximum value: int32
var timestampOffset uint32 = nodeOffset - 32

// Next 14 bits are the offset (or "microtime") to ensure uniqueness
// Maximum value: 16383
// var uniqueOffset = timestampOffset - 14

// New creates a new unique identifier taking into account
// the uid's type, the node which created this uid, the timestamp
// of the creation and an offset to ensure uniqueness
func New(uidType int64, nodeId int64, timestamp int64, offset int64) (Uid, error) {
	var err error

	if int64(timestamp) > maxSignedValue(32) {
		return 0, errors.New("Timestamp overflow")
	}

	if int64(nodeId) > maxUnSignedValue(8) {
		return 0, errors.New("Node overflow")
	}

	if int64(uidType) > maxUnSignedValue(9) {
		return 0, errors.New("Type overflow")
	}

	if int64(offset) > maxUnSignedValue(14) {
		return 0, errors.New("Offset overflow")
	}

	return Uid(((int64(uidType) & int64(511)) << (typeOffset)) |
			((int64(nodeId) & int64(255)) << (nodeOffset)) |
			((int64(timestamp) & int64(0xFFFFFFFF)) << (timestampOffset)) |
			(int64(offset) & int64(16383))), err
}

func (t Uid) Type() int64 {
	return ((int64(t) >> (typeOffset)) & int64(0x1FF))
}

func (t Uid) NodeId() int64 {
	return ((int64(t) >> (nodeOffset)) & int64(0x1F))
}

func (t Uid) Timestamp() int64 {
	return ((int64(t) >> (timestampOffset)) & int64(0xFFFFFFFF))
}

func (t Uid) Offset() int64 {
	return (int64(t) & int64(0xFFFF))
}

func (t Uid) String() string {
	fmt.Printf("%d\n", int64(t))
	fmt.Printf("%s\n", strconv.FormatInt(int64(t), 10))
	return strconv.FormatInt(int64(t), 10)
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
