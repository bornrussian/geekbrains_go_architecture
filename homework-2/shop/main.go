package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"shop/config"
	"shop/repository"
	"shop/tools/email"
	"shop/tools/tgbot"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	flagConfigPath := flag.String("c", "./config.yaml", "config file with yaml format")
	flag.Parse()
	cfg, err := config.ReadConfig(*flagConfigPath)
	if err != nil {
		panic(fmt.Sprintf("Config file read failed: %s", err))
	}

	bot, err := tgbot.NewShopTgBot(cfg.Telegram)
	if err != nil {
		log.Fatal("Unable to init telegram bot")
	}

	mail, err := email.NewShopEmailServer(cfg.Email)
	if err != nil {
		log.Fatal("Unable to init SMTP server")
	}

	handler := &shopHandler{
		db: repository.NewMapDB(),
		bot: bot,
		mail: mail,
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	router.HandleFunc("/order", handler.createOrderHandler).Methods("POST")
	router.HandleFunc("/order/{id}", handler.getOrderHandler).Methods("GET")

	srv := &http.Server{
		Addr:         cfg.Server.Listen,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	if cfg.Server.DebugMode {
		log.Printf("Listening %v%v ...", cfg.Server.Domain, cfg.Server.Listen)
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
