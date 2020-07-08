package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MrWitold/Gvote/handlers"
	"github.com/MrWitold/Gvote/models"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "vote-api ", log.LstdFlags)

	// create the handlers
	ph := handlers.NewItems(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/vote/{id}", ph.ShowAll)
	getR.HandleFunc("/vote/{id}/{token}", ph.ShowSingle)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/vote/{id}/{token}", ph.UpdateVote)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/vote", ph.CreateVote)

	// create a new server
	s := http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// initila database
	models.InitialMigration()

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		l.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
	s.Shutdown(ctx)
}
