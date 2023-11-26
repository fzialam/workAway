package pegawairepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type PegawaiRepoImpl struct {
}

func NewPegawaieRpo() PegawaiRepo {
	return &PegawaiRepoImpl{}
}

// CreatePermohonan implements PegawaiRepo.
func (pr *PegawaiRepoImpl) CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error) {
	SQL := "INSERT INTO `surat_tugas`(`tipe`, `user_id`,`lokasi_tujuan`,`jenis_program`,`dokumen_name`, `dokumen_pdf`, `dok_pendukung_name`, `dok_pendukung_pdf`,`tgl_awal`, `tgl_akhir`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.Exec(SQL, surat.Tipe, surat.LokasiTujuan, surat.LokasiTujuan, surat.JenisProgram, surat.DokumenName, surat.DokumenPDF, surat.DokPendukungName, surat.DokPendukungPdf, surat.TglAwal, surat.TglAkhir)
	if err != nil {
		return surat, errors.New("can't create new surat")
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		surat.Id = int(id)
		fmt.Println(surat.Id)
		return surat, nil
	}
}

// AddParticipans implements PegawaiRepo.
func (pr *PegawaiRepoImpl) AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error) {
	SQL := "INSERT INTO `participan`(`user_id`, `surat_tugas_id`) VALUES"

	valueStrings := []string{}
	valueArgs := []interface{}{}

	for _, data := range participans.UserId {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, data, participans.SuratTugasId)
	}

	SQL += strings.Join(valueStrings, ", ")
	result, err := tx.ExecContext(
		ctx,
		SQL,
		valueArgs...,
	)
	if err != nil {
		return participans, errors.New("can't add participan")
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		participans.Id = int(id)
		return participans, nil
	}
}

// Set0Approved implements PegawaiRepo.
func (pr *PegawaiRepoImpl) Set0Approved(ctx context.Context, tx *sql.Tx, suratId int) error {
	SQL := "INSERT INTO `approved`(`surat_tugas_id`, `status`, `status_ttd`) VALUES(?, 0, 0);"
	_, err := tx.ExecContext(ctx, SQL, suratId)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUserID implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetAllUserID(ctx context.Context, tx *sql.Tx) []entity.User {
	SQL := "SELECT `id`, `name` FROM `user` WHERE `rank`=0;"
	var users []entity.User

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.Id, &user.Name)
		helper.PanicIfError(err)

		users = append(users, user)
	}
	return users
}

// CheckIzin implements PegawaiRepo.
func (pr *PegawaiRepoImpl) CheckIzin(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) error {
	SQL := "SELECT `id`, `surat_tugas_id`, `user_id`, `status`, `CreateAt` FROM approved WHERE status=1 AND surat_tugas_id = ?"

	izin := entity.Izin{}
	tx.QueryRowContext(ctx, SQL, presensi.SuratTugasId).Scan(&izin.Id, &izin.SuratTugasId, &izin.UserId, &izin.Status, &izin.CreateAt)
	if izin.Status == "1" {
		return nil
	} else {
		return errors.New("Surat belum disetujui")
	}
}

// PresensiFoto implements PegawaiRepo.
func (pr *PegawaiRepoImpl) PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error) {
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

// GetSurat implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetSurat(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApproved, error) {
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
