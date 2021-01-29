package accountant

import (
	"context"
	"log"
	"sync"
	"time"
	"workerpool/config"
)

type Accountant struct {
	mutex     sync.Mutex
	started   time.Time
	okJobs    jobResults
	faultJobs jobResults
}

type jobResults struct {
	count     int64
	spentTime time.Duration
}

func NewAccountant() *Accountant{
	return &Accountant {
		started:   time.Now(),
		okJobs: jobResults{
			count:     0,
			spentTime: time.Duration(0),
		},
		faultJobs: jobResults{
			count:     0,
			spentTime: time.Duration(0),
		},
	}
}

func (c *Accountant) AddOk (duration time.Duration) {
	c.mutex.Lock()
	c.okJobs.count += 1
	c.okJobs.spentTime += duration
	c.mutex.Unlock()
}

func (c *Accountant) AddFault (duration time.Duration) {
	c.mutex.Lock()
	c.faultJobs.count += 1
	c.faultJobs.spentTime += duration
	c.mutex.Unlock()
}

func (c *Accountant) AVGSpentTimePerJob() time.Duration {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	time := c.okJobs.spentTime + c.faultJobs.spentTime
	jobs := c.okJobs.count + c.faultJobs.count
	return divDuration(time, jobs)
}

func (c *Accountant) ControlTotalJobsAndThenRun(jobsLimit int64, cancel context.CancelFunc) {
	defer cancel()

	for {
		c.mutex.Lock()
		jobs := c.okJobs.count + c.faultJobs.count
		c.mutex.Unlock()

		if jobs >= jobsLimit {
			log.Printf("%v requests completed.\n",jobsLimit)
			break
		}
		time.Sleep(500*time.Millisecond)
	}
}

func (c *Accountant) LogResults(cfg *config.Config) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Выводим итоги
	log.Printf("\n")
	log.Printf("=== RESULTS: ===\n")
	log.Printf("\n")
	log.Printf("Mode was %v = %v", cfg.Options.Mode, cfg.Options.Until)
	log.Printf("\n")
	log.Printf("Total requests per seconds (RPS) made: %v\n",
		calcRPS(c.okJobs.count+c.faultJobs.count, time.Since(c.started)))
	log.Printf("\n")
	log.Printf("Okay requests:\n")
	log.Printf("\tcount: %v\n",c.okJobs.count)
	log.Printf("\tavg time: %v\n",divDuration(c.okJobs.spentTime, c.okJobs.count))
	log.Printf("\n")
	log.Printf("Failed requests:\n")
	log.Printf("\tcount: %v\n",c.faultJobs.count)
	log.Printf("\tavg time: %v\n",divDuration(c.faultJobs.spentTime, c.faultJobs.count))
}



func divDuration(duration time.Duration, divBy int64) time.Duration {
	if divBy == 0 { return 0 }
	ms := int64( duration / time.Millisecond )
	divResult := ms / divBy
	return time.Duration( divResult * int64(time.Millisecond) )
}

func calcRPS(requests int64, duration time.Duration) int64 {
	if requests == 0 { return 0 }
	seconds := int64(duration / time.Second)
	return requests/seconds
}