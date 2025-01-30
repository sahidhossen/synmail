package middleware

import "github.com/sahidhossen/synmail/src/config"

type MiddlewareHandler struct {
	*config.Config
}

func CreateMiddleware(cfg *config.Config) *MiddlewareHandler {
	return &MiddlewareHandler{Config: cfg}
}
