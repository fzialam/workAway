package pimpinanrepository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	pimpinanreqres "github.com/fzialam/workAway/model/req_res/pimpinan_req_res"
)

type PimpinanRepoImpl struct {
}

func NewPimpinanRepo() PimpinanRepo {
	return &PimpinanRepoImpl{}
}

// Index implements PimpinanRepo.
func (pr *PimpinanRepoImpl) Index(ctx context.Context, tx *sql.Tx) (pimpinanreqres.IndexPimpinan, error) {

	var index pimpinanreqres.IndexPimpinan

	SQL := "SELECT "
	SQL += "SUM(CASE WHEN `approved`.status = '0' AND `approved`.status_ttd = '0' AND `surat_tugas`.dokumen_name != '-' THEN 1 ELSE 0 END) , "
	SQL += "SUM(CASE WHEN `approved`.status = '1' AND `approved`.status_ttd = '0' AND `surat_tugas`.dokumen_name != '-' THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `approved`.status = '2' AND `approved`.status_ttd = '0' AND `surat_tugas`.dokumen_name != '-' THEN 1 ELSE 0 END) "
	SQL += "FROM `surat_tugas` "
	SQL += "LEFT JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "WHERE YEAR(tgl_awal) = YEAR(CURDATE()) AND MONTH(tgl_awal) = MONTH(CURDATE());"

	row := tx.QueryRowContext(ctx, SQL)
	err := row.Scan(&index.Permohonan.Belum, &index.Permohonan.Approved, &index.Permohonan.Reject)
	helper.PanicIfError(err)

	SQL = "SELECT "
	SQL += "SUM(CASE WHEN `approved`.status_ttd = '1' AND `approved`.status = '1' AND `surat_tugas`.dokumen_name != '-' AND `surat_tugas`.tgl_awal < NOW() THEN 1 ELSE 0 END) , "
	SQL += "SUM(CASE WHEN `approved`.status_ttd = '1' AND `approved`.status = '1' AND `surat_tugas`.dokumen_name != '-' AND `surat_tugas`.tgl_awal > NOW() THEN 1 ELSE 0 END) "
	SQL += "FROM `surat_tugas` "
	SQL += "LEFT JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "WHERE YEAR(tgl_awal) = YEAR(CURDATE());"

	row = tx.QueryRowContext(ctx, SQL)
	err = row.Scan(&index.Penugasan.Belum, &index.Penugasan.Sudah)
	helper.PanicIfError(err)

	SQL = "SELECT "
	SQL += "SUM(CASE WHEN `approved_lap_ak`.status= '0' AND `surat_tugas`.dokumen_name != '-' AND `surat_tugas`.tgl_akhir > NOW() THEN 1 ELSE 0 END) , "
	SQL += "SUM(CASE WHEN `approved_lap_ak`.status= '1' AND `surat_tugas`.dokumen_name != '-' AND `surat_tugas`.tgl_akhir > NOW() THEN 1 ELSE 0 END) , "
	SQL += "SUM(CASE WHEN `approved_lap_ak`.status= '2' AND `surat_tugas`.dokumen_name != '-' AND `surat_tugas`.tgl_akhir > NOW() THEN 1 ELSE 0 END)  "
	SQL += "FROM `surat_tugas` "
	SQL += "LEFT JOIN `laporan_aktivitas` ON `surat_tugas`.id = `laporan_aktivitas`.surat_tugas_id "
	SQL += "LEFT JOIN `approved_lap_ak` ON `laporan_aktivitas`.id = `approved_lap_ak`.laporan_id "
	SQL += "WHERE YEAR(tgl_awal) = YEAR(CURDATE()) AND MONTH(tgl_awal) = MONTH(CURDATE());"

	row = tx.QueryRowContext(ctx, SQL)
	err = row.Scan(&index.Laporan.Belum, &index.Laporan.Approved, &index.Laporan.Reject)
	helper.PanicIfError(err)

	return index, nil
}

