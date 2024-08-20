package main

import (
	"context"
	"fmt"
	"go-starter-webapp/config"
	"go-starter-webapp/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	  if err != nil {
	    log.Fatal("Error loading .env file")
	  }
	var wait time.Duration

	config.ConnectDB()
	r := routes.Getroutes()
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}
	fmt.Println("Server started at port "+os.Getenv("PORT"))
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
			os.Exit(0)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	config.CloseDB()
	os.Exit(0)
}
