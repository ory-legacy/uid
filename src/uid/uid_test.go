package uid

import (
	"testing"
	"fmt"
)

func TestNew(t *testing.T) {
	var idType int64 = 1
	var timestamp int64 = 510239510
	var nodeId int64 = 5
	var offset int64 = 12345

	obj, _ := New(idType, nodeId, timestamp, offset)

	if(obj.Type() != idType) {
		t.Error(fmt.Sprintf("Expected type %d to match %d", obj.Type(), idType))
	}

	if(obj.Timestamp() != timestamp) {
		t.Error(fmt.Sprintf("Expected timestamp %d to match %d", obj.Timestamp(), timestamp))
	}

	if(obj.NodeId() != nodeId) {
		t.Error(fmt.Sprintf("Expected node %d to match %d", obj.NodeId(), nodeId))
	}

	if(obj.Offset() != offset) {
		t.Error(fmt.Sprintf("Expected offset %d to match %d", obj.Offset(), offset))
	}
}

func TestMaxSignedValue(t *testing.T) {
	if maxSignedValue(8) != 127 {
		t.Error(fmt.Sprintf("maxSignedValue(2) should be 3 but is %d", maxSignedValue(2)))
	}
}

func TestMaxUnSignedValue(t *testing.T) {
	if maxUnSignedValue(8) != 255 {
		t.Error(fmt.Sprintf("maxUnSignedValue(8) should be 255 but is %d", maxUnSignedValue(2)))
	}
}

func TestOverflows(t *testing.T) {
	var idType int64 = maxUnSignedValue(16)
	var timestamp int64 = 1
	var nodeId int64 = 1
	var offset int64 = 1
	var err error

	_, err = New(idType, nodeId, timestamp, offset)
	if err == nil {
		t.Error("Type overflow not detected")
	}

	idType = 1
	timestamp = maxSignedValue(32) + 1
	_, err = New(idType, nodeId, timestamp, offset)
	if err == nil {
		t.Error("Timestamp overflow not detected")
	}

	timestamp = 1
	nodeId = maxUnSignedValue(8) + 1
	_, err = New(idType, nodeId, timestamp, offset)
	if err == nil {
		t.Error("Node overflow not detected")
	}

	nodeId = 1
	offset = maxUnSignedValue(32) + 1
	_, err = New(idType, nodeId, timestamp, offset)
	if err == nil {
		t.Error("Offset overflow not detected")
	}
}
