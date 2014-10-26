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
	fmt.Printf("Identifier %s (%d) created with params id=%d type=%d timestamp=%d count=%d\n", reply, reply, reply.Creator(), reply.Type(), reply.Timestamp(), reply.MicroTime())

	return nil
}

func (t *Service) createId(uidType, nodeId, timestamp, timestampCount int64) int64 {

	// Highest bit is 0 for serialization algorithm //
	// 0l << (64 - 1) |

	// Highest 9 bits are type //
	//(int64(args.Type) & 511) << (64 - 1 - 9)) |

	// Next bit is empty so that the first two alphanumerics only depend on
	// type and not also the source

	// 0l << (64 - 1 - 9 -1) */
	// Next 5 bits are source */
	//(int64(t.Id & 31) << (64 - 1 - 9 - 1 - 5)) |

	// Next 4 bytes are TS */
	//((now & int64(0xFFFFFFFF)) << (64 - 1 - 9 - 1 - 5 - 32)) |

	// Next 2 bytes are "MicroTime" */
	//ID & int64(0xFFFF)

	return ((uidType & int64(0x1FF)) << (64-1-9)) |
		((nodeId & int64(0x1F)) << (64-1-9-1-5)) |
		((timestamp & int64(0xFFFFFFFF)) << (64-1-9-1-5-32)) |
		(timestampCount & int64(0xFFFF))
}
