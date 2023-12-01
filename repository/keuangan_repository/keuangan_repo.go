package keuanganrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type KeuanganRepo interface {
	// List Permohonan
	ListSurat(ctx context.Context, tx *sql.Tx) []entity.SuratTugasJOINApprovedUser

	// Detail Permohonan
	GetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) entity.SuratTugasJOINApprovedUserParticipan
	GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) []entity.ParticipanJoinUser

	GetAllRincianBiayaBySuratId(ctx context.Context, tx *sql.Tx, suratId int) entity.Laporan

	// Upload Rincian Anggaran Permohonan u/ Approved pimpinan
	UploadRincianAnggaran(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) entity.RincianAnggaran
	AddNULLApprovedRincian(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) error

	SetRincianAnggaran(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) entity.RincianAnggaran
	GetIDRincianAnggaran(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) entity.RincianAnggaran
	SetNullApprovedRincianAnggaran(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) error

	// =======================================

	// List SPPD u/ Full Pay anggaran
	ListSPPDIsApproved(ctx context.Context, tx *sql.Tx) []entity.SuratTugasJOINApprovedUserOtherId

	// SetFullAnggaran
	SetFullAnggaran(ctx context.Context, tx *sql.Tx, approved entity.Approved) entity.Approved

	// =======================================

	// List Laporan SPPD
	ListLaporanSPPD(ctx context.Context, tx *sql.Tx) []entity.SuratTugasJOINUserLaporanApproved

	// Detail Laporan SPPD
	LaporanBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) entity.LaporanJoinApproved
	IsLaporanApproved(ctx context.Context, tx *sql.Tx, laporanId int) entity.ApprovedLaporan

	SetApprovedLaporan(ctx context.Context, tx *sql.Tx, laporan entity.ApprovedLaporan) entity.ApprovedLaporan

	// =======================================
}
