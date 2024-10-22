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
type ProjectService interface {
	Create(ctx context.Context, p api_models.CreateProject) (models.Project, error)
	Delete(ctx context.Context, id string) error
	// Search(
	// 	ctx context.Context,
	// 	name string,
	// 	description string,
	// ) ([]models.Project, error)
	Find(id string) (models.Project, error)
	Update(ctx context.Context, id string, p api_models.UpdateProject) error
}

// ProjectHandler ...
type ProjectHandler struct {
	svc ProjectService
}

// NewProjectHandler ...
func NewProjectHandler(svc ProjectService) *ProjectHandler {
	return &ProjectHandler{
		svc: svc,
	}
}

func (h *ProjectHandler) Register(r *mux.Router) {
	r.HandleFunc("/projects", h.create).Methods(http.MethodPost)
	// r.HandleFunc("/projects/search", h.search).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id}", h.find).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id}", h.update).Methods(http.MethodPut)
	r.HandleFunc("/projects/{id}", h.delete).Methods(http.MethodDelete)
}

func (h *ProjectHandler) create(w http.ResponseWriter, r *http.Request) {
	var req api_models.CreateProject
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

func (h *ProjectHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	if err := h.svc.Delete(r.Context(), id); err != nil {
		msg := fmt.Errorf("delete failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w, struct{}{}, http.StatusOK)
}

// ReadTasksResponse defines the response returned back after searching one task.
type ReadFilmsResponse struct {
	Project models.Project `json:"project"`
}

type SearchTasksRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// func (h *ProjectHandler) search(w http.ResponseWriter, r *http.Request) {
// 	var req SearchTasksRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		renderErrorResponse(
// 			w,
// 			"invalid request",
// 			internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder"))
// 		return
// 	}

// 	defer r.Body.Close()

// films, err := h.svc.Search(
// 	r.Context(),
// 	req.Name,
// 	req.Description,
// )
// if err != nil {
// 	msg := fmt.Errorf("search failed: %w", err)
// 	renderErrorResponse(w, msg.Error(), msg)
// 	return
// }

// 	renderResponse(w, Project{}, http.StatusOK)
// }

func (h *ProjectHandler) find(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	project, err := h.svc.Find(id)
	if err != nil {
		msg := fmt.Errorf("find failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w,
		project,
		http.StatusOK)
}

func (h *ProjectHandler) update(w http.ResponseWriter, r *http.Request) {
	var req api_models.UpdateProject
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
