package bufpool

import (
	"bytes"
	"sync"
)

//----------------------------------------------------------------------------------------------------------------------------//

var bufPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

//----------------------------------------------------------------------------------------------------------------------------//

func GetBuf() *bytes.Buffer {
	if bb, ok := bufPool.Get().(*bytes.Buffer); ok {
		return bb
	}
	return new(bytes.Buffer)
}

func PutBuf(b *bytes.Buffer) {
	b.Reset()
	bufPool.Put(b)
}

//----------------------------------------------------------------------------------------------------------------------------//
