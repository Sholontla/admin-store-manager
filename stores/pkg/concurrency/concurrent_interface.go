package concurrency

import (
	"reflect"
	"sync"
)

type IPoolService interface {
	GenericPool(t reflect.Type) (*sync.Pool, error)
}
