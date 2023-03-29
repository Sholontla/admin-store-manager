package pool

import (
	"log"
	"reflect"
	"sync"
)

// type WorkerPoolChann struct {
// 	responseChan chan struct{}
// }

func genericPool(t reflect.Type) (*sync.Pool, error) {
	f := &sync.Pool{
		New: func() interface{} {
			return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, t), 0).Interface()
		},
	}
	return f, nil
}

func WorkerPool(numWorkers int, request chan map[string]interface{}, workerFunc func(orderRequestChan <-chan map[string]interface{}) (chan map[string]interface{}, error)) (chan map[string]interface{}, error) {
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
			worker := pool.Get().(chan map[string]interface{})
			defer pool.Put(worker)

			workerFunc(request)
			if err != nil {
				log.Printf("error coming from worker pool: %s ", err)
			}

		}()
	}

	return make(chan map[string]interface{}), nil
}
