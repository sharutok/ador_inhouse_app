package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 8282, "server listining to port")
	flag.StringVar(&cfg.env, "env", "development", "Application Environment")
	flag.Parse()

	//LOGER FUNCTION
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	//SERVER CONFIG
	ser := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.Router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 1000,
		WriteTimeout: time.Second * 10000,
	}
	logger.Println("listining to port", cfg.port)

	//START SERVER
	err := ser.ListenAndServe()

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("Started")
}
