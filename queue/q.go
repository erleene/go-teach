package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

var wg sync.WaitGroup

//Queue ...
type Queue struct {
	m     *sync.Mutex
	Items []interface{}
}

//NewQueue ...
func NewQueue(nums ...int) *Queue {

	size := 0
	if len(nums) > 0 {
		if nums[0] > 0 {
			size = nums[0]
		}
	}
	return &Queue{Items: make([]interface{}, size), m: new(sync.Mutex)}
}

// curl -XPOST localhost:3000/ -d '{ "test": "alex" }'
//Enqueue ...
func (q *Queue) Enqueue(item interface{}) {
	q.m.Lock()
	defer q.m.Unlock()
	q.Items = append(q.Items, item)
	log.Printf("%v - > queue length %d", q.Items, len(q.Items))
	wg.Done()
}

//curl localhost:3000/
//Dequeue ...
func (q *Queue) Dequeue() (interface{}, error) {
	q.m.Lock()
	defer q.m.Unlock()
	if len(q.Items) == 0 {
		return nil, errors.New("Zero queue length")
	}
	front := q.Items[0]
	q.Items = q.Items[1:]

	return front, nil

}

func main() {
	q := NewQueue(0)
	var i int
	//loop worker, create 50  workers to enqueues
	for i = 1; i < 50; i++ {
		wg.Add(1)
		go q.Enqueue(i)
	}

	_, er := q.Dequeue()
	if er != nil {
		fmt.Println(er)
	}

	_, er = q.Dequeue()
	if er != nil {
		fmt.Println(er)
	}

	// fmt.Println(item)
	// fmt.Println(item2)

	//workers enqueueing and dequeueing
	//dequeue - problematic when trying to enqueue at the same time
	wg.Wait()

}
