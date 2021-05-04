package server

import "net/http"

func (s *tokenService) routes() {
	// Issues a signed JWT
	s.router.HandleFunc("/api/token", s.IssueToken).Methods(http.MethodGet)

	// Server the public JWKs
	s.router.HandleFunc("/api/token/public-key", s.ServePublicKey).Methods(http.MethodGet)
}
