package token

import (
	"net/http"
)

const (
	Issuer = "todanni-account-service"
	KeyURL = "http://token-issuer/api/public-key"
)

type Service interface {
	// IssueToken returns a signed JWT for the provided user details
	IssueToken(w http.ResponseWriter, r *http.Request)

	// ServePublicKey serves the JWKs to use to validate a JWT
	ServePublicKey(w http.ResponseWriter, r *http.Request)
}
