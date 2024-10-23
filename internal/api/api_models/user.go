package api_models

import (
	"fyt/internal"

	"github.com/go-playground/validator/v10"
)

type CreateUser struct {
	Email    string `json:"email" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (p *CreateUser) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "CreateUser")
	}

	return nil
}

type UpdateUser struct {
	Email    string `json:"email" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password"`
}

func (p *UpdateUser) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "UpdateUser")
	}

	return nil
}
