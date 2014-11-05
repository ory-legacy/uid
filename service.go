package uid

import (
	"time"
	"sync"
)

// Service represents a RPC node for creating uids
type Service struct {
	// Needed for locking count
	sync.Mutex

	// Current count relative to timestamp
	count                int64
	timestampForCounting int64

	// This nodes ID
	Id             int64
}

// CreatorArguments is a structure for the RPC Create function
type CreatorArguments struct {
	Type  int64
}

// Create is a RPC wrapper for uid's New method
func (t *Service) New(args *CreatorArguments, reply *Uid) error {
	now := int64(time.Now().Unix())

	t.Mutex.Lock()
	if now == t.timestampForCounting {
		t.count++
	} else {
		t.count = 0
		t.timestampForCounting = now
	}
	t.Mutex.Unlock()

	id, err := New(args.Type, t.Id, now, t.count)

	if err != nil {
		return err
	}

	*reply = *id

	return nil
}
