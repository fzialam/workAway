package userservice

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/middleware"
	"github.com/fzialam/workAway/model/entity"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
	userrepository "github.com/fzialam/workAway/repository/user_repository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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
func (us *UserServiceImpl) Login(ctx context.Context, request userreqres.UserLoginRequest) (userreqres.LoginResponse, error) {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Email: request.Email,
	}

	user, err = us.UserRepo.Login(ctx, tx, user)
	if err != nil {
		return helper.ToLoginResponse(user, ""), err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return helper.ToLoginResponse(user, ""), errors.New("email atau password salah")
	} else {

		token := middleware.LoginSetJWT(user)

		return helper.ToLoginResponse(user, token), nil
	}
}

// Register implements UserService.
func (us *UserServiceImpl) Register(ctx context.Context, request userreqres.UserRegisterRequest) userreqres.UserResponse {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	user := entity.User{
		NIP:      request.NIP,
		NIK:      request.NIK,
		NPWP:     request.NPWP,
		Name:     request.Name,
		Rank:     0,
		NoTelp:   request.NoTelp,
		TglLahir: request.TglLahir,
		Status:   "1",
		Gender:   request.Gender,
		Alamat:   request.Alamat,
		Email:    request.Email,
		Password: string(hashPassword),
		Gambar:   "data:image/png;base64,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",
	}

	go func() {

		// Chechk Duplicate Email
		user, err = us.UserRepo.CheckEmail(ctx, tx, user)
		if err != nil {
			panic(exception.NewDuplicatedError(err.Error()))
		}

	}()

	go func() {

		// Chechk Duplicate NIP
		user, err = us.UserRepo.CheckNIP(ctx, tx, user)
		if err != nil {
			panic(exception.NewDuplicatedError(err.Error()))
		}
	}()

	go func() {

		// Chechk Duplicate NIK
		user, err = us.UserRepo.CheckNIK(ctx, tx, user)
		if err != nil {
			panic(exception.NewDuplicatedError(err.Error()))
		}
	}()

	go func() {

		// Chechk Duplicate NPWP
		user, err = us.UserRepo.CheckNPWP(ctx, tx, user)
		if err != nil {
			panic(exception.NewDuplicatedError(err.Error()))
		}
	}()

	user, err = us.UserRepo.Register(ctx, tx, user)
	helper.PanicIfError(err)

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

// Profil implements UserService.
func (us *UserServiceImpl) Profile(ctx context.Context, userId int) userreqres.UserResponse {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := us.UserRepo.Profile(ctx, tx, userId)

	return helper.ToUserResponse(user)
}

// UpdateProfile implements UserService.
func (us *UserServiceImpl) UpdateProfile(ctx context.Context, request userreqres.UserUpdateRequest) userreqres.UserResponse {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Id:       request.Id,
		NIP:      request.NIP,
		NIK:      request.NIK,
		NPWP:     request.NPWP,
		Name:     request.Name,
		NoTelp:   request.NoTelp,
		TglLahir: request.TglLahir,
		Alamat:   request.Alamat,
		Email:    request.Email,
	}

	user, err = us.UserRepo.UpdateProfile(ctx, tx, user)

	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

// ChangePassword implements UserService.
func (us *UserServiceImpl) ChangePassword(ctx context.Context, request userreqres.ChangePasswordReq) userreqres.UserResponse {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := us.UserRepo.Profile(ctx, tx, request.Id)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicIfError(err)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user.Password = string(hashPassword)

	user, err = us.UserRepo.ChangePassword(ctx, tx, user)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

// ChangeImage implements UserService.
func (us *UserServiceImpl) ChangeImage(ctx context.Context, request userreqres.ChangeImageReq) userreqres.UserResponse {
	err := us.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Id:     request.Id,
		Gambar: request.Gambar,
	}
	user, err = us.UserRepo.ChangeImage(ctx, tx, user)

	helper.PanicIfError(err)

	return helper.ToUserResponse(user)

}
