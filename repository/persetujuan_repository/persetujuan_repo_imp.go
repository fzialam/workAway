package persetujuanrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type PersetujuanRepoImpl struct {
}

func NewPersetujuanRepo() PersetujuanRepo {
	return &PersetujuanRepoImpl{}
}

// SetApproved implements PersetujuanRepo.
func (pr *PersetujuanRepoImpl) SetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin {
	SQL := "UPDATE `approved` SET `status` = ?, `create_at` = NOW(), `status_ttd` = ?, `status_ttd_created_at` = NOW() WHERE `surat_tugas_id` = ?;"
	_, err := tx.ExecContext(ctx, SQL, izin.Status, izin.StatusTTD, izin.SuratTugasId)
	helper.PanicIfError(err)
	return izin
}

// GetAllParticipanBySuratId implements PersetujuanRepo.
func (pr *PersetujuanRepoImpl) GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error) {
	SQL := "SELECT `participan`.user_id, `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `participan` INNER JOIN `user` ON `participan`.user_id = `user`.id WHERE `surat_tugas_id`=?"
	participans := []entity.ParticipanJoinUser{}
	rows, err := tx.QueryContext(ctx, SQL, suratId)
	helper.PanicIfError(err)
	for rows.Next() {
		participan := entity.ParticipanJoinUser{}
		rows.Scan(
			&participan.UserId,
			&participan.NIP,
			&participan.Name,
			&participan.NoTelp,
			&participan.Email,
		)
		participans = append(participans, participan)
	}
	if err != nil {
		return participans, errors.New("Tidak ada participan tugas")
	}
	return participans, nil

}

// GetAllSuratTugas implements PersetujuanRepo.
func (pr *PersetujuanRepoImpl) GetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status, `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.tgl_awal > NOW() AND;"
	surats := []entity.SuratTugasJOINApprovedUser{}
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	for rows.Next() {
		surat := entity.SuratTugasJOINApprovedUser{}
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
			&surat.UserNIP,
			&surat.UserName,
			&surat.UserNoTelp,
			&surat.UserEmail,
		)

		surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
		surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)
		surats = append(surats, surat)
	}
	if err != nil {
		return surats, errors.New("Tidak ada surat tugas")
	}
	return surats, nil
}

// GetSuratTugasById implements PersetujuanRepo.
func (pr *PersetujuanRepoImpl) GetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINApprovedUserParticipan, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status, `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.id= ?;"
	// log.Fatal(len(SQL))
	surat := entity.SuratTugasJOINApprovedUserParticipan{}
	row := tx.QueryRowContext(ctx, SQL, suratId)
	err := row.Scan(
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
		&surat.UserNIP,
		&surat.UserName,
		&surat.UserNoTelp,
		&surat.UserEmail,
	)
	surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
	surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)

	if err != nil {
		return surat, err
	}
	return surat, nil
}
