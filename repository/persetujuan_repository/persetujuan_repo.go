package persetujuanrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type PersetujuanRepo interface {
	SetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin
	GetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	GetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINApprovedUserParticipan, error)
	GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error)
}
