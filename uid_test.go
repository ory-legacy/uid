package uid

import (
	"testing"
	"fmt"
)

// TestNew tests the uids New method
func TestNew(t *testing.T) {
	var idType int64 = 1
	var timestamp int64 = 510239510
	var Node int64 = 5
	var offset int64 = 12345

	obj, _ := New(idType, Node, timestamp, offset)

	if(obj.Type() != idType) {
		t.Error(fmt.Sprintf("Expected type %d to match %d", obj.Type(), idType))
	}

	if(obj.Timestamp() != timestamp) {
		t.Error(fmt.Sprintf("Expected timestamp %d to match %d", obj.Timestamp(), timestamp))
	}

	if(obj.Node() != Node) {
		t.Error(fmt.Sprintf("Expected node %d to match %d", obj.Node(), Node))
	}

	if(obj.Offset() != offset) {
		t.Error(fmt.Sprintf("Expected offset %d to match %d", obj.Offset(), offset))
	}
}


func TestMarshalling(t *testing.T) {
	var idType int64 = 1
	var timestamp int64 = 510239510
	var Node int64 = 5
	var offset int64 = 12345

	obj, _ := New(idType, Node, timestamp, offset)
	text, _ := obj.MarshalText();

	marhsalled := new(Uid)
	_ = marhsalled.UnmarshalText(text)

	if *marhsalled != *obj {
		t.Error(fmt.Sprintf("Marshalling failed: %d should be %d", *marhsalled, *obj))
	}
}

// TestMaxSignedValue tests the internal maxSignedValue method
func TestMaxSignedValue(t *testing.T) {
	if maxSignedValue(8) != 127 {
		t.Error(fmt.Sprintf("maxSignedValue(2) should be 3 but is %d", maxSignedValue(2)))
	}
}

// TestMaxSignedValue tests the internal maxUnSignedValue method
func TestMaxUnSignedValue(t *testing.T) {
	if maxUnSignedValue(8) != 255 {
		t.Error(fmt.Sprintf("maxUnSignedValue(8) should be 255 but is %d", maxUnSignedValue(2)))
	}
}

// TestOverflows tests if overflows are detected correctly
func TestOverflows(t *testing.T) {
	var idType int64 = maxUnSignedValue(16)
	var timestamp int64 = 1
	var Node int64 = 1
	var offset int64 = 1
	var err error

	_, err = New(idType, Node, timestamp, offset)
	if err == nil {
		t.Error("Type overflow not detected")
	}

	idType = 1
	timestamp = maxSignedValue(32) + 1
	_, err = New(idType, Node, timestamp, offset)
	if err == nil {
		t.Error("Timestamp overflow not detected")
	}

	timestamp = 1
	Node = maxUnSignedValue(9) + 1
	_, err = New(idType, Node, timestamp, offset)
	if err == nil {
		t.Error("Node overflow not detected")
	}

	Node = 1
	offset = maxUnSignedValue(32) + 1
	_, err = New(idType, Node, timestamp, offset)
	if err == nil {
		t.Error("Offset overflow not detected")
	}
}

func BenchmarkTestNew(b *testing.B) {
	b.StopTimer()
	var idType int64 = 1
	var timestamp int64 = 1
	var Node int64 = 1
	var offset int64 = 1
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = New(idType, Node, timestamp, offset)
	}
}
