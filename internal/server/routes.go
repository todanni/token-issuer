package server

import "net/http"

func (s *tokenService) routes() {
	// GET an template
	s.router.HandleFunc("/api/token", s.IssueToken).Methods(http.MethodGet)

	s.router.HandleFunc("/api/token/public-key", s.ServePublicKey).Methods(http.MethodGet)
}
