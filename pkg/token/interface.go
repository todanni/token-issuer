package token

import (
	"net/http"
)

const (
	Issuer = "todanni-account-service"
)

type Service interface {
	// IssueToken returns a signed JWT for the provided user details
	IssueToken(w http.ResponseWriter, r *http.Request)

	// ServePublicKey serves the JWK to use to validate a JWT
	ServePublicKey(w http.ResponseWriter, r *http.Request)
}

type Repository interface {
}
