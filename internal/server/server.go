package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/token-issuer/internal/config"
	"github.com/token-issuer/pkg/token"
)

func NewTokenService(router *mux.Router, cfg config.Config) token.Service {
	service := &tokenService{
		router: router,
		config: cfg,
	}
	service.routes()
	return service
}

type tokenService struct {
	router *mux.Router
	config config.Config
}

func (s *tokenService) IssueToken(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request
	params := r.URL.Query()

	// Generate JWT and sign with JWK
	t := jwt.New()
	t.Set(jwt.ExpirationKey, time.Now().Add(time.Hour*24).Unix())
	t.Set(jwt.IssuerKey, token.Issuer)
	t.Set("uid", params["uid"][0])

	signedJWT, err := jwt.Sign(t, jwa.RS256, s.config.PrivateJWK)
	if err != nil {
		http.Error(w, "Couldn't sign JWT", http.StatusInternalServerError)
	}

	type response struct {
		Token string `json:"token"`
	}
	marshalled, _ := json.Marshal(&response{
		Token: string(signedJWT),
	})
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(marshalled)
}

func (s *tokenService) ServePublicKey(w http.ResponseWriter, r *http.Request) {
	//var keyset jwk.Set
	keyset := jwk.NewSet()
	keyset.Add(s.config.PublicJWK)

	buf, err := json.Marshal(keyset)
	if err != nil {
		http.Error(w, "Failed to marshal key", http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(buf)
}
