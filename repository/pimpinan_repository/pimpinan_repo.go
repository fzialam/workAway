package pimpinanrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type PimpinanRepo interface {
	// Penugasan
	CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error)
	GetAllUserId(ctx context.Context, tx *sql.Tx) []entity.User
	GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error)

	PermohonanGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINApprovedUserParticipan, error)
	PermohonanGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	PermohonanSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin

	SPPDGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	SPPDSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin
	SPPDGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugas, error)
}
