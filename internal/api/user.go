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

//go:generate mockgen -source=user.go -destination=mock_restapi/mockuser.go
type UserService interface {
	Create(ctx context.Context, p api_models.CreateUser) (models.User, error)
	Delete(ctx context.Context, id string) error
	Find(id string) (models.User, error)
	Update(ctx context.Context, id string, p api_models.UpdateUser) error
}

// UserHandler ...
type UserHandler struct {
	svc UserService
}

// NewUserHandler ...
func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Register(r *mux.Router) {
	r.HandleFunc("/users", h.create).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", h.find).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", h.update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", h.delete).Methods(http.MethodDelete)
}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	var req api_models.CreateUser
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

	user, err := h.svc.Create(r.Context(), req)
	if err != nil {
		// fmt.Println(err)
		msg := fmt.Errorf("create failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w, user, http.StatusCreated)
}

func (h *UserHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	if err := h.svc.Delete(r.Context(), id); err != nil {
		msg := fmt.Errorf("delete failed: %w", err)
		renderErrorResponse(w, msg.Error(), msg)
		return
	}

	renderResponse(w, struct{}{}, http.StatusOK)
}

func (h *UserHandler) find(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	var req api_models.UpdateUser
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
