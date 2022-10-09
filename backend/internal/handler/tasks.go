package handler

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CreateTaskRequest struct {
	Task *domain.Task `json:"task"`
}

type CreateTaskResponse struct {
	TaskId int `json:"taskId"`
}

type GetAllTasksResponse struct {
	Tasks []*domain.Task `json:"tasks"`
}

// @Summary Create Task
// @Security ApiKeyAuth
// @Tags tasks
// @Description create task
// @Accept json
// @Produce json
// @Param task body CreateTaskRequest true "task"
// @Success 200 {object} CreateTaskResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/tasks [post]
func (h *Handler) CreateTask(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	req := new(CreateTaskRequest)

	if err := c.Bind(req); err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	taskId, err := h.service.CreateTask(ctx, userId, req.Task)
	if err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"taskId": taskId,
	})
}

// @Summary Get All Tasks
// @Tags tasks
// @Description get all tasks
// @Accept json
// @Produce json
// @Success 200 {object} GetAllTasksResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/tasks [get]
func (h *Handler) GetAllTasks(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	tasks, err := h.service.GetAllTasks(ctx)
	if err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}

// @Summary Delete Task
// @Security ApiKeyAuth
// @Tags tasks
// @Description delete task
// @Param task_id   path int true "Task Id"
// @Accept json
// @Produce json
// @Success 204
// @Failure 401 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/cards/{task_id} [delete]
func (h *Handler) DeleteTask(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	taskIdString := c.Param("task_id")
	taskId, err := strconv.Atoi(taskIdString)
	if err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	if err = h.service.DeleteTask(ctx, userId, taskId); err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
