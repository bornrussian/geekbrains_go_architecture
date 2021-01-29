package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
	"workerpool/accountant"
	"workerpool/leader"

	"workerpool/config"
)

func main() {
	// Парсим конфиг-файл
	flagConfigPath := flag.String("c", "./config.yaml", "config file with yaml format")
	flag.Parse()
	cfg, err := config.ReadConfig(*flagConfigPath)
	if err != nil {
		log.Fatalf("Config file read failed: %s", err)
	}

	// Мы делаем консольный софт, потому обеспечиваем возможность прерваться по нажатию Ctrl+C
	var cancel context.CancelFunc
	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt)
	go func() {
		oscall := <-ctrlC
		log.Printf("System call: %+v", oscall)
		cancel()
	}()

	// Запускаем нужный режим:
	log.Printf("Running in mode '%v' %v ...", cfg.Options.Mode, cfg.Options.Until)
	log.Printf("Target = %v", cfg.Victim.HTTPUrl)
	accountant := accountant.NewAccountant()
	var ctx context.Context
	switch cfg.Options.Mode {
	case "until_timeout": // Режим по срабатыванию таймера
		ctx, cancel = context.WithTimeout(context.Background(),	time.Duration(cfg.Options.Until) * time.Second)
	case "until_requests_sent": // Режим по выполнению нужного количества запросов к Victim
		ctx, cancel = context.WithCancel(context.Background())
		go accountant.ControlTotalJobsAndThenRun(cfg.Options.Until, cancel)
	default:
		log.Fatal("Unexpected mode in config.yaml ! Use 'till_timeout' or 'till_requests_sent'.")
	}

	wg := &sync.WaitGroup{}

	leader.Run(ctx, wg, cfg, accountant)

	wg.Wait()

	accountant.LogResults(cfg)
}
