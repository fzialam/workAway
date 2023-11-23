package penugasanrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type PenugasanRepo interface {
	CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error)
	GetAllUserID(ctx context.Context, tx *sql.Tx) []entity.User
	SetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin
	GetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	GetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugas, error)
}
