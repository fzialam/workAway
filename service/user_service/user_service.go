package userservice

import (
	"context"

	userreqres "github.com/fzialam/workAway/model/user_request_response"
)

type UserService interface {
	Login(ctx context.Context, request userreqres.UserLoginRequest) userreqres.UserResponse
	Register(ctx context.Context, request userreqres.UserRegisterRequest) userreqres.UserResponse
	Update(ctx context.Context, request userreqres.UserUpdateRequest) userreqres.UserResponse
	Delete(ctx context.Context, userID int)
	FindByNIP(ctx context.Context, nip string) userreqres.UserResponse
	FindByEmail(ctx context.Context, email string) userreqres.UserResponse
	FindAll(ctx context.Context) []userreqres.UserResponse
}
