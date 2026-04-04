package tasks_service

import (
	"context"
	"fmt"

	"github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/domain"
	core_errors "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/errors"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID, limit, offset *int,
) ([]domain.Task, error) {
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

	tasks, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}

	return tasks, nil
}
