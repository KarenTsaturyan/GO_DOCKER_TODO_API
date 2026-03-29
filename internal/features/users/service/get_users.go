package users_service

import (
	"context"
	"fmt"

	"github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/domain"
	core_errors "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/errors"
)

func (s *UsersService) GetUsers(
	ctx context.Context,
	limit, offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit <= 0 {
		return nil, fmt.Errorf(
			"invalid 'limit' query param: must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"invalid 'offset' query param: must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users from repository: %w", err)
	}

	return users, nil
}
