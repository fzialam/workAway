package pimpinanservice

import (
	"context"

	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	pimpinanreqres "github.com/fzialam/workAway/model/req_res/pimpinan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type PimpinanService interface {
	CreatePenugasan(ctx context.Context, request penugasanreqres.PenugasanRequest) penugasanreqres.PenugasanResponse
	GetAllUserId(ctx context.Context) []userreqres.UserResponse

	PermohonanSetApproved(ctx context.Context, request izinreqres.IzinRequest) izinreqres.IzinResponse
	PermohonanGetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	PermohonanGetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINApprovedUserParticipanResponse

	SPPDSetApproved(ctx context.Context, request pimpinanreqres.UploadSPPDRequest) izinreqres.IzinResponse
	SPPDGetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	SPPDGetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasResponse
}
