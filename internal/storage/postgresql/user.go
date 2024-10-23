package postgresql

import (
	"context"
	"errors"
	"strings"

	"fyt/internal"
	"fyt/internal/api/api_models"
	"fyt/internal/app/models"
	"fyt/internal/storage/postgresql/db"

	"github.com/jackc/pgx/v5"
)

// UserRepository represents the repository used for interacting with user records.
type UserRepository struct {
	q *db.Queries
}

// NewFilm instantiates the Project repository.
func NewUserRepo(d db.DBTX) *UserRepository {
	r := UserRepository{
		q: db.New(d),
	}

	return &r
}

// Create inserts a new record in db.
func (r *UserRepository) Create(u api_models.CreateUser) (models.User, error) {
	row, err := r.q.InsertUser(context.Background(), db.InsertUserParams{
		Email:    u.Email,
		UserName: u.UserName,
		Password: u.Password,
	})
	if err != nil {
		return models.User{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert User")
	}

	return models.User{
		Id:        row.ID,
		Email:     u.Email,
		UserName:  u.UserName,
		Password:  u.Password,
		CreatedAt: row.CreatedAt.Time,
	}, nil
}

// Delete deletes the existing record matching the id.
func (r *UserRepository) Delete(id int32) error {
	result, err := r.q.DeleteUser(context.Background(), id)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "delete user")
	}

	if result == 0 {
		return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "delete user")
	}

	return nil
}

func (r *UserRepository) Find(id int32) (models.User, error) {
	user, err := r.q.SelectUser(context.Background(), id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "user not found")
		}

		return models.User{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "select user")
	}

	return models.User{
		Id:        user.ID,
		Email:     user.Email,
		UserName:  user.UserName,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Time,
	}, nil
}

func (r *UserRepository) Update(id int32, u api_models.UpdateUser) error {
	result, err := r.q.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       id,
		Email:    u.Email,
		UserName: u.UserName,
		Password: u.Password,
	})
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return internal.WrapErrorf(err, internal.ErrorCodeUniqueConstraints, "update user")
		} else {
			return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "update user")
		}
	}

	if result == 0 {
		return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "update user")
	}

	return nil
}
