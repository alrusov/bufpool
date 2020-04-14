package bufpool

import (
	"bytes"
	"sync"
)

//----------------------------------------------------------------------------------------------------------------------------//

var (
	enabled = false
	bufPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
)

//----------------------------------------------------------------------------------------------------------------------------//

// GetBuf --
func GetBuf() *bytes.Buffer {
	if enabled {
		if bb, ok := bufPool.Get().(*bytes.Buffer); ok {
			return bb
		}
	}
	return new(bytes.Buffer)
}

// PutBuf --
func PutBuf(b *bytes.Buffer) {
	if enabled {
		b.Reset()
		bufPool.Put(b)
	}
}

//----------------------------------------------------------------------------------------------------------------------------//

// Enable --
func Enable() {
	enabled = true
}

// Disable --
func Disable() {
	enabled = false
}

//----------------------------------------------------------------------------------------------------------------------------//
