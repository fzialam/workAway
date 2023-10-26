package helper

import (
	"github.com/fzialam/workAway/model/entity"
	presensireqres "github.com/fzialam/workAway/model/presensi_request_response"
	userreqres "github.com/fzialam/workAway/model/user_request_response"
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

func ToPresensiResponse(presensi entity.Presensi) presensireqres.PresensiFotoResponse {
	return presensireqres.PresensiFotoResponse{
		Id:           presensi.Id,
		UserId:       presensi.UserId,
		SuratTugasId: presensi.SuratTugasId,
		Gambar:       "Sukses Upload Gambar",
		Lokasi:       presensi.Lokasi,
	}
}
