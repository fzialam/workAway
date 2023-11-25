package permohonanrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type PermohonanRepoImpl struct {
}

func NewPermohonanRepo() PermohonanRepo {
	return &PermohonanRepoImpl{}
}

// CreatePermohonan implements PermohonanRepo.
func (pr *PermohonanRepoImpl) CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error) {
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

// AddParticipans implements PermohonanRepo.
func (pr *PermohonanRepoImpl) AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error) {
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

// Set0Approved implements PermohonanRepo.
func (pr *PermohonanRepoImpl) Set0Approved(ctx context.Context, tx *sql.Tx, suratId int) error {
	SQL := "INSERT INTO `approved`(`surat_tugas_id`, `status`, `status_ttd`) VALUES(?, 0, 0);"
	_, err := tx.ExecContext(ctx, SQL, suratId)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUserID implements PresensiRepo.
func (pr *PermohonanRepoImpl) GetAllUserID(ctx context.Context, tx *sql.Tx) []entity.User {
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
