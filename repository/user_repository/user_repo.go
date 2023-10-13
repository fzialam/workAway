package userrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type UserRepo interface {
	Login(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	Register(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	CheckEmailNIP(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	Delete(ctx context.Context, tx *sql.Tx, user entity.User)
	FindByNIP(ctx context.Context, tx *sql.Tx, nip string) (entity.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.User
}