package util

import (
	"sync"
	"time"

	"github.com/oklog/ulid"
	"golang.org/x/exp/rand"
)


var ulidPool = sync.Pool{
	New: func() interface{} {
		entropy := ulid.Monotonic(rand.New(rand.NewSource(uint64(time.Now().UnixNano()))), 0)
		return ulid.MustNew(ulid.Now(), entropy)
	},
}

func GetUlid() string {
	id := ulidPool.Get().(ulid.ULID)
	//? Ensure uniqueness with a bit of performance tradeoff
	time.Sleep(1 * time.Millisecond)
	ulidPool.Put(id)

	return id.String()
}