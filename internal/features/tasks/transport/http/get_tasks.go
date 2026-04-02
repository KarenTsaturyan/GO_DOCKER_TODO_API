package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/logger"
	core_http_request "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/transport/http/request"
	core_http_response "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, limit, offset, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'user_id', 'limit' or 'offset' query parameters",
		)
		return
	}

	tasksDomain, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get tasks",
		)
		return
	}

	response := GetTasksResponse(taskDTOsFromDomains(tasksDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		userIDQueryParam = "user_id"
		limitQueryParam  = "limit"
		offsetQueryParam = "offset"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query params: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query params: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query params: %w", err)
	}

	return userID, limit, offset, nil
}
