package adminrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type AdminRepo interface {
	Permohonan(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINUser, error)
	Penugasan(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINUser, error)
	LapAKK(ctx context.Context, tx *sql.Tx) ([]entity.Laporan, error)
	LapAGG(ctx context.Context, tx *sql.Tx) ([]entity.Laporan, error)
	UserGETById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error)
	UserGET(ctx context.Context, tx *sql.Tx) ([]entity.User, error)
	UserPOST(ctx context.Context, tx *sql.Tx, user entity.User) error

	Profile(ctx context.Context, tx *sql.Tx, userId int) entity.User
}
