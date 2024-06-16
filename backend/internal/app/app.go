package app

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"

	"qi-rec/internal/handler"
	"qi-rec/internal/handlergen"
	"qi-rec/internal/service/recommendation/adapter"
	"qi-rec/internal/service/recommendation/recommend"
	"qi-rec/internal/service/recommendation/spotify"
	"qi-rec/pkg/config"
	"qi-rec/pkg/signal"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
		log.Println("Creating spotify client...")
		cl, err := spotify.NewClient(a.config.ClientID, a.config.ClientSecret)
		if err != nil {
			log.Fatalf("Failed to create spotify client: %v", err)
		}
		log.Println("Spotify client created successfully!")

		ml := adapter.NewAdapter(a.config.MLHost, a.config.MLPort)

		rec := recommend.NewRecommender(cl, ml)

		h := handler.NewHandler(rec)
		r := gin.Default()

		cfg := cors.DefaultConfig()
		cfg.AllowAllOrigins = true
		cfg.AllowMethods = []string{"GET", "POST", "OPTIONS"}
		cfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
		cfg.ExposeHeaders = []string{"Content-Length"}
		cfg.AllowCredentials = true

		r.Use(cors.New(cfg))

		handlergen.RegisterHandlers(r, h)

		a.srv = &http.Server{
			Handler: r,
			Addr:    net.JoinHostPort("", a.config.HTTPPort),
		}

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
