package pimpinanservice

import (
	"context"

	approvedreqres "github.com/fzialam/workAway/model/req_res/approved_req_res"
	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	pimpinanreqres "github.com/fzialam/workAway/model/req_res/pimpinan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type PimpinanService interface {
	Index(ctx context.Context) (pimpinanreqres.IndexPimpinan, error)
	CreatePenugasan(ctx context.Context, request penugasanreqres.PenugasanRequest) penugasanreqres.PenugasanResponse
	GetAllUserId(ctx context.Context) []userreqres.UserResponse

	IndexPermohonan(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	PermohonanGetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINApprovedUserParticipanResponse
	PermohonanSetApproved(ctx context.Context, request izinreqres.IzinRequest) approvedreqres.ApprovedResponse

	IndexPenugasan(ctx context.Context) ([]surattugasreqres.SuratTugasJOINSPPDApprovedAnggaranResponse, error)
	SPPDGetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	SPPDGetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINSPPDApprovedAnggaranResponse
	SPPDSetApproved(ctx context.Context, request pimpinanreqres.UploadSPPDRequest) approvedreqres.ApprovedResponse

	IndexLaporan(ctx context.Context) []surattugasreqres.SuratTugasJOINUserLaporanApprovedResponse
	LaporanSPPDById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINUserFotoParticipanFotoLaporanStatusResponse
	SetApprovedLaporan(ctx context.Context, request laporanreqres.ApprovedLaporanRequest) laporanreqres.ApprovedLaporanResponse
	Profile(ctx context.Context, userId int) userreqres.UserResponse
}
