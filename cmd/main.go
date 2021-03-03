package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LoadBalancer struct {
	workers []*Worker
	jobChan chan func()
}

type Worker struct {
	state    bool
	name     string
	taskChan chan func()
}

func NewWorker(num int) *Worker {

	worker := &Worker{name: fmt.Sprintf("worker-%d", num), taskChan: make(chan func()), state: true}
	go worker.lookup()
	fmt.Printf("%s初始化完毕\n", worker.name)
	return worker
}

func (w *Worker) lookup() {

	for {

		select {

		case task := <-w.taskChan:
			w.state = false

			fmt.Printf("%s 开始执行任务\n", w.name)
			task()
			w.state = true

		default:
			time.Sleep(time.Second)
		}
	}

}
func LB(workerNum int) *LoadBalancer {

	workers := make([]*Worker, 0)
	for i := 1; i <= workerNum; i++ {
		workers = append(workers, NewWorker(i))
	}

	lb := &LoadBalancer{workers: workers, jobChan: make(chan func())}
	go lb.lookup()

	return lb
}

func (lb *LoadBalancer) lookup() {

	for {

		select {
		case job := <-lb.jobChan:
			worker := lb.selectAvaliableWorker()
			worker.taskChan <- job
		}
	}

}

func (lb *LoadBalancer) selectAvaliableWorker() *Worker {

RETRY:
	for _, worker := range lb.workers {
		if worker.state {
			return worker
		}
	}

	fmt.Println("not found the available worker after 1s try again....")
	time.Sleep(time.Second)
	goto RETRY

}

func (lb *LoadBalancer) Submit(fun func()) {

	lb.jobChan <- fun
}

func main() {

	lb := LB(4)

	for i := 0; i < 100; i++ {
		lb.Submit(func() {

			time.Sleep(time.Second * 2)

			fmt.Printf("-----已完成任务---%d\n", rand.Int63())
		})
	}

	time.Sleep(time.Second * 100)

}
