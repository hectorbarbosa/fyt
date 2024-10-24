package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"fyt/internal"
	"fyt/internal/api/api_models"
	"fyt/internal/app/models"
)

//go:generate mockgen -source=project.go -destination=mock_restapi/mockproject.go
type TaskService interface {
	Create(ctx context.Context, p api_models.CreateTask) (models.Task, error)
	Delete(ctx context.Context, id string) error
	// Search(
	// 	ctx context.Context,
	// 	name string,
	// 	description string,
	// ) ([]models.Project, error)
	Find(id string) (models.Task, error)
	Update(ctx context.Context, id string, p api_models.UpdateTask) error
	UpdateDoer(ctx context.Context, id string, p api_models.UpdateDoer) error
	UpdateDone(ctx context.Context, id string, p api_models.UpdateDone) error
}

// TaskHandler ...
type TaskHandler struct {
	svc TaskService
}

// NewTaskHandler ...
func NewTaskHandler(svc TaskService) *TaskHandler {
	return &TaskHandler{
		svc: svc,
	}
}

func (h *TaskHandler) Register(r *mux.Router) {
	r.HandleFunc("/tasks", h.create).Methods(http.MethodPost)
	// r.HandleFunc("/projects/search", h.search).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", h.find).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", h.update).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", h.delete).Methods(http.MethodDelete)
	r.HandleFunc("/tasks/{id}/doer", h.updateDoer).Methods(http.MethodPatch)
	r.HandleFunc("/tasks/{id}/done", h.updateDone).Methods(http.MethodPatch)
}

func (h *TaskHandler) create(w http.ResponseWriter, r *http.Request) {
	var req api_models.CreateTask
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e := internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder")
		msg := fmt.Errorf("invalid request %w", e)
		renderErrorResponse(w, msg.Error(), msg)

		if err = req.Validate(); err != nil {
			e := internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "params validation")
			msg := fmt.Errorf("invalid request %w", e)
			renderErrorResponse(w, msg.Error(), msg)
		}

		return
	}

	defer r.Body.Close()

	project, err := h.svc.Create(r.Context(), req)
	if err != nil {
		// fmt.Println(err)
		msg := fmt.Errorf("create failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w,
		project,
		http.StatusCreated)
}

func (h *TaskHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	if err := h.svc.Delete(r.Context(), id); err != nil {
		msg := fmt.Errorf("delete failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w, struct{}{}, http.StatusOK)
}

func (h *TaskHandler) find(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	task, err := h.svc.Find(id)
	if err != nil {
		msg := fmt.Errorf("find failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w,
		task,
		http.StatusOK)
}

func (h *TaskHandler) update(w http.ResponseWriter, r *http.Request) {
	var req api_models.UpdateTask
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e := internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder")
		msg := fmt.Errorf("invalid request %w", e)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	defer r.Body.Close()

	id := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	err := h.svc.Update(r.Context(), id, req)
	if err != nil {
		msg := fmt.Errorf("update failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w, &struct{}{}, http.StatusOK)
}

func (h *TaskHandler) updateDoer(w http.ResponseWriter, r *http.Request) {
	var req api_models.UpdateDoer
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e := internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder")
		msg := fmt.Errorf("invalid request %w", e)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	defer r.Body.Close()

	id := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	err := h.svc.UpdateDoer(r.Context(), id, req)
	if err != nil {
		msg := fmt.Errorf("update failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w, &struct{}{}, http.StatusOK)
}

func (h *TaskHandler) updateDone(w http.ResponseWriter, r *http.Request) {
	var req api_models.UpdateDone
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e := internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder")
		msg := fmt.Errorf("invalid request %w", e)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	defer r.Body.Close()

	id := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	err := h.svc.UpdateDone(r.Context(), id, req)
	if err != nil {
		msg := fmt.Errorf("update failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w, &struct{}{}, http.StatusOK)
}
