package presensirepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type PresensiRepoImpl struct {
}

func NewPresensiRepo() PresensiRepo {
	return &PresensiRepoImpl{}
}

// CheckIzin implements PresensiRepo.
func (pr *PresensiRepoImpl) CheckIzin(ctx context.Context, tx *sql.Tx, surat_id int) error {
	SQL := "SELECT `id`, `surat_tugas_id`, `user_id`, `rank_user`, `status`, `create_at` FROM izin_pengguna WHERE `status`='disetujui' AND `rank_user`=1 AND `surat_tugas_id`=?"

	rows, err := tx.QueryContext(ctx, SQL, surat_id)
	helper.PanicIfError(err)
	defer rows.Close()

	newIzin := entity.Izin{}
	if rows.Next() {
		err := rows.Scan(&newIzin.Id, &newIzin.SuratTugasId, &newIzin.UserId, &newIzin.RankUser, &newIzin.Status, newIzin.Create_at)
		helper.PanicIfError(err)
		return nil
	} else {
		return errors.New("Surat belum disetujui")

	}
}

// PresensiFoto implements PresensiRepo.
func (pr *PresensiRepoImpl) PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error) {
	SQL := "INSERT INTO `presensi`(`user_id`, `surat_tugas_id`, `gambar`, `lokasi`) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, presensi.UserId, presensi.SuratTugasId, presensi.Gambar, presensi.Lokasi)
	if err != nil {
		return presensi, errors.New("Error Upload Gambar")
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		presensi.Id = int(id)
		return presensi, nil
	}
}
