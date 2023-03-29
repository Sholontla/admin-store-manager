package pool

import (
	"log"
	"reflect"
	"sync"
)

func genericPool(t reflect.Type) (*sync.Pool, error) {
	f := &sync.Pool{
		New: func() interface{} {
			return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, t), 0).Interface()
		},
	}
	return f, nil
}

func ConnectionPool(t reflect.Type) (*sync.Pool, error) {
	pool := &sync.Pool{
		New: func() interface{} {
			return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, t), 0).Interface()
		},
	}
	return pool, nil
}

func WorkerPool(numWorkers int, request chan interface{}, workerFunc func(orderRequestChan <-chan interface{}) (chan interface{}, error)) (chan interface{}, error) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	// Define a pool to reuse workers

	pool, err := genericPool(reflect.TypeOf(map[string]interface{}{}))
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			// Get a worker from the pool or create a new one
			worker := pool.Get().(chan interface{})
			defer pool.Put(worker)

			workerFunc(request)
			if err != nil {
				log.Printf("error coming from worker pool: %s ", err)
			}
		}()
	}
	var response chan interface{}
	return response, nil
}
