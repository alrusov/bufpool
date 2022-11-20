package bufpool

import (
	"bytes"
	"sync"
	"sync/atomic"
)

//----------------------------------------------------------------------------------------------------------------------------//

var (
	enabled = true

	bufPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	issued   = int64(0)
	released = int64(0)
)

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

// GetBuf --
func GetBuf() (b *bytes.Buffer) {
	if enabled {
		b = bufPool.Get().(*bytes.Buffer)
		atomic.AddInt64(&issued, 1)
		return
	}

	return new(bytes.Buffer)
}

// PutBuf --
func PutBuf(b *bytes.Buffer) {
	if enabled && b != nil {
		b.Reset()
		bufPool.Put(b)
		atomic.AddInt64(&released, 1)
	}
}

//----------------------------------------------------------------------------------------------------------------------------//

type (
	stat struct {
		Issued   int64 `json:"issued"`
		Released int64 `json:"released"`
		InUse    int64 `json:"inUse"`
	}
)

// GetStat --
func GetStat() any {
	stat := &stat{
		Issued:   atomic.LoadInt64(&issued),
		Released: atomic.LoadInt64(&released),
	}
	stat.InUse = stat.Issued - stat.Released
	return stat
}

//----------------------------------------------------------------------------------------------------------------------------//
