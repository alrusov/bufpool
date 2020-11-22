package bufpool

import (
	"bytes"
	"sync"
	"sync/atomic"
)

//----------------------------------------------------------------------------------------------------------------------------//

type (
	// Stat --
	Stat struct {
		Issued   int64 `json:"issued"`
		Released int64 `json:"released"`
		InUse    int64 `json:"inUse"`
	}

	// GetStatFunc --
	GetStatFunc func() (stat *Stat)
)

var (
	enabled  = true
	bufPool  = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
	issued   = int64(0)
	released = int64(0)
)

//----------------------------------------------------------------------------------------------------------------------------//

// GetBuf --
func GetBuf() (b *bytes.Buffer) {
	defer func() {
		if enabled && b != nil {
			atomic.AddInt64(&issued, 1)
		}
	}()

	if enabled {
		ok := false
		b, ok = bufPool.Get().(*bytes.Buffer)
		if ok {
			return
		}
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

// Enable --
func Enable() {
	enabled = true
}

// Disable --
func Disable() {
	enabled = false
}

//----------------------------------------------------------------------------------------------------------------------------//

// GetStat --
func GetStat() (stat *Stat) {
	return &Stat{
		Issued:   atomic.LoadInt64(&issued),
		Released: atomic.LoadInt64(&released),
	}
}

//----------------------------------------------------------------------------------------------------------------------------//
