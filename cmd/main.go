package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KhoalaS/guitar-girl-offline/pkg/core/auth"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	authMux := http.NewServeMux()
	authService := auth.NewAuthService()

	authApi := auth.NewAuthApi(authService)
	authApi.RegisterStaticRoutes(authMux)
	authApi.RegisterAccountsRoutes(authMux)

	authServer := http.Server{
		Addr:    ":10000",
		Handler: authMux,
	}
	go func() {

		if err := authServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Listen error")
		}
	}()

	<-ctx.Done()

	// give server 5 seconds to wrap up
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := authServer.Shutdown(shutdownCtx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server shut down gracefully")
}
