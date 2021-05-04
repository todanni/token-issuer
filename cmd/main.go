package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/token-issuer/internal/config"
	"github.com/token-issuer/internal/server"
)

func main() {
	// Read config
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Initialise router
	r := mux.NewRouter()

	// Create servers by passing DB connection and router
	server.NewTokenService(r, cfg)

	// Start the servers and listen
	log.Fatal(http.ListenAndServe(":8083", r))
}
