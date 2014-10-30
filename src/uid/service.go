package uid

import (
	"time"
	"sync"
	"fmt"
)

type Service struct {
	sync.Mutex
	// Current count relative to timestamp
	count                int64
	timestampForCounting int64

	// This nodes ID
	Id             int64
}

type CreatorArguments struct {
	Type  int64
}

//
func (t *Service) Create(args *CreatorArguments, reply *Uid) error {
	now := int64(time.Now().Unix())

	t.Mutex.Lock()
	if now == t.timestampForCounting {
		t.count++
	} else {
		t.count = 0
		t.timestampForCounting = now
	}
	t.Mutex.Unlock()

	*reply, _ = New(args.Type, t.Id, now, t.count)

	fmt.Printf("Identifier %s (%d) requested with params id=%d type=%d timestamp=%d count=%d\n", reply, reply, t.Id, args.Type, now, t.count)
	fmt.Printf("Identifier %s (%d) created with params id=%d type=%d timestamp=%d count=%d\n", reply, reply, reply.NodeId(), reply.Type(), reply.Timestamp(), reply.Offset())

	return nil
}
