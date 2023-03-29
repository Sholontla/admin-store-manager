package concurrency

import (
	"reflect"
	"sync"
)

func GenericPool(t reflect.Type) (*sync.Pool, error) {
	f := &sync.Pool{
		New: func() interface{} {
			return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, t), 0).Interface()
		},
	}
	return f, nil
}
