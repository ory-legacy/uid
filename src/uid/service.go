package uid

import (
	"time"
	"sync"
	"fmt"
)

type Service struct {
	sync.Mutex
	// Current count relative to timestamp
	count          int16
	timestampForCounting int32

	// This nodes ID
	Id             int16
}

type CreatorArguments struct {
	Type  int16
}

//
func (t *Service) Create(args *CreatorArguments, reply *Uid) error {
	now := int32(time.Now().Unix())

	t.Mutex.Lock()
	if now == t.timestampForCounting {
		t.count++
	} else {
		t.count = 0
		t.timestampForCounting = now
	}
	t.Mutex.Unlock()

	id := t.createId(int64(args.Type), int64(t.Id), int64(now), int64(t.count))

	*reply = Uid(id)
	fmt.Printf("Identifier %s (%d) created with params id=%d type=%d timestamp=%d count=%d\n", id, id, t.Id, args.Type, now, t.count)
	fmt.Printf("Identifier %s (%d) created with params id=%d type=%d timestamp=%d count=%d\n", reply, reply, reply.NodeId(), reply.Type(), reply.Timestamp(), reply.MicroTime())

	return nil
}
