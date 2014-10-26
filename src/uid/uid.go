package uid

import "strconv"

type Identifier interface {
	Id() Uid
	SetId(uid Uid)
}

type Uid int64

func (t Uid) Type() int16 {
	return int16((int64(t) >> (64 - 9 - 1)) & int64(0x1FF))
}

func (t Uid) Creator() byte {
	return byte((int64(t) >> (64 - 1 - 9 - 1 - 5)) & int64(0x1F))
}

func (t Uid) Timestamp() int64 {
	return ((int64(t) >> (64 - 1 - 9 - 1 - 5 - 32)) & int64(0xFFFFFFFF))
}

func (t Uid) MicroTime() int64 {
	return (int64(t) & int64(0xFFFF))
}

func (t Uid) String() string {
	return strconv.FormatInt(int64(t), 10)
}
