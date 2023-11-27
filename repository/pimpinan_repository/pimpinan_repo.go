package pimpinanrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
	pimpinanreqres "github.com/fzialam/workAway/model/req_res/pimpinan_req_res"
)

type PimpinanRepo interface {
	// Penugasan
	GetAllUserId(ctx context.Context, tx *sql.Tx) []entity.User
	CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error)
	GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error)

	// Permohonan
	PermohonanGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	PermohonanGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINApprovedUserParticipan, error)
	PermohonanSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin

	// SPPD
	SPPDGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	SPPDGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugas, error)
	SPPDSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin

	SetNullLaporanAktivitas(ctx context.Context, tx *sql.Tx, suratId int) error
	SetNullLaporanAnggaran(ctx context.Context, tx *sql.Tx, suratId int) error

	SetNullApprovedAktivitas(ctx context.Context, tx *sql.Tx, suratId int) error
	SetNullApprovedAnggaran(ctx context.Context, tx *sql.Tx, suratId int) error

	UploadSPPDAproved(ctx context.Context, tx *sql.Tx, request pimpinanreqres.UploadSPPDRequest) error
}