// IndexPenugasan implements PimpinanRepo.
func (pr *PimpinanRepoImpl) IndexPenugasan(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugas, error) {
	SQL := "SELECT * "
	SQL += "FROM `surat_tugas` "
	SQL += "WHERE tgl_akhir < NOW() AND tipe =1;"
	surats := []entity.SuratTugas{}
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
	for rows.Next() {
		surat := entity.SuratTugas{}
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
	SQL := "SELECT `surat_tugas`.*, `approved`.status_ttd, `user`.nip, `user`.name, `user`.no_telp, `user`.email "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas`.tgl_awal > NOW() AND `approved`.status = '1' AND `surat_tugas`.dokumen_name != '-';"
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
			&surat.UserNIP,
			&surat.UserName,
			&surat.UserNoTelp,
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

// GetSuratTugasByIdSPPD implements PimpinanRepo.
func (pr *PimpinanRepoImpl) SPPDGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINSPPDApprovedAnggaran, error) {
	SQL := "SELECT `s`.*, `a`.status_ttd "
	SQL += "FROM `surat_tugas` `s` "
	SQL += "INNER JOIN `approved` `a` ON `s`.id = `a`.surat_tugas_id "
	SQL += "WHERE `s`.id= ?;"
	surat := entity.SuratTugasJOINSPPDApprovedAnggaran{}
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
	)
	surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
	surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)
	surat.CreateAt = helper.ConvertSQLTimeStamp(surat.CreateAt)

	if err != nil {
		return surat, err
	}
	return surat, nil
}

// SPPDSetApproved implements PimpinanRepo.
func (pr *PimpinanRepoImpl) SPPDSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Approved) entity.Approved {
	SQL := "UPDATE `approved` SET `status_ttd` = ?, `message_ttd`= ?, `status_ttd_created_at` = NOW() WHERE `surat_tugas_id` = ?;"
	_, err := tx.ExecContext(ctx, SQL, izin.StatusTTD, izin.MessageTTD, izin.SuratTugasId)
	helper.PanicIfError(err)
	return izin
}

// SPPDSetApproved implements PimpinanRepo.
func (pr *PimpinanRepoImpl) RincianSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Approved) error {
	SQL := "UPDATE `approved_rincian_anggaran` SET `status` = ?, `user_id`=2, `message`= ?, `create_at` = NOW() WHERE `rincian_id` = ?;"
	_, err := tx.ExecContext(ctx, SQL, izin.StatusTTD, izin.MessageTTD, izin.Id)
	helper.PanicIfError(err)
	return nil
}

// UploadSPPDAproved implements PimpinanRepo.
func (pr *PimpinanRepoImpl) UploadSPPDApproved(ctx context.Context, tx *sql.Tx, request pimpinanreqres.UploadSPPDRequest) error {
	SQL := "UPDATE `surat_tugas` SET `dokumen_name` = ?, `dokumen_pdf` = ? WHERE id = ?;"
	_, err := tx.ExecContext(ctx, SQL, request.DokName, request.DokPDF, request.SuratTugasId)
	helper.PanicIfError(err)
	return nil
}

// GetAllParticipanJOINUserBySuratId implements PimpinanRepo.
func (pr *PimpinanRepoImpl) GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) []entity.ParticipanJoinUser {
	SQL := "SELECT `participan`.user_id, `participan`.surat_tugas_id, `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `participan` INNER JOIN `user` ON `participan`.user_id = `user`.id WHERE `surat_tugas_id`=?"

	rows, err := tx.QueryContext(ctx, SQL, suratId)
	helper.PanicIfError(err)

	defer rows.Close()

	participans := []entity.ParticipanJoinUser{}
	for rows.Next() {
		participan := entity.ParticipanJoinUser{}
		rows.Scan(
			&participan.UserId,
			&participan.SuratTugasId,
			&participan.NIP,
			&participan.Name,
			&participan.NoTelp,
			&participan.Email,
		)
		participans = append(participans, participan)
	}
	return participans
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

// GetAllSuratTugasJOINApprovedUserPermohonan implements PimpinanRepo.
func (pr *PimpinanRepoImpl) PermohonanGetAllSuratTugasJOINApprovedUser(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINApprovedUser, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status, `user`.nip, `user`.name, `user`.no_telp, `user`.email FROM `surat_tugas` INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.tipe=0 AND `surat_tugas`.tgl_awal > NOW();"
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
			&surat.UserNIP,
			&surat.UserName,
			&surat.UserNoTelp,
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

// GetSuratTugasByIdPermohonan implements PimpinanRepo.
func (pr *PimpinanRepoImpl) PermohonanGetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINApprovedUserParticipan, error) {
	SQL := "SELECT `surat_tugas`.*, `approved`.status, `user`.nip, `user`.name, `user`.no_telp, `user`.email "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas`.id= ?;"
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
	surat.CreateAt = helper.ConvertSQLTimeStamp(surat.CreateAt)

	if err != nil {
		return surat, err
	}
	return surat, nil
}

// SetApprovedPermohonan implements PimpinanRepo.
func (pr *PimpinanRepoImpl) PermohonanSetApproved(ctx context.Context, tx *sql.Tx, izin entity.Approved) entity.Approved {
	SQL := "UPDATE `approved` SET `status` = ?, `message`=?, `create_at` = NOW(), `status_ttd` = ?, `message_ttd`=?, `status_ttd_created_at` = NOW() WHERE `surat_tugas_id` = ?;"
	_, err := tx.ExecContext(ctx, SQL, izin.Status, izin.Message, izin.StatusTTD, izin.MessageTTD, izin.SuratTugasId)
	helper.PanicIfError(err)
	return izin
}

// GetRincianBiayaBySuratId implements PimpinanRepo.
func (pr *PimpinanRepoImpl) GetRincianBiayaBySuratId(ctx context.Context, tx *sql.Tx, suratId int) entity.RincianAnggaran {
	SQL := "SELECT * FROM `rincian_anggaran` WHERE `surat_tugas_id`=?"

	var rinci entity.RincianAnggaran
	tx.QueryRowContext(ctx, SQL, suratId).Scan(
		&rinci.Id,
		&rinci.SuratTugasId,
		&rinci.DokName,
		&rinci.DokPDF,
		&rinci.CreateAt,
	)

	return rinci
}

// AddApprovedRincianBiayaById implements PimpinanRepo.
func (pr *PimpinanRepoImpl) AddApprovedRincianBiayaById(ctx context.Context, tx *sql.Tx, approved entity.Approved) error {
	SQL := "INSERT INTO `approved_rincian_anggaran`(`rincian_id`, `status`) VALUES(?,?)"

	_, err := tx.ExecContext(ctx, SQL,
		approved.Id,
		approved.StatusTTD,
	)
	helper.PanicIfError(err)

	return nil
}

// LaproanGetAllSPPD implements PimpinanRepo.
func (pr *PimpinanRepoImpl) LaporanGetAllSPPD(ctx context.Context, tx *sql.Tx) []entity.SuratTugasJOINUserLaporanApproved {
	SQL := "SELECT `surat_tugas`.*, `user`.name "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas`.tgl_akhir > NOW() AND `approved`.status_ttd = '1' AND `surat_tugas`.dokumen_name != '-';"
	surats := []entity.SuratTugasJOINUserLaporanApproved{}
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()
	for rows.Next() {
		surat := entity.SuratTugasJOINUserLaporanApproved{}
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
			&surat.UserName,
		)

		surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
		surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)
		surats = append(surats, surat)
	}

	return surats
}

