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
	SQL := "SELECT `id`, `surat_tugas_id`, `user_id`, `status`, `CreateAt` FROM approved WHERE status=1 AND surat_tugas_id = ?"

	izin := entity.Izin{}
	tx.QueryRowContext(ctx, SQL, presensi.SuratTugasId).Scan(&izin.Id, &izin.SuratTugasId, &izin.UserId, &izin.Status, &izin.CreateAt)
	if izin.Status == "1" {
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
func (pr *PresensiRepoImpl) GetSurat(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApproved, error) {
	SQL := "SELECT surat_tugas.*, approved.status, approved.status_ttd, approved.status_ttd_created_at FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id WHERE `surat_tugas`.tgl_akhir > NOW() AND `surat_tugas`.user_id = ? AND (approved.status_ttd_created_at = '0' OR approved.status_ttd_created_at = '1');"
	surats := []entity.SuratTugasJOINApproved{}
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	for rows.Next() {
		surat := entity.SuratTugasJOINApproved{}
		rows.Scan(
			&surat.Id,
			&surat.Tipe,
			&surat.UserId,
			&surat.LokasiTujuan,
			&surat.JenisProgram,
			&surat.DokumenName,
			&surat.DokumenPDF,
			&surat.DokPendukungName,
			&surat.DokPendukungPdf,
			&surat.TglAwal,
			&surat.TglAkhir,
			&surat.CreateAt,
			&surat.Status,
			&surat.StatusTTD,
			&surat.StatusTTDCreateAt,
		)
		surats = append(surats, surat)
	}
	if err != nil {
		return surats, errors.New("Tidak ada surat tugas")
	}
	return surats, nil
}
