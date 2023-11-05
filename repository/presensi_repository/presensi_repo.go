package presensirepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type PresensiRepo interface {
	CheckIzin(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) error
	GetSurat(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugas, error)
	PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error)
}
