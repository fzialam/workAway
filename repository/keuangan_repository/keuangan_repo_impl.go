package keuanganrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type KeuanganRepoImpl struct {
}

func NewKeuanganRepo() KeuanganRepo {
	return &KeuanganRepoImpl{}
}

// ListSurat implements KeuanganRepo.
func (kr *KeuanganRepoImpl) ListSurat(ctx context.Context, tx *sql.Tx) []entity.SuratTugasJOINApprovedUser {
	SQL := "SELECT `surat_tugas`.*, `user`.name "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas`.tgl_awal >= NOW() AND `approved`.status = '1';"
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
			&surat.UserName,
		)

		surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
		surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)

		surats = append(surats, surat)
	}

	var idS []int
	SQL = "SELECT `rincian_anggaran`.id FROM `rincian_anggaran` "
	SQL += "INNER JOIN `surat_tugas` ON `surat_tugas`.id = `rincian_anggaran`.surat_tugas_id "
	SQL += "WHERE `surat_tugas`.id=?;"
	for i := 0; i < len(surats); i++ {
		var id int
		tx.QueryRowContext(ctx, SQL, surats[i].Id).Scan(&id)

		idS = append(idS, id)
	}

	SQL = "SELECT `approved_rincian_anggaran`.status FROM `approved_rincian_anggaran`  "
	SQL += "INNER JOIN `rincian_anggaran` ON `rincian_anggaran`.id = `approved_rincian_anggaran`.rincian_id "
	SQL += "WHERE `rincian_anggaran`.id=?;"
	for i := 0; i < len(surats); i++ {
		var status string
		tx.QueryRowContext(ctx, SQL, idS[i]).Scan(&status)

		surats[i].Status = status
	}

	return surats
}

// GetSuratTugasById implements KeuanganRepo.
func (kr *KeuanganRepoImpl) GetSuratTugasById(ctx context.Context, tx *sql.Tx, suratId int) entity.SuratTugasJOINApprovedUserParticipan {
	SQL := "SELECT `surat_tugas`.*, `approved`.status_ttd , `user`.nip, `user`.name, `user`.no_telp, `user`.email "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `approved`.status ='1' AND `surat_tugas`.id= ?;"
	surat := entity.SuratTugasJOINApprovedUserParticipan{}
	row := tx.QueryRowContext(ctx, SQL, suratId)
	row.Scan(
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

	return surat
}

// GetAllParticipanJOINUserBySuratId implements KeuanganRepo.
func (*KeuanganRepoImpl) GetAllParticipanJOINUserBySuratId(ctx context.Context, tx *sql.Tx, suratId int) []entity.ParticipanJoinUser {
	SQL := "SELECT `participan`.user_id, `user`.nip, `user`.name, `user`.no_telp, `user`.email "
	SQL += "FROM `participan` "
	SQL += "INNER JOIN `user` ON `participan`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas_id`=?"

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
	return participans
}

// GetAllRincianBiayaBySuratId implements KeuanganRepo.
func (kr *KeuanganRepoImpl) GetAllRincianBiayaBySuratId(ctx context.Context, tx *sql.Tx, suratId int) entity.Laporan {
	SQL := "SELECT * FROM `rincian_anggaran` WHERE `surat_tugas_id`=?"

	var laporan entity.Laporan
	tx.QueryRowContext(ctx, SQL, suratId).Scan(
		&laporan.Id,
		&laporan.SuratTugasId,
		&laporan.DokLaporanName,
		&laporan.DokLaporanPDF,
		&laporan.CreateAt,
	)

	return laporan
}

// UploadRincianAnggaran implements KeuanganRepo.
func (kr *KeuanganRepoImpl) UploadRincianAnggaran(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) entity.RincianAnggaran {
	SQL := "INSERT INTO `rincian_anggaran`(`surat_tugas_id`, `dok_name`,`dok_pdf`) VALUES(?,?,?)"

	row, err := tx.ExecContext(ctx, SQL,
		rinci.SuratTugasId,
		rinci.DokName,
		rinci.DokPDF,
	)
	helper.PanicIfError(err)

	id, _ := row.LastInsertId()

	rinci.Id = int(id)
	return rinci
}

// AddNULLApprovedRincian implements KeuanganRepo.
func (kr *KeuanganRepoImpl) AddNULLApprovedRincian(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) error {
	SQL := "INSERT INTO `approved_rincian_anggaran`(`rincian_id`, `status`, `message`) VALUES(?,'0','0')"

	_, err := tx.ExecContext(ctx, SQL,
		rinci.Id,
	)
	helper.PanicIfError(err)

	return nil
}

// SetRincianAnggaran implements KeuanganRepo.
func (kr *KeuanganRepoImpl) SetRincianAnggaran(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) entity.RincianAnggaran {
	SQL := "UPDATE `rincian_anggaran` SET `dok_name`=?, `dok_pdf`=?, `create_at`=NOW() WHERE `surat_tugas_id`=?;"
	_, err := tx.ExecContext(ctx, SQL, rinci.DokName, rinci.DokPDF, rinci.SuratTugasId)
	helper.PanicIfError(err)

	return rinci
}

