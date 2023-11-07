package permohonanrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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
	log.Println("LOG CREATE SURAT")
	fmt.Println(len(surat.DokPendukungPdf))
	SQL := "INSERT INTO `surat_tugas`(`lokasi_tujuan`,`jenis_program`,`dokumen_name`, `dokumen_pdf`, `dok_pendukung_name`, `dok_pendukung_pdf`,`tgl_awal`, `tgl_akhir`) VALUES (?,?,?,?,?,?,?,?)"
	result, err := tx.Exec(SQL, surat.LokasiTujuan, surat.JenisProgram, surat.DokumenName, surat.DokumenPDF, surat.DokPendukungName, surat.DokPendukungPdf, surat.TglAwal, surat.TglAkhir)
	if err != nil {
		log.Println(err)
		return surat, errors.New("can't create new surat")
	} else {
		id, err := result.LastInsertId()
		log.Println(err)
		helper.PanicIfError(err)

		surat.Id = int(id)
		fmt.Println(surat.Id)
		return surat, nil
	}
}

// WhoCreate implements PermohonanRepo.
func (*PermohonanRepoImpl) AddPemohon(ctx context.Context, tx *sql.Tx, pemohon entity.Pemohon) (entity.Pemohon, error) {
	log.Println("LOG ADD PEMOHON")
	SQL := "INSERT INTO `pemohon`(`user_id`, `surat_tugas_id`) VALUES (?,?);"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		pemohon.UserId,
		pemohon.SuratTugasId,
	)
	if err != nil {
		return pemohon, err
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		pemohon.Id = int(id)
		return pemohon, nil
	}
}

// AddParticipans implements PermohonanRepo.
func (pr *PermohonanRepoImpl) AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error) {
	log.Println("LOG ADD PARTICIPANS")
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
