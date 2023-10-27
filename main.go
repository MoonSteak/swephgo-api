package main

import (
	"context"
	"github.com/MoonSteak/swephgo-api/internal/core"
	"github.com/MoonSteak/swephgo-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RequestDurationLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)

		// Calculate the duration and log it.
		duration := time.Since(startTime)
		log.Printf("URL: %s, Duration: %s", r.URL.Path, duration)
	})
}

func main() {
	core.LibInfo()

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	r := chi.NewRouter()

	r.Use(RequestDurationLoggerMiddleware)
	r.Get("/bodies-degree", handlers.BodiesDegreeHandler)
	r.Get("/jdn-by-offset", handlers.JdnByOffsetHandler)

	server := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Second,
		Addr:         ":3000",
		Handler:      r,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatalf("HttpServer: run error")
		}
	}()

	<-killSignal
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()
	log.Info("HttpServer: shutdown")
	if err := server.Shutdown(ctx); err != nil {
		log.WithError(err).Fatalf("HttpServer: shutdown failed")
	}
	log.Info("HttpServer: exited properly")
}