// GetIDRincianAnggaran implements KeuanganRepo.
func (*KeuanganRepoImpl) GetIDRincianAnggaran(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) entity.RincianAnggaran {
	SQL := "SELECT * FROM `rincian_anggaran` WHERE `surat_tugas_id`=?;"
	err := tx.QueryRowContext(ctx, SQL, rinci.SuratTugasId).Scan(
		&rinci.Id,
		&rinci.SuratTugasId,
		&rinci.DokName,
		&rinci.DokPDF,
		&rinci.CreateAt,
	)
	helper.PanicIfError(err)

	return rinci
}

// SetNullRincianAnggaran implements KeuanganRepo.
func (kr *KeuanganRepoImpl) SetNullApprovedRincianAnggaran(ctx context.Context, tx *sql.Tx, rinci entity.RincianAnggaran) error {
	SQL := "UPDATE `approved_rincian_anggaran` SET `status`='0', `message`='0', `create_at`=NOW() WHERE `id`=?;"
	_, err := tx.ExecContext(ctx, SQL, rinci.Id)
	helper.PanicIfError(err)

	return nil
}

// ListSPPDIsApproved implements KeuanganRepo.
func (kr *KeuanganRepoImpl) ListSPPDIsApproved(ctx context.Context, tx *sql.Tx) []entity.SuratTugasJOINApprovedUserOtherId {
	SQL := "SELECT `surat_tugas`.*, `approved_lap_ak`.status, `user`.name, `rincian_anggaran`.id "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "INNER JOIN `laporan_aktivitas` ON `laporan_aktivitas`.surat_tugas_id = `surat_tugas`.id "
	SQL += "INNER JOIN `approved_lap_ak` ON `approved_lap_ak`.laporan_id = `laporan_aktivitas`.id "
	SQL += "INNER JOIN `rincian_anggaran` ON `rincian_anggaran`.surat_tugas_id = `surat_tugas`.id "
	SQL += "WHERE `surat_tugas`.tgl_akhir >= NOW() AND `approved`.status_ttd = '1';"
	surats := []entity.SuratTugasJOINApprovedUserOtherId{}
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	for rows.Next() {
		surat := entity.SuratTugasJOINApprovedUserOtherId{}
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
			&surat.OtherId,
		)

		surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
		surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)
		surats = append(surats, surat)
	}

	SQL = "SELECT status FROM `full_anggaran` WHERE rincian_id=?;"
	for i := 0; i < len(surats); i++ {
		tx.QueryRowContext(ctx, SQL, surats[i].OtherId).Scan(
			&surats[i].OtherStatus,
		)
	}
	return surats
}

// SetFullAnggaran implements KeuanganRepo.
func (kr *KeuanganRepoImpl) SetFullAnggaran(ctx context.Context, tx *sql.Tx, approved entity.Approved) entity.Approved {
	SQL := "UPDATE `full_anggaran` SET status=1, create_at=NOW() WHERE rincian_id=?"
	_, err := tx.ExecContext(ctx, SQL,
		approved.Id,
	)
	helper.PanicIfError(err)

	return approved
}

// ListLaporanSPPD implements KeuanganRepo.
func (kr *KeuanganRepoImpl) ListLaporanSPPD(ctx context.Context, tx *sql.Tx) []entity.SuratTugasJOINUserLaporanApproved {
	SQL := "SELECT `surat_tugas`.*, `user`.name "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas`.tgl_akhir >= NOW() AND "
	SQL += "`approved`.status_ttd = '1' AND `surat_tugas`.dokumen_name != '-';"
	surats := []entity.SuratTugasJOINUserLaporanApproved{}
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
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
		surat.CreateAt = helper.ConvertSQLTimeStamp(surat.CreateAt)
		surats = append(surats, surat)
	}
	return surats
}

// LaporanBySPPDId implements KeuanganRepo.
func (kr *KeuanganRepoImpl) LaporanBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) entity.LaporanJoinApproved {
	SQL := "SELECT `laporan_anggaran`.* "
	SQL += "FROM `laporan_anggaran` "
	SQL += "WHERE `laporan_anggaran`.surat_tugas_id=?;"

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

// IsLaporanApproved implements KeuanganRepo.
func (kr *KeuanganRepoImpl) IsLaporanApproved(ctx context.Context, tx *sql.Tx, laporanId int) entity.ApprovedLaporan {
	SQL := "SELECT `approved_lap_angg`.status "
	SQL += "FROM `approved_lap_angg` "
	SQL += "WHERE `approved_lap_angg`.laporan_id=?;"

	row := tx.QueryRowContext(ctx, SQL, laporanId)

	var laporanAprroved entity.ApprovedLaporan
	row.Scan(
		&laporanAprroved.Status,
	)

	return laporanAprroved
}

// SetApprovedLaporan implements KeuanganRepo.
func (kr *KeuanganRepoImpl) SetApprovedLaporan(ctx context.Context, tx *sql.Tx, laporan entity.ApprovedLaporan) entity.ApprovedLaporan {
	SQL := "UPDATE `approved_lap_angg` SET `user_id`=4, `status`=?, `message`=?, `create_at` = NOW() WHERE `laporan_id`=?"
	_, err := tx.ExecContext(ctx, SQL, laporan.Status, laporan.Message, laporan.LaporanId)
	helper.PanicIfError(err)
	return laporan
}
