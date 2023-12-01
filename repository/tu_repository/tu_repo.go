package turepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type TURepo interface {
	CreateSPPD(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	SetNULLStatus(ctx context.Context, tx *sql.Tx, suratId int) error
	ListSurat(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error)
	GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error)
	GetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINDoubleApprovedUserParticipan, error)
}
