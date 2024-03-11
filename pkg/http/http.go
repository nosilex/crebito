package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(hdl http.Handler, addr string) {
	// The HTTP Server
	srv := &http.Server{
		Addr:    addr,
		Handler: hdl,
	}

	// Server run context
	srvCtx, srvStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		log.Println("shutting down server...")

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(srvCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		srvStopCtx()
	}()

	log.Printf("Runing at: %s", addr)

	// Run the srv
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}

	// Wait for server context to be stopped
	<-srvCtx.Done()

	log.Println("server exiting")
}
