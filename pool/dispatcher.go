package pool

import (
	"fmt"

	"github.com/Kevin8428/channels-and-worker-pools/work"
)

var WorkerChannel = make(chan chan work.Work)

const workerCount = 2

type Collection struct {
	Work chan work.Work
	End  chan bool
}

func BuildCollector() Collection {
	workers := []Worker{}
	for i := 0; i < workerCount; i++ { // create workerss
		worker := Worker{
			ID:            i,
			WorkerChannel: WorkerChannel,
			Channel:       make(chan work.Work),
			End:           make(chan bool),
		}
		workers = append(workers, worker)
		worker.Start()
	}
	work := make(chan work.Work)
	end := make(chan bool)
	collection := Collection{
		Work: work,
		End:  end,
	}
	go func() {
		for {
			ch := <-WorkerChannel
			select {
			case job := <-work:
				ch <- job
			case <-end:
				fmt.Println("end reached")
				for _, w := range workers {
					w.Stop()
				}
			}
		}
	}()
	return collection
}
