package vars

import "sync"

var (
	ThreadNum = 100
	Result    *sync.Map
)

func init() {
	Result = &sync.Map{}
}
