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
func (pr *PresensiRepoImpl) CheckIzin(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) error {
	SQL := "SELECT `id`, `surat_tugas_id`, `user_id`, `status`, `create_at` FROM approved WHERE status=1 AND surat_tugas_id = ?"

	newIzin := entity.Izin{}
	tx.QueryRowContext(ctx, SQL, presensi.SuratTugasId).Scan(&newIzin.Id, &newIzin.SuratTugasId, &newIzin.UserId, &newIzin.Status, &newIzin.Create_at)
	if newIzin.Status == 1 {
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

// GetSurat implements PresensiRepo.
func (pr *PresensiRepoImpl) GetSurat(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugas, error) {
	SQL := "SELECT surat_tugas.*, approved.status FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.id WHERE `surat_tugas`.tgl_akhir > NOW() AND `surat_tugas`.user_id = ?;"
	surats := []entity.SuratTugas{}
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	for rows.Next() {
		surat := entity.SuratTugas{}
		rows.Scan(
			&surat.Id,
			&surat.UserId,
			&surat.JudulSurat,
			&surat.TglAwal,
			&surat.TglAkhir,
			&surat.Create_at,
			&surat.Status,
		)
		surats = append(surats, surat)

	}
	if err != nil {
		return surats, errors.New("Tidak ada surat tugas")
	}

	return surats, nil

}
