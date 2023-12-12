package pegawairepository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	pegawaireqres "github.com/fzialam/workAway/model/req_res/pegawai_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
)

type PegawaiRepoImpl struct {
}

func NewPegawaiRepo() PegawaiRepo {
	return &PegawaiRepoImpl{}
}

// Index implements PegawaiRepo.
func (pr *PegawaiRepoImpl) Index(ctx context.Context, tx *sql.Tx, userId int) (pegawaireqres.IndexPegawai, error) {
	var index pegawaireqres.IndexPegawai

	SQL := "SELECT "
	SQL += "SUM(CASE WHEN `sub_quer`.s_0 != 0 THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `sub_quer`.s_1 != 0 THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `sub_quer`.s_2 != 0 THEN 1 ELSE 0 END) "
	SQL += "FROM ( SELECT DISTINCT "
	SQL += "(CASE WHEN `approved`.status = '0' AND `approved`.status_ttd = '0' THEN `s`.id ELSE 0 END) AS `s_0` ,  "
	SQL += "(CASE WHEN `approved`.status = '1' AND `approved`.status_ttd = '0' THEN `s`.id ELSE 0 END) AS `s_1`, "
	SQL += "(CASE WHEN `approved`.status = '2' AND `approved`.status_ttd = '0' THEN `s`.id ELSE 0 END) AS `s_2` "
	SQL += "FROM `surat_tugas` `s`"
	SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
	SQL += "LEFT JOIN `user` `u` on `u`.id =`s`.user_id "
	SQL += "LEFT JOIN `approved` ON `s`.id = `approved`.surat_tugas_id "
	SQL += "WHERE YEAR(tgl_awal) = YEAR(CURDATE()) AND MONTH(tgl_awal) = MONTH(CURDATE()) AND "
	SQL += "(`s`.user_id = ? OR `p`.user_id = ?) AND `s`.tipe=0 "
	SQL += ") AS sub_quer;"

	row := tx.QueryRowContext(ctx, SQL, userId, userId)
	err := row.Scan(&index.Permohonan.Belum, &index.Permohonan.Approved, &index.Permohonan.Reject)
	helper.PanicIfError(err)

	SQL = "SELECT "
	SQL += "SUM(CASE WHEN `sub_quer`.s_0 != 0 THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `sub_quer`.s_1 != 0 THEN 1 ELSE 0 END) "
	SQL += "FROM ( SELECT DISTINCT "
	SQL += "(CASE WHEN `approved`.status_ttd = '1'AND `s`.tgl_awal < NOW() THEN `s`.id ELSE 0 END) AS `s_0`, "
	SQL += "(CASE WHEN `approved`.status_ttd = '1'AND `s`.tgl_awal > NOW() THEN `s`.id ELSE 0 END) AS `s_1` "
	SQL += "FROM `surat_tugas` `s` "
	SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
	SQL += "LEFT JOIN `user` `u` on `u`.id =`s`.user_id "
	SQL += "LEFT JOIN `approved` ON `s`.id = `approved`.surat_tugas_id "
	SQL += "WHERE YEAR(tgl_awal) = YEAR(CURDATE()) AND MONTH(tgl_awal) = MONTH(CURDATE()) AND "
	SQL += "(`s`.user_id = ? OR `p`.user_id = ?) AND `s`.tipe=1 "
	SQL += ") AS sub_quer;"

	row = tx.QueryRowContext(ctx, SQL, userId, userId)
	err = row.Scan(&index.Penugasan.Belum, &index.Penugasan.Sudah)
	helper.PanicIfError(err)

	SQL = "SELECT "
	SQL += "SUM(CASE WHEN `sub_quer`.s_0 != 0 THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `sub_quer`.s_1 != 0 THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `sub_quer`.s_2 != 0 THEN 1 ELSE 0 END) "
	SQL += "FROM ( SELECT DISTINCT "
	SQL += "(CASE WHEN `approved_lap_ak`.status= '0' AND `s`.dokumen_name != '-' AND `s`.tgl_akhir > NOW() THEN `s`.id ELSE 0 END) AS `s_0` , "
	SQL += "(CASE WHEN `approved_lap_ak`.status= '1' AND `s`.dokumen_name != '-' AND `s`.tgl_akhir > NOW() THEN `s`.id ELSE 0 END) AS `s_1` , "
	SQL += "(CASE WHEN `approved_lap_ak`.status= '2' AND `s`.dokumen_name != '-' AND `s`.tgl_akhir > NOW() THEN `s`.id ELSE 0 END) AS `s_2`  "
	SQL += "FROM `surat_tugas` `s` "
	SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
	SQL += "LEFT JOIN `laporan_aktivitas` ON `s`.id = `laporan_aktivitas`.surat_tugas_id "
	SQL += "LEFT JOIN `approved_lap_ak` ON `laporan_aktivitas`.id = `approved_lap_ak`.laporan_id "
	SQL += "WHERE YEAR(tgl_awal) = YEAR(CURDATE()) AND MONTH(tgl_awal) = MONTH(CURDATE()) AND "
	SQL += "(`s`.user_id = ? OR `p`.user_id = ?)"
	SQL += ") AS sub_quer; "

	row = tx.QueryRowContext(ctx, SQL, userId, userId)
	err = row.Scan(&index.Aktivitas.Belum, &index.Aktivitas.Approved, &index.Aktivitas.Reject)
	helper.PanicIfError(err)

	SQL = "SELECT "
	SQL += "SUM(CASE WHEN `sub_quer`.s_0 != 0 THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `sub_quer`.s_1 != 0 THEN 1 ELSE 0 END), "
	SQL += "SUM(CASE WHEN `sub_quer`.s_2 != 0 THEN 1 ELSE 0 END) "
	SQL += "FROM ( SELECT DISTINCT "
	SQL += "(CASE WHEN `approved_lap_angg`.status= '0' AND `s`.dokumen_name != '-' AND `s`.tgl_akhir > NOW() THEN `s`.id ELSE 0 END) AS `s_0` , "
	SQL += "(CASE WHEN `approved_lap_angg`.status= '1' AND `s`.dokumen_name != '-' AND `s`.tgl_akhir > NOW() THEN `s`.id ELSE 0 END) AS `s_1` , "
	SQL += "(CASE WHEN `approved_lap_angg`.status= '2' AND `s`.dokumen_name != '-' AND `s`.tgl_akhir > NOW() THEN `s`.id ELSE 0 END) AS `s_2`  "
	SQL += "FROM `surat_tugas` `s` "
	SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
	SQL += "LEFT JOIN `laporan_anggaran` ON `s`.id = `laporan_anggaran`.surat_tugas_id "
	SQL += "LEFT JOIN `approved_lap_angg` ON `laporan_anggaran`.id = `approved_lap_angg`.laporan_id "
	SQL += "WHERE YEAR(tgl_awal) = YEAR(CURDATE()) AND MONTH(tgl_awal) = MONTH(CURDATE()) AND "
	SQL += "(`s`.user_id = ? OR `p`.user_id = ?)"
	SQL += ") AS sub_quer; "

	row = tx.QueryRowContext(ctx, SQL, userId, userId)
	err = row.Scan(&index.Anggaran.Belum, &index.Anggaran.Approved, &index.Anggaran.Reject)
	helper.PanicIfError(err)

	return index, nil
}

