package uid

import "strconv"

type Identifier interface {
	Id() Uid
	SetId(uid Uid)
}

type Uid int64

// New creates a new unique identifier taking into account
// the uid's type, the node which created this uid, the timestamp
// of the creation and an offset to ensure uniqueness
func New(uidType int16, nodeId byte, timestamp, offset int64) int64 {

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

	return ((int64(uidType) & int64(0x1FF)) << (64-1-9)) |
			((int64(nodeId) & int64(0x1F)) << (64-1-9-1-5)) |
			((int64(timestamp) & int64(0xFFFFFFFF)) << (64-1-9-1-5-32)) |
			(int64(offset) & int64(0xFFFF))
}

func (t *Uid) Type() int16 {
	return int16((int64(t) >> (64 - 9 - 1)) & int64(0x1FF))
}

func (t *Uid) NodeId() byte {
	return byte((int64(t) >> (64 - 1 - 9 - 1 - 5)) & int64(0x1F))
}

func (t *Uid) Timestamp() int64 {
	return ((int64(t) >> (64 - 1 - 9 - 1 - 5 - 32)) & int64(0xFFFFFFFF))
}

func (t *Uid) Offset() int64 {
	return (int64(t) & int64(0xFFFF))
}

func (t *Uid) String() string {
	return strconv.FormatInt(int64(t), 10)
}
