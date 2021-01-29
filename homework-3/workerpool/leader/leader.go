package leader

import (
	"context"
	"log"
	"sync"
	"time"
	"workerpool/accountant"
	"workerpool/config"
	"workerpool/worker"
)

func Run(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, acnt *accountant.Accountant) {
	jobCh := make(chan config.Victim)

	for i := 0; i < int(cfg.Options.Threads); i++ {
		wg.Add(1)
		w := worker.NewWorker(i, wg, jobCh, acnt)
		go w.HandleJobs()
	}

	for {
		select {
		case <-ctx.Done():
			close(jobCh)
			log.Println("Gracefully stop:", ctx.Err())
			return
		default:
			jobCh <- cfg.Victim
			time.Sleep(100 * time.Millisecond )
		}
	}
}
