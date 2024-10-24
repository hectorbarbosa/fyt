package api_models

import (
	"fyt/internal"
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateTask struct {
	ProjectId   int64  `json:"project_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	DueDate     string `json:"due_date" validate:"required"`
}

func (t *CreateTask) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "CreateTask")
	}

	if _, err := time.Parse("2006-01-02", t.DueDate); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "CreateTask")
	}

	return nil
}

type UpdateTask struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	DueDate     string `json:"due_date" validate:"required"`
}

func (t *UpdateTask) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "UpdateTask")
	}

	if _, err := time.Parse("2006-01-02", t.DueDate); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "UpdateTask")
	}

	return nil
}

type UpdateDoer struct {
	Doer int32 `json:"doer" validate:"required,number"`
}

func (t *UpdateDoer) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "UpdateDoer")
	}

	return nil
}

type UpdateDone struct {
	Done bool `json:"done" validate:"required"`
}

func (t *UpdateDone) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "UpdateDone")
	}

	return nil
}
