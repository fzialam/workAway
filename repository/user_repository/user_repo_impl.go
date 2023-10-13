package userrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type UserRepoImpl struct {
}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

// Login implements UserRepo.
func (*UserRepoImpl) Login(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	panic("unimplemented")
}

// Save implements UserRepo.
func (*UserRepoImpl) Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	panic("unimplemented")
}

// Update implements UserRepo.
func (*UserRepoImpl) Update(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	panic("unimplemented")
}

// Delete implements UserRepo.
func (*UserRepoImpl) Delete(ctx context.Context, tx *sql.Tx, user entity.User) {
	panic("unimplemented")
}

// FindAll implements UserRepo.
func (*UserRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.User, error) {
	panic("unimplemented")
}

// FindByEmail implements UserRepo.
func (*UserRepoImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error) {
	panic("unimplemented")
}

// FindByNIP implements UserRepo.
func (*UserRepoImpl) FindByNIP(ctx context.Context, tx *sql.Tx, nip string) (entity.User, error) {
	panic("unimplemented")
}
