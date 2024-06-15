package main

import (
	"log"
	"net/http"

	"qi-rec/internal/handler"
	"qi-rec/internal/handlergen"
	"qi-rec/internal/service/recommendation/spotify"
	"qi-rec/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	cl, err := spotify.NewClient(cfg.ClientID, cfg.ClientSecret)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	h := handler.NewHandler(cl)
	r := gin.Default()

	handlergen.RegisterHandlers(r, h)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(srv.ListenAndServe())
}
