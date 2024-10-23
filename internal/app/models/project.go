package models

import (
	"time"

	"github.com/go-playground/validator/v10"

	"fyt/internal"
)

type ProjectType int32

const (
	Educational ProjectType = iota
	Petproject
	Startup
)

type Project struct {
	Id          int64
	Owner       int32
	ProjectType ProjectType
	Title       string
	Description string
	SocialUrl   []string
	SourceUrl   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Closed      bool
	ClosedAt    time.Time
}

func (p *Project) Validate() error {
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "invalid data format")
	}

	return nil
}
