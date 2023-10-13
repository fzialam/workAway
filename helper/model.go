package helper

import (
	"github.com/fzialam/workAway/model/entity"
	userreqres "github.com/fzialam/workAway/model/web/user_request_response"
)

func ToUserResponse(user entity.User) userreqres.UserResponse {
	return userreqres.UserResponse{
		Id:    user.Id,
		NIP:   user.NIP,
		Email: user.Email,
	}
}

func ToUserResponses(categories []entity.User) []userreqres.UserResponse {
	var userResponses []userreqres.UserResponse
	for _, user := range categories {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
