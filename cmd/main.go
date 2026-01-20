package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KhoalaS/guitar-girl-offline/pkg/core/auth"
	"github.com/KhoalaS/guitar-girl-offline/pkg/core/cdn"
	"github.com/KhoalaS/guitar-girl-offline/pkg/core/game"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

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

	userDatabase := game.NewMockDb()
	gameService := &game.GameServiceImpl{
		UserRepository:            game.NewUserRepositoryImpl(userDatabase),
		UserAreaRepository:        game.NewUserAreaRepositoryImpl(userDatabase),
		UserAchievementRepository: game.NewUserAchievementRepository(userDatabase),
		UserCharacterRepository:   game.NewUserCharacterRepository(userDatabase),
		UserCostumeRepository:     game.NewUserCostumeRepository(userDatabase),
		Timezone:                  "Asia/Seoul",
	}
	gameApi := game.NewGameApi(gameService)
	gameRpc := game.NewGameRpc(gameApi)

	gameMux := http.NewServeMux()
	gameRpc.RegisterRoutes(gameMux)
	gameServer := http.Server{
		Addr:    ":10001",
		Handler: gameMux,
	}

	go func() {
		if err := gameServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Listen error")
		}
	}()

	cdnServer := cdn.NewCdnServer(10002)
	go func() {
		if err := cdnServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
