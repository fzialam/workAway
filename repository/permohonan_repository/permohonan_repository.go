package permohonanrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type PermohonanRepo interface {
	CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	AddPemohon(ctx context.Context, tx *sql.Tx, pemohon entity.Pemohon) (entity.Pemohon, error)
	AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error)
	GetAllUserID(ctx context.Context, tx *sql.Tx) []entity.User
}
