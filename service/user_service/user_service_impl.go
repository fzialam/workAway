package userservice

import (
	"context"
	"database/sql"

	userrequestresponse "github.com/fzialam/workAway/model/web/user_request_response"
	userrepository "github.com/fzialam/workAway/repository/user_repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepo userrepository.UserRepo
	DB       *sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepo userrepository.UserRepo, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		DB:       DB,
		Validate: validate,
	}
}

// Login implements UserService.
func (*UserServiceImpl) Login(ctx context.Context, request userrequestresponse.UserLoginRequest) userrequestresponse.UserResponse {
	panic("unimplemented")
}

// Register implements UserService.
func (*UserServiceImpl) Register(ctx context.Context, request userrequestresponse.UserRegisterRequest) userrequestresponse.UserResponse {
	panic("unimplemented")
}

// Update implements UserService.
func (*UserServiceImpl) Update(ctx context.Context, request userrequestresponse.UserUpdateRequest) userrequestresponse.UserResponse {
	panic("unimplemented")
}

// Delete implements UserService.
func (*UserServiceImpl) Delete(ctx context.Context, userID int) {
	panic("unimplemented")
}

// FindAll implements UserService.
func (*UserServiceImpl) FindAll(ctx context.Context) []userrequestresponse.UserResponse {
	panic("unimplemented")
}

// FindByEmail implements UserService.
func (*UserServiceImpl) FindByEmail(ctx context.Context, email string) userrequestresponse.UserResponse {
	panic("unimplemented")
}

// FindByNIP implements UserService.
func (*UserServiceImpl) FindByNIP(ctx context.Context, nip string) userrequestresponse.UserResponse {
	panic("unimplemented")
}
