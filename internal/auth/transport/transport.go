package transport

import "github.com/RickDred/ozinse/internal/auth"

type transport struct {
	service auth.Service
}

func New(s auth.Service) auth.Handlers {
	return &transport{service: s}
}
