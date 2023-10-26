package presensirepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type PresensiRepo interface {
	CheckIzin(ctx context.Context, tx *sql.Tx, surat_id int) error
	PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error)
}