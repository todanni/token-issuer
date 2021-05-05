package server

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	uid, err := strconv.Atoi(params["uid"][0])
	if err != nil {
		http.Error(w, "Couldn't convert user ID to int", http.StatusBadRequest)
	}

	// Generate JWT and sign with JWK
	t := jwt.New()
	t.Set(jwt.ExpirationKey, time.Now().Add(time.Hour*24).Unix())
	t.Set(jwt.IssuerKey, token.Issuer)
	t.Set("uid", uid)

	signedJWT, err := jwt.Sign(t, jwa.RS256, s.config.PrivateJWK)
	if err != nil {
		http.Error(w, "Couldn't sign JWT", http.StatusInternalServerError)
	}
	_, err = w.Write(signedJWT)
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
