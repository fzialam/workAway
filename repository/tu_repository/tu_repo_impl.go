package turepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type TURepoImpl struct {
}

func NewTURepo() TURepo {
	return &TURepoImpl{}
}

// CreateSPPD implements TURepo.
func (tr *TURepoImpl) CreateSPPD(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error) {
	SQL := "UPDATE `surat_tugas` SET dokumen_name=?, dokumen_pdf=?, WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL,
		surat.DokumenName,
		surat.DokumenPDF,
		surat.Id,
	)

	if err != nil {
		return surat, errors.New("can't create sppd")
	} else {
		return surat, nil
	}
}

// ListSurat implements TURepo.
func (tr *TURepoImpl) ListSurat(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status_ttd ,`user`.name FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.tgl_awal > NOW() AND `approved`.status = '1';"
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
			&surat.UserName,
		)

		surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
		surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)
		surats = append(surats, surat)
	}
	if err != nil {
		return surats, errors.New("tidak ada surat tugas")
	}
	return surats, nil
}

// GetAllParticipanJOINUserBySuratId implements TURepo.
func (tr *TURepoImpl) GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error) {
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
		return participans, errors.New("tidak ada participan tugas")
	}
	return participans, nil
}

// GetSuratTugasById implements TURepo.
func (tr *TURepoImpl) GetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINApprovedUserParticipan, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status, `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.id= ?;"
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
