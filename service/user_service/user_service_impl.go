package userservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
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
func (us *UserServiceImpl) Login(ctx context.Context, request userreqres.UserLoginRequest) userreqres.UserResponse {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err = us.UserRepo.Login(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

// Register implements UserService.
func (us *UserServiceImpl) Register(ctx context.Context, request userreqres.UserRegisterRequest) userreqres.UserResponse {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		NIP:      request.NIP,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err = us.UserRepo.CheckEmailNIP(ctx, tx, user)
	if err != nil {
		panic(exception.NewDuplicatedError(err.Error()))
	}
	user, err = us.UserRepo.Register(ctx, tx, user)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

// Update implements UserService.
func (us *UserServiceImpl) Update(ctx context.Context, request userreqres.UserUpdateRequest) userreqres.UserResponse {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := us.UserRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user = entity.User{
		NIP:      request.NIP,
		Email:    request.Email,
		Password: request.Password,
	}

	user = us.UserRepo.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

// Delete implements UserService.
func (us *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := us.UserRepo.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	us.UserRepo.Delete(ctx, tx, user)
}

// FindAll implements UserService.
func (us *UserServiceImpl) FindAll(ctx context.Context) []userreqres.UserResponse {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := us.UserRepo.FindAll(ctx, tx)
	return helper.ToUserResponses(user)
}

// FindByEmail implements UserService.
func (us *UserServiceImpl) FindByEmail(ctx context.Context, email string) userreqres.UserResponse {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := us.UserRepo.FindByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

// FindByNIP implements UserService.
func (us *UserServiceImpl) FindByNIP(ctx context.Context, nip string) userreqres.UserResponse {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := us.UserRepo.FindByNIP(ctx, tx, nip)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}