// CreatePermohonan implements PegawaiRepo.
func (pr *PegawaiRepoImpl) CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error) {
	SQL := "INSERT INTO `surat_tugas`(`tipe`, `user_id`,`lokasi_tujuan`,`jenis_program`,`dokumen_name`, `dokumen_pdf`, `dok_pendukung_name`, `dok_pendukung_pdf`,`tgl_awal`, `tgl_akhir`) "
	SQL += "VALUES (?,?,?,?,'-','-',?,?,?,?)"
	result, err := tx.Exec(SQL,
		surat.Tipe,
		surat.UserId,
		surat.LokasiTujuan,
		surat.JenisProgram,
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

// GetAllUserID implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetAllUserID(ctx context.Context, tx *sql.Tx, userId int) []entity.User {
	SQL := "SELECT `id`, `name` FROM `user` WHERE `rank`=0 AND `id` != ?;"
	var users []entity.User

	rows, err := tx.QueryContext(ctx, SQL, userId)
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
	SQL := "INSERT INTO `approved`(`surat_tugas_id`, `status`, `message`, `status_ttd`, `message_ttd`) VALUES(?, 0, 0, 0, 0);"
	_, err := tx.ExecContext(ctx, SQL, suratId)
	if err != nil {
		return err
	}
	return nil
}

// GetSuratById implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetSuratById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINUserParticipan, error) {
	SQL := "SELECT `surat_tugas`.*, `user`.nip, `user`.name, `user`.no_telp, `user`.email "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas`.id=?;"
	surat := entity.SuratTugasJOINUserParticipan{}
	err := tx.QueryRowContext(ctx, SQL, suratId).Scan(
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
	surat.CreateAt = helper.ConvertSQLTimeStamp(surat.CreateAt)
	if err != nil {
		return surat, errors.New("tidak ada surat tugas")
	}
	return surat, nil
}

// GetAllParticipanBySPPDId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetAllParticipanBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) ([]entity.ParticipanJoinUser, error) {
	SQL := "SELECT `participan`.user_id, `user`.nip, `user`.name, `user`.no_telp, `user`.email "
	SQL += "FROM `participan` "
	SQL += "INNER JOIN `user` ON `participan`.user_id = `user`.id "
	SQL += "WHERE `surat_tugas_id`=?"
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

// PresensiFoto implements PegawaiRepo.
func (pr *PegawaiRepoImpl) PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error) {
	SQL := "INSERT INTO `presensi`(`user_id`, `surat_tugas_id`, `gambar`, `lokasi`, `koordinat`) VALUES (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, presensi.UserId, presensi.SuratTugasId, presensi.Gambar, presensi.Lokasi, presensi.Koordinat)
	if err != nil {
		return presensi, errors.New("error upload gambar")
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		presensi.Id = int(id)
		return presensi, nil
	}
}

// GetSurat implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetSurat(ctx context.Context, tx *sql.Tx, request surattugasreqres.GetSuratRequest) ([]entity.SuratTugasJOINSPPDApprovedAnggaran, error) {
	if request.Tipe == "permohonan" {
		SQL := "SELECT `surat_tugas`.*, `approved`.status "
		SQL += "FROM `surat_tugas` "
		SQL += "INNER JOIN `approved` ON `approved`.surat_tugas_id = `surat_tugas`.id "
		SQL += "WHERE `surat_tugas`.tgl_awal >= NOW() AND `surat_tugas`.tipe=0 AND "
		SQL += "`surat_tugas`.user_id = ?;"
		surats := []entity.SuratTugasJOINSPPDApprovedAnggaran{}
		rows, err := tx.QueryContext(ctx, SQL, request.UserId)
		helper.PanicIfError(err)

		defer rows.Close()
		for rows.Next() {
			surat := entity.SuratTugasJOINSPPDApprovedAnggaran{}
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
			)
			surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
			surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)

			surats = append(surats, surat)
		}
		if err != nil {
			return surats, errors.New("tidak ada surat tugas")
		}
		return surats, nil
	} else {
		SQL := "SELECT `s`.* "
		SQL += "FROM `surat_tugas` `s` "
		SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
		SQL += "LEFT JOIN `user` `u` on `u`.id =`s`.user_id "
		SQL += "LEFT JOIN `approved` `a` on `s`.id = `a`.surat_tugas_id "
		SQL += "WHERE (`s`.user_id = ? OR `p`.user_id = ?) AND `a`.status_ttd = '1' AND `s`.tgl_akhir >= NOW();"
		surats := []entity.SuratTugasJOINSPPDApprovedAnggaran{}
		rows, err := tx.QueryContext(ctx, SQL, request.UserId, request.UserId)
		helper.PanicIfError(err)

		defer rows.Close()

		for rows.Next() {
			surat := entity.SuratTugasJOINSPPDApprovedAnggaran{}
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

			surats = append(surats, surat)
		}
		if err != nil {
			return surats, errors.New("tidak ada surat tugas")
		}
		SQL = "SELECT * FROM `rincian_anggaran` WHERE surat_tugas_id=?"
		for i := 0; i < len(surats); i++ {
			rinci := entity.RincianAnggaran{}
			tx.QueryRowContext(ctx, SQL, surats[i].Id).Scan(
				&rinci.Id,
				&rinci.SuratTugasId,
				&rinci.DokName,
				&rinci.DokPDF,
				&rinci.CreateAt,
			)
			surats[i].Rincian = rinci
		}
		return surats, nil
	}
}

// GetSuratPresensi implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetSuratPresensi(ctx context.Context, tx *sql.Tx, userId int) []entity.SuratTugasJOINPresensi {
	SQL := "SELECT DISTINCT `s`.id, `s`.lokasi_tujuan, `s`.jenis_program, `s`.tgl_awal, `s`.tgl_akhir "
	SQL += "FROM `surat_tugas` `s` "
	SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
	SQL += "LEFT JOIN `user` `u` on `u`.id =`s`.user_id "
	SQL += "LEFT JOIN `approved` `a` on `s`.id = `a`.surat_tugas_id "
	SQL += "WHERE (`s`.user_id = ? OR `p`.user_id = ?) AND `a`.status_ttd = '1' AND `s`.tgl_akhir > NOW();"

	var hasil []entity.SuratTugasJOINPresensi

	rows, err := tx.QueryContext(ctx, SQL, userId, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	for rows.Next() {
		surat := entity.SuratTugasJOINPresensi{}
		rows.Scan(
			&surat.Id,
			&surat.LokasiSurat,
			&surat.JenisProgram,
			&surat.TglAwal,
			&surat.TglAkhir,
		)

		surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
		surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)

		hasil = append(hasil, surat)
	}

	SQL = "SELECT id, gambar, lokasi, koordinat FROM `presensi` WHERE user_id=? AND surat_tugas_id=?"
	for i := 0; i < len(hasil); i++ {
		tx.QueryRowContext(ctx, SQL, userId, hasil[i].Id).Scan(
			&hasil[i].GambarId,
			&hasil[i].Gambar,
			&hasil[i].Lokasi,
			&hasil[i].Koordinat,
		)
	}

	return hasil
}

// LaporanGetAllSPPDByUserId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) LaporanGetAllSPPDByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApprovedLaporan, error) {
	SQL := "SELECT `s`.* "
	SQL += "FROM `surat_tugas` `s` "
	SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
	SQL += "LEFT JOIN `user` `u` on `u`.id =`s`.user_id "
	SQL += "LEFT JOIN `approved` `a` on `s`.id = `a`.surat_tugas_id "
	SQL += "WHERE (`s`.user_id = ? OR `p`.user_id = ?) AND `a`.status_ttd = '1' AND `s`.tgl_akhir > NOW();"

	surats := []entity.SuratTugasJOINApprovedLaporan{}

	rows, err := tx.QueryContext(ctx, SQL, userId, userId)
	if err != nil {
		return surats, err
	}
	defer rows.Close()

	for rows.Next() {
		surat := entity.SuratTugasJOINApprovedLaporan{}
		rows.Scan(
			&surat.Id,
			&surat.Tipe,
			&surat.UserId,
			&surat.LokasiTujuan,
			&surat.JenisProgram,
			&surat.DokPendukungName,
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
	return surats, nil
}

// GetStatusLaporanAkBySPPDId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetStatusLaporanAkBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) entity.LaporanJoinApproved {
	SQL := "SELECT * FROM `laporan_aktivitas` WHERE `surat_tugas_id` =?;"

	var laporanJoin entity.LaporanJoinApproved
	tx.QueryRowContext(ctx, SQL, suratId).Scan(
		&laporanJoin.Id,
		&laporanJoin.UserId,
		&laporanJoin.SuratTugasId,
		&laporanJoin.DokName,
		&laporanJoin.DokPDF,
		&laporanJoin.CreateAt,
	)

	SQL = "SELECT `status`, `message` FROM `approved_lap_ak` WHERE `laporan_id` =?;"
	tx.QueryRowContext(ctx, SQL, laporanJoin.Id).Scan(
		&laporanJoin.Status,
		&laporanJoin.Message,
	)

	return laporanJoin
}

// GetStatusLaporanAngBySPPDId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetStatusLaporanAngBySPPDId(ctx context.Context, tx *sql.Tx, suratId int) entity.LaporanJoinApproved {
	SQL := "SELECT * FROM `laporan_anggaran` WHERE `surat_tugas_id` =?;"

	var laporanJoin entity.LaporanJoinApproved
	tx.QueryRowContext(ctx, SQL, suratId).Scan(
		&laporanJoin.Id,
		&laporanJoin.UserId,
		&laporanJoin.SuratTugasId,
		&laporanJoin.DokName,
		&laporanJoin.DokPDF,
		&laporanJoin.CreateAt,
	)

	SQL = "SELECT `status`, `message` FROM `approved_lap_angg` WHERE `laporan_id` =?;"
	tx.QueryRowContext(ctx, SQL, laporanJoin.Id).Scan(
		&laporanJoin.Status,
		&laporanJoin.Message,
	)

	return laporanJoin
}

// LaporanGetSPPDById implements PegawaiRepo.
func (pr *PegawaiRepoImpl) LaporanGetSPPDById(ctx context.Context, tx *sql.Tx, suratId int) (entity.SuratTugasJOINUserParticipan, error) {
	SQL := "SELECT `surat_tugas`.*, `user`.nip, `user`.name, `user`.no_telp, `user`.email "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "INNER JOIN `user` ON `surat_tugas`.user_id = `user`.id WHERE `surat_tugas`.id= ?;"
	surat := entity.SuratTugasJOINUserParticipan{}
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
	surat.CreateAt = helper.ConvertSQLTimeStamp(surat.CreateAt)

	if err != nil {
		return surat, err
	}
	return surat, nil
}

// GetFotoByUserIdAndSPPDId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetFotoByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, request laporanreqres.LaporanGetSPPDByIdRequest) entity.Presensi {
	SQL := "SELECT * FROM `presensi` WHERE `surat_tugas_id` =? AND `user_id`=?"
	presensi := entity.Presensi{}
	tx.QueryRowContext(ctx, SQL, request.SuratTugasId, request.UserId).Scan(
		&presensi.Id,
		&presensi.UserId,
		&presensi.SuratTugasId,
		&presensi.Gambar,
		&presensi.Lokasi,
		&presensi.Koordinat,
		&presensi.CreateAt,
	)
	return presensi
}

// GetLaporanAktivitasByUserIdAndSPPDId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetLaporanAktivitasByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) entity.LaporanJoinApproved {
	SQL := "SELECT * FROM `laporan_aktivitas` WHERE `surat_tugas_id` =?;"

	var laporanJoin entity.LaporanJoinApproved
	tx.QueryRowContext(ctx, SQL, laporan.SuratTugasId).Scan(
		&laporanJoin.Id,
		&laporanJoin.UserId,
		&laporanJoin.SuratTugasId,
		&laporanJoin.DokName,
		&laporanJoin.DokPDF,
		&laporanJoin.CreateAt,
	)

	SQL = "SELECT `status`, `message` FROM `approved_lap_ak` WHERE `laporan_id` =?;"
	tx.QueryRowContext(ctx, SQL, laporanJoin.Id).Scan(
		&laporanJoin.Status,
		&laporanJoin.Message,
	)

	return laporanJoin
}

// GetLaporanAnggaranByUserIdAndSPPDId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetLaporanAnggaranByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) entity.LaporanJoinApproved {
	SQL := "SELECT * FROM `laporan_anggaran` WHERE `surat_tugas_id` =?;"

	var laporanJoin entity.LaporanJoinApproved
	tx.QueryRowContext(ctx, SQL, laporan.SuratTugasId).Scan(
		&laporanJoin.Id,
		&laporanJoin.UserId,
		&laporanJoin.SuratTugasId,
		&laporanJoin.DokName,
		&laporanJoin.DokPDF,
		&laporanJoin.CreateAt,
	)

	SQL = "SELECT `status`, `message` FROM `approved_lap_angg` WHERE `laporan_id` =?;"
	tx.QueryRowContext(ctx, SQL, laporanJoin.Id).Scan(
		&laporanJoin.Status,
		&laporanJoin.Message,
	)

	return laporanJoin
}

// UploadLaporanAct implements PegawaiRepo.
func (pr *PegawaiRepoImpl) UploadLaporanAct(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error) {
	SQL := "INSERT INTO `laporan_aktivitas`(`surat_tugas_id`, `user_id`, `dok_laporan_name`, `dok_laporan_pdf`) VALUES(?, ?, ?, ?);"
	result, err := tx.ExecContext(ctx, SQL,
		laporan.SuratTugasId,
		laporan.UserId,
		laporan.DokLaporanName,
		laporan.DokLaporanPDF,
	)
	if err != nil {
		return laporan, err
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		laporan.Id = int(id)
		SQL = "INSERT INTO `approved_lap_ak`(`laporan_id`, `user_id`, `status`, `message`) VALUES(?, ?, '0', '0');"
		_, err = tx.ExecContext(ctx, SQL,
			laporan.Id,
			laporan.UserId,
		)
		helper.PanicIfError(err)
		return laporan, nil
	}
}

// UploadLaporanAngg implements PegawaiRepo.
func (pr *PegawaiRepoImpl) UploadLaporanAngg(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error) {
	SQL := "INSERT INTO `laporan_anggaran` (`surat_tugas_id`, `user_id`, `dok_laporan_name`, `dok_laporan_pdf`) VALUES(?, ?, ?, ?);"
	result, err := tx.ExecContext(ctx, SQL,
		laporan.SuratTugasId,
		laporan.UserId,
		laporan.DokLaporanName,
		laporan.DokLaporanPDF,
	)
	if err != nil {
		return laporan, errors.New("error upload laporan")
	} else {
		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		laporan.Id = int(id)

		SQL = "INSERT INTO `approved_lap_angg`(`laporan_id`, `user_id`, `status`, `message`) VALUES(?, ?, '0', '0');"
		_, err = tx.ExecContext(ctx, SQL,
			laporan.Id,
			laporan.UserId,
		)
		helper.PanicIfError(err)
		return laporan, nil
	}
}

// SetLaporanAct implements PegawaiRepo.
func (pr *PegawaiRepoImpl) SetLaporanAct(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error) {
	SQL := "SELECT id FROM`laporan_aktivitas` WHERE `surat_tugas_id`=?;"
	err := tx.QueryRowContext(ctx, SQL, laporan.SuratTugasId).Scan(&laporan.Id)

	helper.PanicIfError(err)
	SQL = "UPDATE `laporan_aktivitas` SET `user_id`=?, `dok_laporan_name`=?, `dok_laporan_pdf`=?, `create_at`=NOW() WHERE `id`=?;"
	_, err = tx.ExecContext(ctx, SQL, laporan.UserId, laporan.DokLaporanName, laporan.DokLaporanPDF, laporan.Id)
	helper.PanicIfError(err)

	SQL = "UPDATE `approved_lap_ak` SET `user_id`=?, `status`='0', `message`='0', `create_at`=NOW() WHERE `laporan_id`=?;"
	_, err = tx.ExecContext(ctx, SQL, laporan.UserId, laporan.Id)
	helper.PanicIfError(err)

	return laporan, nil
}

// SetLaporanAngg implements PegawaiRepo.
func (pr *PegawaiRepoImpl) SetLaporanAngg(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error) {
	SQL := "SELECT id FROM `laporan_anggaran` WHERE `surat_tugas_id`=?;"
	err := tx.QueryRowContext(ctx, SQL, laporan.SuratTugasId).Scan(&laporan.Id)
	helper.PanicIfError(err)

	SQL = "UPDATE `laporan_anggaran` SET `user_id`=?, `dok_laporan_name`=?, `dok_laporan_pdf`=?, `create_at`=NOW() WHERE `id`=?;"
	_, err = tx.ExecContext(ctx, SQL, laporan.UserId, laporan.DokLaporanName, laporan.DokLaporanPDF, laporan.Id)
	helper.PanicIfError(err)

	SQL = "UPDATE `approved_lap_angg` SET `user_id`=?, `status`='0', `message`='0', `create_at`=NOW() WHERE `laporan_id`=?;"
	_, err = tx.ExecContext(ctx, SQL, laporan.UserId, laporan.Id)
	helper.PanicIfError(err)

	return laporan, nil
}

// Profile implements PegawaiRepo.
func (pr *PegawaiRepoImpl) Profile(ctx context.Context, tx *sql.Tx, userId int) entity.User {
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
