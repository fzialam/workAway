package pegawairepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
)

type PegawaiRepo interface {
	// Permohonan
	CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	Set0Approved(ctx context.Context, tx *sql.Tx, suratId int) error
	AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error)
	GetAllUserID(ctx context.Context, tx *sql.Tx, userId int) []entity.User

	// Presensi
	PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error)
	GetSurat(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApproved, error)

	// Laporan
	LaporanGetAllSPPDByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApprovedLaporan, error)

	LaporanGetSPPDById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINUserParticipan, error)
	GetAllParticipanBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error)
	GetFotoByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, request laporanreqres.LaporanGetSPPDByIdRequest) entity.Presensi

	GetLaporanAnggaranByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, laporan entity.LaporanAktivitasAnggaran) (entity.LaporanAktivitasAnggaran, error)
	GetLaporanAktivitasByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, laporan entity.LaporanAktivitasAnggaran) (entity.LaporanAktivitasAnggaran, error)

	UploadLaporanAct(ctx context.Context, tx *sql.Tx, laporan entity.LaporanAktivitas) (entity.LaporanAktivitas, error)
	UploadLaporanAngg(ctx context.Context, tx *sql.Tx, laporan entity.LaporanAnggaran) (entity.LaporanAnggaran, error)
}
