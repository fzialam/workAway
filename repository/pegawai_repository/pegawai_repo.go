package pegawairepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
)

type PegawaiRepo interface {
	// Permohonan
	CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	GetAllUserID(ctx context.Context, tx *sql.Tx, userId int) []entity.User
	AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error)
	Set0Approved(ctx context.Context, tx *sql.Tx, suratId int) error

	// Presensi
	PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error)
	GetSurat(ctx context.Context, tx *sql.Tx, request surattugasreqres.GetSuratRequest) ([]entity.SuratTugasJOINSPPDApprovedAnggaran, error)

	GetSuratById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINUserParticipan, error)

	// Laporan
	LaporanGetAllSPPDByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApprovedLaporan, error)
	GetStatusLaporanAkBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) entity.LaporanJoinApproved
	GetStatusLaporanAngBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) entity.LaporanJoinApproved

	LaporanGetSPPDById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINUserParticipan, error)
	GetAllParticipanBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error)
	GetFotoByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, request laporanreqres.LaporanGetSPPDByIdRequest) entity.Presensi

	GetLaporanAnggaranByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) entity.LaporanJoinApproved
	GetLaporanAktivitasByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) entity.LaporanJoinApproved

	UploadLaporanAct(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error)
	UploadLaporanAngg(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error)

	SetLaporanAct(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error)
	SetLaporanAngg(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error)
}
