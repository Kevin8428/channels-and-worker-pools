package main

import (
	"fmt"
	"math/rand"

	"github.com/Kevin8428/channels-and-worker-pools/pool"
	"github.com/Kevin8428/channels-and-worker-pools/work"
)

func main() {

	collector := pool.BuildCollector()
	fmt.Println("collector: ", collector)

	jobs := BuildJobs()
	for i := 0; i < len(jobs); i++ {
		collector.Work <- work.Work{ID: i, JobName: jobs[i]}
	}
}

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func BuildJobs() []string {
	jobs := []string{}
	jobLength := 8
	for i := 0; i < 6; i++ {
		job := ""
		for j := 0; j < jobLength; j++ {
			l := letters[rand.Intn(len(letters))]
			job += l
		}
		jobs = append(jobs, job)
	}
	return jobs
}
