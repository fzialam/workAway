package adminservice

import (
	"context"

	"github.com/fzialam/workAway/model/entity"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type AdminService interface {
	Permohonan(ctx context.Context) ([]surattugasreqres.SuratTugasJOINUserResponse, error)
	Penugasan(ctx context.Context) ([]surattugasreqres.SuratTugasJOINUserResponse, error)
	LapAKK(ctx context.Context) ([]entity.Laporan, error)
	LapAGG(ctx context.Context) ([]entity.Laporan, error)
	UserGET(ctx context.Context) ([]userreqres.UserResponse, error)
	UserGETById(ctx context.Context, userId int) (userreqres.UserResponse, error)
	UserPOST(ctx context.Context, request userreqres.RankChangeRequest) error

	Profile(ctx context.Context, userId int) userreqres.UserResponse
}
