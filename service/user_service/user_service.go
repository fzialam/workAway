package userservice

import (
	"context"

	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type UserService interface {
	Login(ctx context.Context, request userreqres.UserLoginRequest) (userreqres.LoginResponse, error)
	Register(ctx context.Context, request userreqres.UserRegisterRequest) userreqres.UserResponse
	Delete(ctx context.Context, userID int)
	FindByNIP(ctx context.Context, nip string) userreqres.UserResponse
	FindByEmail(ctx context.Context, email string) userreqres.UserResponse
	FindAll(ctx context.Context) []userreqres.UserResponse

	Profile(ctx context.Context, userId int) userreqres.UserResponse
	UpdateProfile(ctx context.Context, request userreqres.UserUpdateRequest) userreqres.UserResponse
	ChangePassword(ctx context.Context, request userreqres.ChangePasswordReq) userreqres.UserResponse
	ChangeImage(ctx context.Context, request userreqres.ChangeImageReq) userreqres.UserResponse
}
