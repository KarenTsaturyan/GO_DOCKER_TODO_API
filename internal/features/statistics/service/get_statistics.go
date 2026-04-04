package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/domain"
	core_errors "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"invalid date range: 'to' date must be after 'from' date: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.statisticsRepository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("failed to get tasks for statistics: %w", err)
	}

	statistics := calculateStatistics(tasks)

	return statistics, nil
}

func calculateStatistics(tasks []domain.Task) domain.Statistics {
	taskCreated := len(tasks)

	if taskCreated == 0 {
		return domain.NewStatistics(0, 0, nil, nil)
	}

	tasksCompleted := 0
	var totalCompletionDuration time.Duration
	for _, task := range tasks {
		if task.Completed {
			tasksCompleted++
		}

		completionDuration := task.CompletionDuration()
		if completionDuration != nil {
			totalCompletionDuration += *completionDuration
		}
	}

	taskCompletedRate := float64(tasksCompleted) / float64(taskCreated) * 100

	var tasksAverageCompletionTime *time.Duration
	if tasksCompleted > 0 && totalCompletionDuration != 0 {
		avgDuration := totalCompletionDuration / time.Duration(tasksCompleted)
		tasksAverageCompletionTime = &avgDuration
	}

	return domain.NewStatistics(
		taskCreated,
		tasksCompleted,
		&taskCompletedRate,
		tasksAverageCompletionTime,
	)
}
