package penugasanservice

import (
	"context"

	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type PenugasanService interface {
	SetApproved(ctx context.Context, request izinreqres.IzinRequest) izinreqres.IzinResponse
	CreatePenugasan(ctx context.Context, request penugasanreqres.PenugasanRequest) penugasanreqres.PenugasanResponse
	GetAllUserId(ctx context.Context) []userreqres.UserResponse
	GetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	GetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasResponse
}
