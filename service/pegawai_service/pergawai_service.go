package pegawaiservice

import (
	"context"

	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	presensireqres "github.com/fzialam/workAway/model/req_res/presensi_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type PegawaiService interface {
	// Create Permohonan
	CreatePermohonan(ctx context.Context, request permohonanreqres.PermohonanRequest) permohonanreqres.PermohonanResponse
	GetAllUserId(ctx context.Context) []userreqres.UserResponse

	// Mobile
	PresensiFoto(ctx context.Context, request presensireqres.PresensiFotoRequest) presensireqres.PresensiFotoResponse
	GetSurat(ctx context.Context, request presensireqres.GetSuratForPresensiRequest) []surattugasreqres.SuratTugasJOINApprovedResponse

	// Laporan
	LaporanGetAllSPPDByUserId(ctx context.Context, userId int) []surattugasreqres.SuratTugasJOINApprovedLaporanResponse
	LaporanGetSPPDById(ctx context.Context, request laporanreqres.LaporanGetSPPDByIdRequest) surattugasreqres.SuratTugasJOINUserParticipanLaporanResponse

	UploadLapAktivitas(ctx context.Context, request laporanreqres.UploadLaporanRequest) laporanreqres.UploadLaporanResponse
	UploadLapAnggaran(ctx context.Context, request laporanreqres.UploadLaporanRequest) laporanreqres.UploadLaporanResponse
}
