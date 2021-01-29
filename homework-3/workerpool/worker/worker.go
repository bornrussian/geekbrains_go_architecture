package worker

import (
	"log"
	"net/http"
	"sync"
	"time"
	"workerpool/accountant"
	"workerpool/config"
)

type Worker struct {
	id    int
	jobCh <-chan config.Victim
	accountant *accountant.Accountant
	wg *sync.WaitGroup
}

func (w *Worker) HandleJobs() {
	defer w.wg.Done()
	for job := range w.jobCh {
		started := time.Now()
		// TODO: добавить разные методы, помимо http.Get
		_, err := http.Get(job.HTTPUrl)
		elapsed := time.Since(started)
		if err != nil {
			w.accountant.AddFault(elapsed)
			log.Printf("Worker #%v fault in %v (avg=%v)\n", w.id, elapsed, w.accountant.AVGSpentTimePerJob())
		} else {
			w.accountant.AddOk(elapsed)
			log.Printf("Worker #%v success in %v (avg=%v)\n", w.id, elapsed, w.accountant.AVGSpentTimePerJob())
		}
	}
}

func NewWorker(id int, wg *sync.WaitGroup, jobCh <-chan config.Victim, accountant *accountant.Accountant) *Worker {
	return &Worker{
		id:    id,
		wg: wg,
		jobCh: jobCh,
		accountant: accountant,
	}
}
