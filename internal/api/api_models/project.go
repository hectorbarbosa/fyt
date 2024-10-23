package api_models

import (
	"fyt/internal"
	"fyt/internal/app/models"

	"github.com/go-playground/validator/v10"
)

type CreateProject struct {
	Owner       int32    `json:"owner" validate:"required"`
	ProjectType int32    `json:"project_type" validate:"required"`
	Title       string   `json:"title" validate:"required,min=2,max=150"`
	Description string   `json:"description" validate:"required,min=5,max=500"`
	SocialUrl   []string `json:"social_url" validate:"required"`
	SourceUrl   string   `json:"source_url"`
}

func (p *CreateProject) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "CreateProject")
	}

	if p.ProjectType != int32(models.Educational) && p.ProjectType != int32(models.Petproject) && p.ProjectType != int32(models.Startup) {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, "invalid project type")
	}

	return nil
}

type UpdateProject struct {
	ProjectType int32    `json:"project_type" validate:"required"`
	Title       string   `json:"title" validate:"required,min=2,max=150"`
	Description string   `json:"description" validate:"required,min=12,max=500"`
	SocialUrl   []string `json:"social_url" validate:"required"`
	SourceUrl   string   `json:"source_url"`
	Closed      bool     `json:"closed"`
}

func (p *UpdateProject) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "UpdateProject")
	}

	if p.ProjectType != int32(models.Educational) && p.ProjectType != int32(models.Petproject) && p.ProjectType != int32(models.Startup) {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, "Invalid project type")
	}

	return nil
}
