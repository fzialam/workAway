package turepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	tureqres "github.com/fzialam/workAway/model/req_res/tu_req_res"
)

type TURepoImpl struct {
}

func NewTURepo() TURepo {
	return &TURepoImpl{}
}

// Index implements TURepo.
func (tr *TURepoImpl) Index(ctx context.Context, tx *sql.Tx) (tureqres.IndexTU, error) {
	SQL := "SELECT "
	SQL += "SUM(CASE WHEN `approved`.status = '1' AND `approved`.status_ttd = '0' AND `surat_tugas`.dokumen_name = '-' THEN 1 ELSE 0 END) , "
	SQL += "SUM(CASE WHEN `approved`.status = '1' AND `approved`.status_ttd = '1' AND `surat_tugas`.dokumen_name != '-' THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `approved`.status = '1' AND `approved`.status_ttd = '2' AND `surat_tugas`.dokumen_name != '-' THEN 1 ELSE 0 END) "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "WHERE YEAR(tgl_awal) = YEAR(CURDATE()) AND MONTH(tgl_awal) = MONTH(CURDATE());"

	row := tx.QueryRowContext(ctx, SQL)

	var index tureqres.IndexTU

	err := row.Scan(&index.Iscreated, &index.Approved, &index.Reject)
	helper.PanicIfError(err)

	return index, nil
}

// CreateSPPD implements TURepo.
func (tr *TURepoImpl) CreateSPPD(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error) {
	SQL := "UPDATE `surat_tugas` SET dokumen_name=?, dokumen_pdf=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL,
		surat.DokumenName,
		surat.DokumenPDF,
		surat.Id,
	)

	if err != nil {
		return surat, err
	} else {
		return surat, nil
	}
}

// SetNULLStatus implements TURepo.
func (tr *TURepoImpl) SetNULLStatus(ctx context.Context, tx *sql.Tx, suratId int) error {
	SQL := "UPDATE `approved` SET user_id=3, status_ttd=0, message_ttd=0 , status_ttd_created_at=NOW() WHERE surat_tugas_id=?"
	_, err := tx.ExecContext(ctx, SQL, suratId)
	return err
}

// ListSurat implements TURepo.
func (tr *TURepoImpl) ListSurat(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status_ttd ,`user`.name, `approved`.message_ttd "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas`.tgl_akhir > NOW() AND `approved`.status = '1' "
	SQL += "ORDER BY `surat_tugas`.create_at DESC;"
	surats := []entity.SuratTugasJOINApprovedUser{}
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
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
			&surat.UserEmail,
		)

		surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
		surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)
		surat.CreateAt = helper.ConvertSQLTimeStamp(surat.CreateAt)
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

	defer rows.Close()
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
func (tr *TURepoImpl) GetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINDoubleApprovedUserParticipan, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status_ttd, `approved`.message_ttd , `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.id= ?;"
	surat := entity.SuratTugasJOINDoubleApprovedUserParticipan{}
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
		&surat.OtherStatus,
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

// Profile implements TURepo.
func (tr *TURepoImpl) Profile(ctx context.Context, tx *sql.Tx, userId int) entity.User {
	SQL := "select * from `user` where id = ?"
	row := tx.QueryRowContext(ctx, SQL, userId)

	var user entity.User
	err := row.Scan(
		&user.Id,
		&user.NIK,
		&user.NPWP,
		&user.NIP,
		&user.Name,
		&user.Rank,
		&user.NoTelp,
		&user.TglLahir,
		&user.Status,
		&user.Gender,
		&user.Alamat,
		&user.Email,
		&user.Password,
		&user.Gambar,
	)

	user.TglLahir = helper.ConvertSQLTimeToHTML(user.TglLahir)

	helper.PanicIfError(err)
	return user
}
