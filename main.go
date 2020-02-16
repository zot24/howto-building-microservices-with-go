package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// global timeouts
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// no blocking goroutine (thread) to start the server
	go func() {
		err := s.ListenAndServe() // blocking function
		if err != nil {
			l.Println("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// it will immediatly terminate as there is nothing to stop the execution
	// so we need to block the execution and listen for the os signal events
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// reading from a channel will block the execution of the app
	// til we receive a notification
	sig := <-sigChan
	l.Println("Recievied terminate, grateful shutdown", sig)

	// greatefull shutdown
	// note: if the handlers are still working after 30s forcely close them
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
