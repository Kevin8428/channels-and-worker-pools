package work

import (
	"fmt"
)

type Work struct {
	ID      int
	JobName string
}

func Process(job Work) {
	fmt.Println("processing job: ", job.JobName)
}
