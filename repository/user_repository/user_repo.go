package userrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type UserRepo interface {
	Login(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	Register(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	CheckEmail(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	CheckNIP(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	CheckNIK(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	CheckNPWP(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	Delete(ctx context.Context, tx *sql.Tx, user entity.User)
	FindByNIP(ctx context.Context, tx *sql.Tx, nip string) (entity.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.User

	Profile(ctx context.Context, tx *sql.Tx, userId int) entity.User
	UpdateProfile(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	ChangePassword(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	ChangeImage(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
}
