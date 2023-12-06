package keuanganservice

import (
	"context"

	approvedreqres "github.com/fzialam/workAway/model/req_res/approved_req_res"
	dokumenreqres "github.com/fzialam/workAway/model/req_res/dokumen_req_res"
	keuanganreqres "github.com/fzialam/workAway/model/req_res/keuangan_req_res"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type KeuanganService interface {
	Index(ctx context.Context) (keuanganreqres.IndexKeuangan, error)
	ListPermohonanApproved(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	PermohonanApprovedById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINApprovedUserParticipanLaporanResponse
	UploadRincianBiaya(ctx context.Context, request keuanganreqres.UploadRincianAnggaranRequest) dokumenreqres.UploadDokumenResponse
	SetRincian(ctx context.Context, request keuanganreqres.UploadRincianAnggaranRequest) dokumenreqres.UploadDokumenResponse

	ListSPPDApproved(ctx context.Context) []keuanganreqres.SuratTugasJOINApprovedUserOtherIdResponse
	SetFullAnggaran(ctx context.Context, request approvedreqres.ApprovedRequest) approvedreqres.ApprovedResponse

	ListLaporanSPPD(ctx context.Context) []surattugasreqres.SuratTugasJOINUserLaporanApprovedResponse
	SetApprovedLaporan(ctx context.Context, request laporanreqres.ApprovedLaporanRequest) laporanreqres.ApprovedLaporanResponse
	Profile(ctx context.Context, userId int) userreqres.UserResponse
}
