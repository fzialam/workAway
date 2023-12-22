package adminrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
	adminreqres "github.com/fzialam/workAway/model/req_res/admin_req_res"
)

type AdminRepo interface {
	Index(ctx context.Context, tx *sql.Tx) (adminreqres.IndexResponse, error)
	Permohonan(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINUser, error)
	Penugasan(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINUser, error)
	LapAKK(ctx context.Context, tx *sql.Tx) ([]entity.Laporan, error)
	LapAGG(ctx context.Context, tx *sql.Tx) ([]entity.Laporan, error)
	UserGETById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error)
	UserGET(ctx context.Context, tx *sql.Tx) ([]entity.User, error)
	UserPOST(ctx context.Context, tx *sql.Tx, user entity.User) error

	Profile(ctx context.Context, tx *sql.Tx, userId int) entity.User
}
