package pool

import (
	"fmt"

	"github.com/Kevin8428/channels-and-worker-pools/work"
)

type Worker struct {
	ID            int
	WorkerChannel chan chan work.Work
	Channel       chan work.Work
	End           chan bool
}

func (worker *Worker) Start() {
	go func() {
		for {
			worker.WorkerChannel <- worker.Channel
			select {
			case job := <-worker.Channel:
				fmt.Printf("going to process using channel %+v\n", worker.Channel)
				work.Process(job)
			case end := <-worker.End:
				fmt.Println("done: ", end)
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.End <- true
}
