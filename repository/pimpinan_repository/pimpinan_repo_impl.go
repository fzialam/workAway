package pimpinanrepository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type PimpinanRepoImpl struct {
}

// GetSuratTugasByIdPermohonan implements PimpinanRepo.
func (pr *PimpinanRepoImpl) PermohonanGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINApprovedUserParticipan, error) {
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

func NewPimpinanRepo() PimpinanRepo {
	return &PimpinanRepoImpl{}
}

// CreateSurat implements PimpinanRepo.
func (pr *PimpinanRepoImpl) CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error) {
	SQL := "INSERT INTO `surat_tugas`(`tipe`,`user_id`,`lokasi_tujuan`,`jenis_program`,`dokumen_name`, `dokumen_pdf`, `dok_pendukung_name`, `dok_pendukung_pdf`,`tgl_awal`, `tgl_akhir`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.Exec(SQL,
		surat.Tipe,
		surat.UserId,
		surat.LokasiTujuan,
		surat.JenisProgram,
		surat.DokumenName,
		surat.DokumenPDF,
		surat.DokPendukungName,
		surat.DokPendukungPdf,
		surat.TglAwal,
		surat.TglAkhir,
	)
	if err != nil {
		return surat, errors.New("can't create new surat")
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		surat.Id = int(id)
		return surat, nil
	}
}

// AddParticipans implements PimpinanRepo.
func (pr *PimpinanRepoImpl) AddParticipans(ctx context.Context, tx *sql.Tx, participans entity.Participan) (entity.Participan, error) {
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

// GetAllSuratTugasJOINApprovedUserSPPD implements PimpinanRepo.
func (pr *PimpinanRepoImpl) SPPDGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status, `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.tgl_awal > NOW() AND `approved`.status = '1' AND `surat_tugas`.dokumen_name != '-';"
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

// GetAllParticipanJOINUserBySuratId implements PimpinanRepo.
func (pr *PimpinanRepoImpl) GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error) {
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

// GetAllUserID implements PimpinanRepo.
func (pr *PimpinanRepoImpl) GetAllUserId(ctx context.Context, tx *sql.Tx) []entity.User {
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

// GetSuratTugasByIdSPPD implements PimpinanRepo.
func (pr *PimpinanRepoImpl) SPPDGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugas, error) {
	SQL := "SELECT * FROM `surat_tugas` WHERE `surat_tugas`.id= ?;"
	surat := entity.SuratTugas{}
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
	)
	surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
	surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)

	if err != nil {
		return surat, err
	}
	return surat, nil
}

// SetApprovedSPPD implements PimpinanRepo.
func (pr *PimpinanRepoImpl) SPPDSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin {
	SQL := "UPDATE `approved` SET `status` = 1, `status_ttd` = ?, `status_ttd_created_at` = NOW() WHERE `surat_tugas_id` = ?;"
	_, err := tx.ExecContext(ctx, SQL, izin.StatusTTD, izin.SuratTugasId)
	helper.PanicIfError(err)
	return izin
}

// GetAllSuratTugasJOINApprovedUserPermohonan implements PimpinanRepo.
func (pr *PimpinanRepoImpl) PermohonanGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status, `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.tgl_awal > NOW();"
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

// SetApprovedPermohonan implements PimpinanRepo.
func (pr *PimpinanRepoImpl) PermohonanSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Izin) entity.Izin {
	SQL := "UPDATE `approved` SET `status` = ?, `create_at` = NOW(), `status_ttd` = ?, `status_ttd_created_at` = NOW() WHERE `surat_tugas_id` = ?;"
	_, err := tx.ExecContext(ctx, SQL, izin.Status, izin.StatusTTD, izin.SuratTugasId)
	helper.PanicIfError(err)
	return izin
}
