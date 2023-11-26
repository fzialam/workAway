package pegawairepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/model/entity"
)

type PegawaiRepo interface {
	// Permohonan
	CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error)
	Set0Approved(ctx context.Context, tx *sql.Tx, suratId int) error
	AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error)
	GetAllUserID(ctx context.Context, tx *sql.Tx) []entity.User

	// Presensi
	CheckIzin(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) error
	PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error)
	GetSurat(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApproved, error)
}
