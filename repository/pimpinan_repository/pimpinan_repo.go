package pimpinanrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
	pimpinanreqres "github.com/fzialam/workAway/model/req_res/pimpinan_req_res"
)

type PimpinanRepo interface {
	Index(ctx context.Context, tx *sql.Tx) (pimpinanreqres.IndexPimpinan, error)
	IndexPenugasan(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugas, error)

	// Penugasan
	GetAllUserId(ctx context.Context, tx *sql.Tx) []entity.User
	CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error)
	GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) []entity.ParticipanJoinUser

	// Permohonan
	PermohonanGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	PermohonanGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINApprovedUserParticipan, error)
	PermohonanSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Approved) entity.Approved

	// SPPD
	SPPDGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	SPPDGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINSPPDApprovedAnggaran, error)
	SPPDSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Approved) entity.Approved

	RincianSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Approved) error

	UploadSPPDApproved(ctx context.Context, tx *sql.Tx, request pimpinanreqres.UploadSPPDRequest) error
	GetRincianBiayaBySuratId(ctx context.Context, tx *sql.Tx, suratId int) entity.RincianAnggaran
	AddApprovedRincianBiayaById(ctx context.Context, tx *sql.Tx, approved entity.Approved) error

	// Laporan
	LaporanGetAllSPPD(ctx context.Context, tx *sql.Tx) []entity.SuratTugasJOINUserLaporanApproved
	LaporanIsApproved(ctx context.Context, tx *sql.Tx, laporanId int) entity.ApprovedLaporan

	LaporanBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) entity.LaporanJoinApproved
	LaporanSPPDById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINUserFoto, error)
	GetLaporanSPPDById(ctx context.Context, tx *sql.Tx, suratId int) entity.Laporan
	IsLaporanApproved(ctx context.Context, tx *sql.Tx, laporanId int) string

	GetFotoKetuaSPPDById(ctx context.Context, tx *sql.Tx, surat entity.SuratTugasJOINUserFoto) entity.SuratTugasJOINUserFoto
	GetAllFotoParticipanById(ctx context.Context, tx *sql.Tx, participan entity.ParticipanJoinUser) entity.ParticipanJoinUserFoto
	ApprovedLaporan(ctx context.Context, tx *sql.Tx, laporan entity.ApprovedLaporan) entity.ApprovedLaporan

	Profile(ctx context.Context, tx *sql.Tx, userId int) entity.User
}
