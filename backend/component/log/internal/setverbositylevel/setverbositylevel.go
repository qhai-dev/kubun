package setverbositylevel

import "sync"

var (
	Mutex sync.Mutex

	Callbacks []func(v uint32) error
)