// LaporanBySPPDId implements PimpinanRepo.
func (pr *PimpinanRepoImpl) LaporanBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) entity.LaporanJoinApproved {
	SQL := "SELECT `laporan_aktivitas`.* "
	SQL += "FROM `laporan_aktivitas` "
	SQL += "WHERE `laporan_aktivitas`.surat_tugas_id=?;"

	row := tx.QueryRowContext(ctx, SQL, suratId)

	var laporanAprroved entity.LaporanJoinApproved
	row.Scan(
		&laporanAprroved.Id,
		&laporanAprroved.SuratTugasId,
		&laporanAprroved.UserId,
		&laporanAprroved.DokName,
		&laporanAprroved.DokPDF,
		&laporanAprroved.CreateAt,
	)

	return laporanAprroved
}

// LaporanIsApproved implements PimpinanRepo.
func (pr *PimpinanRepoImpl) LaporanIsApproved(ctx context.Context, tx *sql.Tx, laporanId int) entity.ApprovedLaporan {
	SQL := "SELECT status "
	SQL += "FROM `approved_lap_ak` "
	SQL += "WHERE laporan_id=?;"

	row := tx.QueryRowContext(ctx, SQL, laporanId)

	var laporanAprroved entity.ApprovedLaporan
	row.Scan(
		&laporanAprroved.Status,
	)

	return laporanAprroved
}

