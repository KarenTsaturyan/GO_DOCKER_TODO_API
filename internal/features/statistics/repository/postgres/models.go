package statistics_postgres_repository

import (
	"time"

	"github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/domain"
)

// Code Duplication: This code is duplicated from task/
type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func taskDomainsFromModels(taskModels []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(taskModels))

	for i, taskModel := range taskModels {
		domains[i] = taskDomainFromModel(taskModel)
	}

	return domains
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)
}
