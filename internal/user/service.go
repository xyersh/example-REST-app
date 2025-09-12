package user

import (
	"context"
	"log/slog"
)

type Service struct {
	storage Storage
	logger  *slog.Logger
}

func (s *Service) Create(ctx context.Context, dto *CreateUserDTO) (u User, err error) {
	// TODO

	return
}