// LaproanSPPDById implements PimpinanRepo.
func (pr *PimpinanRepoImpl) LaporanSPPDById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINUserFoto, error) {
	SQL := "SELECT `surat_tugas`.*, `user`.nip, `user`.name, `user`.no_telp, `user`.email "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas`.id= ?;"
	surat := entity.SuratTugasJOINUserFoto{}
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

// GetLaporanSPPDById implements PimpinanRepo.
func (pr *PimpinanRepoImpl) GetLaporanSPPDById(ctx context.Context, tx *sql.Tx, suratId int) entity.Laporan {
	SQL := "SELECT * FROM `laporan_aktivitas` WHERE surat_tugas_id = ? AND dok_laporan_name != '';"

	var laporan entity.Laporan
	row := tx.QueryRowContext(ctx, SQL, suratId)

	row.Scan(
		&laporan.Id,
		&laporan.SuratTugasId,
		&laporan.UserId,
		&laporan.DokLaporanName,
		&laporan.DokLaporanPDF,
		&laporan.CreateAt,
	)

	return laporan
}

// IsLaporanApproved implements PimpinanRepo.
func (pr *PimpinanRepoImpl) IsLaporanApproved(ctx context.Context, tx *sql.Tx, laporanId int) string {
	SQL := "SELECT status FROM `approved_lap_ak` WHERE laporan_id = ?;"

	var status string

	tx.QueryRowContext(ctx, SQL, laporanId).Scan(
		&status,
	)

	return status
}

// ApprovedLaporan implements PimpinanRepo.
func (pr *PimpinanRepoImpl) ApprovedLaporan(ctx context.Context, tx *sql.Tx, laporan entity.ApprovedLaporan) entity.ApprovedLaporan {
	SQL := "UPDATE `approved_lap_ak` SET user_id=?, status=?, message=?, create_at=NOW() WHERE laporan_id=?"
	_, err := tx.ExecContext(ctx, SQL, laporan.UserId, laporan.Status, laporan.Message, laporan.LaporanId)
	helper.PanicIfError(err)
	return laporan
}

// GetFotoKetuaSPPDById implements PimpinanRepo.
func (pr *PimpinanRepoImpl) GetFotoKetuaSPPDById(ctx context.Context, tx *sql.Tx, surat entity.SuratTugasJOINUserFoto) entity.SuratTugasJOINUserFoto {
	SQL := "SELECT name, gambar, lokasi, koordinat FROM `presensi` WHERE user_id=? AND surat_tugas_id=?"

	row := tx.QueryRowContext(ctx, SQL, surat.UserId, surat.Id)

	row.Scan(
		&surat.UserNameGambar,
		&surat.UserGambar,
		&surat.UserLokasi,
		&surat.UserKoordinat,
	)

	return surat
}

// GetAllFotoParticipanById implements PimpinanRepo.
func (pr *PimpinanRepoImpl) GetAllFotoParticipanById(ctx context.Context, tx *sql.Tx, participan entity.ParticipanJoinUser) entity.ParticipanJoinUserFoto {
	SQL := "SELECT * FROM `presensi` WHERE user_id=? AND surat_tugas_id=?"

	row := tx.QueryRowContext(ctx, SQL, participan.UserId, participan.SuratTugasId)

	result := entity.ParticipanJoinUserFoto{
		NIP:    participan.NIP,
		Name:   participan.Name,
		NoTelp: participan.NoTelp,
		Email:  participan.Email,
	}

	row.Scan(
		&result.Id,
		&result.UserId,
		&result.SuratTugasId,
		&result.NameGambar,
		&result.Gambar,
		&result.Lokasi,
		&result.Koordinat,
		&result.CreateAt,
	)

	return result
}

// Profile implements PimpinanRepo.
func (pr *PimpinanRepoImpl) Profile(ctx context.Context, tx *sql.Tx, userId int) entity.User {
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
