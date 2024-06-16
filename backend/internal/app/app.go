package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"qi-rec/internal/handler"
	"qi-rec/internal/handlergen"
	"qi-rec/internal/repository"
	"qi-rec/internal/service/recommendation/adapter"
	"qi-rec/internal/service/recommendation/recommend"
	"qi-rec/internal/service/recommendation/spotify"
	"qi-rec/internal/service/user"
	"qi-rec/pkg/config"
	"qi-rec/pkg/signal"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type App struct {
	config  *config.Config
	sigQuit chan os.Signal // signal channel for graceful shutdown
	srv     *http.Server
}

func New(cfg *config.Config) *App {
	return &App{
		config:  cfg,
		sigQuit: signal.GetShutdownChannel(),
	}
}

func (a *App) Run() {
	go func() {
		ctx := context.Background()

		recService, err := createRecommenderService(a.config)
		if err != nil {
			log.Fatalf("Failed to create recommender service: %v", err)
		}

		userService, err := setupUserRepoAndService(ctx, a.config)
		if err != nil {
			log.Fatalf("Failed to setup user service: %v", err)
		}

		h := handler.NewHandler(recService, *userService, a.config.JWTSecret)

		a.srv = setupServer(a.config, h)

		log.Println("Starting server on port", a.config.HTTPPort)
		if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Failed to start server: ", err)
		}
	}()

	<-a.sigQuit
	log.Println("Gracefully shutting down server...")

	if err := a.srv.Shutdown(context.Background()); err != nil {
		log.Fatalln("Failed to shutdown the server gracefully: ", err)
	}

	log.Println("Server shutdown is successful!")
}

func createSpotifyClient(cfg *config.Config) (*spotify.Client, error) {
	log.Println("Creating spotify client...")
	cl, err := spotify.NewClient(cfg.ClientID, cfg.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create spotify client: %w", err)
	}
	log.Println("Spotify client created successfully!")
	return cl, nil
}

func createMLAdapter(cfg *config.Config) *adapter.MLAdapter {
	return adapter.NewAdapter(cfg.MLHost, cfg.MLPort)
}

func createRecommenderService(cfg *config.Config) (recommend.Recommender, error) {
	cl, err := createSpotifyClient(cfg)
	if err != nil {
		return nil, err
	}
	ml := createMLAdapter(cfg)
	return recommend.NewRecommender(cl, ml), nil
}

func setupUserRepoAndService(ctx context.Context, cfg *config.Config) (*user.Service, error) {
	if err := processMigration(cfg); err != nil {
		return nil, fmt.Errorf("failed to process migration: %w", err)
	}

	userRepo, err := setupUserRepo(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to setup user repository: %w", err)
	}

	return user.NewService(userRepo), nil
}

func setupServer(cfg *config.Config, h *handler.Handler) *http.Server {
	r := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	corsCfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	corsCfg.ExposeHeaders = []string{"Content-Length"}
	corsCfg.AllowCredentials = true

	r.Use(cors.New(corsCfg))

	handlergen.RegisterHandlers(r, h)

	return &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("", cfg.HTTPPort),
	}
}

func setupUserRepo(ctx context.Context, cfg *config.Config) (repository.UserRepository, error) {
	log.Println("Setting up pgx pool...")
	pool, err := repository.SetupPgxPool(ctx, cfg.DbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to setup pgx pool: %w", err)
	}
	log.Println("Pgx pool is set up successfully")

	userRepo := repository.NewUserRepository(pool)

	return userRepo, nil
}

func processMigration(cfg *config.Config) error {
	log.Println("Processing migration...")
	if err := repository.ProcessMigration(cfg.MigrationPath, cfg.DbURL); err != nil {
		return fmt.Errorf("failed to process migration: %w", err)
	}
	log.Println("Migration is successful")

	return nil
}
