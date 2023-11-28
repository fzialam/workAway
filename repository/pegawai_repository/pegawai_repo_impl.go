package pegawairepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
)

type PegawaiRepoImpl struct {
}

func NewPegawaiRepo() PegawaiRepo {
	return &PegawaiRepoImpl{}
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

// LaporanGetAllSPPDByUserId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) LaporanGetAllSPPDByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApprovedLaporan, error) {
	SQL := "SELECT `s`.*, `ala`.status as 'status_ala', `alg`.status as 'status_alg' "
	SQL += "FROM `surat_tugas` `s` "
	SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
	SQL += "LEFT JOIN `user` `u` on `u`.id =`s`.user_id "
	SQL += "LEFT JOIN `approved_lap_ak` `ala` on `ala`.surat_tugas_id =`s`.id  "
	SQL += "LEFT JOIN `approved_lap_angg` `alg` on `alg`.surat_tugas_id =`s`.id "
	SQL += "LEFT JOIN `approved` `a` on `s`.id = `a`.surat_tugas_id "
	SQL += "WHERE (`s`.user_id = ? OR `p`.user_id = ?) AND `a`.status_ttd = '1' AND `s`.tgl_awal > NOW();"

	surats := []entity.SuratTugasJOINApprovedLaporan{}

	rows, err := tx.QueryContext(ctx, SQL, userId, userId)
	if err != nil {
		return surats, err
	}

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
			&surat.StatusPimpinan,
			&surat.StatusKeuangan,
		)
		surat.TglAwal = helper.ConvertSQLTimeToHTML(surat.TglAwal)
		surat.TglAkhir = helper.ConvertSQLTimeToHTML(surat.TglAkhir)
		surat.CreateAt = helper.ConvertSQLTimeStamp(surat.CreateAt)
		surats = append(surats, surat)
	}

	// log.Println(surats[0].Id)
	return surats, nil
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

// CreatePermohonan implements PegawaiRepo.
func (pr *PegawaiRepoImpl) CreateSurat(ctx context.Context, tx *sql.Tx, surat entity.SuratTugas) (entity.SuratTugas, error) {
	SQL := "INSERT INTO `surat_tugas`(`tipe`, `user_id`,`lokasi_tujuan`,`jenis_program`,`dokumen_name`, `dokumen_pdf`, "
	SQL += "`dok_pendukung_name`, `dok_pendukung_pdf`,`tgl_awal`, `tgl_akhir`) "
	SQL += "VALUES (?,?,?,?,?,?,?,?,?,?)"
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

// PresensiFoto implements PegawaiRepo.
func (pr *PegawaiRepoImpl) PresensiFoto(ctx context.Context, tx *sql.Tx, presensi entity.Presensi) (entity.Presensi, error) {
	SQL := "INSERT INTO `presensi`(`user_id`, `surat_tugas_id`, `gambar`, `lokasi`) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, presensi.UserId, presensi.SuratTugasId, presensi.Gambar, presensi.Lokasi)
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
func (pr *PegawaiRepoImpl) GetSurat(ctx context.Context, tx *sql.Tx, userId int) ([]entity.SuratTugasJOINApproved, error) {
	SQL := "SELECT surat_tugas.*, approved.status, approved.status_ttd, approved.status_ttd_created_at "
	SQL += "FROM `surat_tugas` "
	SQL += "INNER JOIN `approved` ON `surat_tugas`.id = `approved`.surat_tugas_id "
	SQL += "WHERE `surat_tugas`.tgl_akhir > NOW() AND "
	SQL += "`surat_tugas`.user_id = ? AND "
	SQL += "(approved.status_ttd_created_at = '0' OR approved.status_ttd_created_at = '1');"
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

// GetFotoByUserIdAndSPPDId implements PegawaiRepo.
func (*PegawaiRepoImpl) GetFotoByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, request laporanreqres.LaporanGetSPPDByIdRequest) entity.Presensi {
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
func (pr *PegawaiRepoImpl) GetLaporanAktivitasByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error) {
	SQL := "SELECT dok_laporan_name, dok_laporan_pdf FROM `laporan_aktivitas` WHERE `surat_tugas_id` =? AND `user_id`=?"
	err := tx.QueryRowContext(ctx, SQL, laporan.SuratTugasId, laporan.UserId).Scan(
		&laporan.DokLaporanName,
		&laporan.DokLaporanPDF,
	)
	helper.PanicIfError(err)

	return laporan, nil
}

// GetLaporanAnggaranByUserIdAndSPPDId implements PegawaiRepo.
func (pr *PegawaiRepoImpl) GetLaporanAnggaranByUserIdAndSPPDId(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error) {
	SQL := "SELECT dok_laporan_name, dok_laporan_pdf FROM `laporan_anggaran` WHERE `surat_tugas_id` =? AND `user_id`=?"
	err := tx.QueryRowContext(ctx, SQL, laporan.SuratTugasId, laporan.UserId).Scan(
		&laporan.DokLaporanName,
		&laporan.DokLaporanPDF,
	)
	helper.PanicIfError(err)

	return laporan, nil
}

// UploadLaporanAct implements PegawaiRepo.
func (pr *PegawaiRepoImpl) UploadLaporanAct(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error) {
	SQL := "INSERT TO `laporan_aktivitas`(`surat_tugas_id`, `user_id`, `dok_laporan_name`, `dok_laporan_pdf`) VALUES(?, ?, ?, ?);"
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
		return laporan, nil
	}
}

// UploadLaporanAngg implements PegawaiRepo.
func (pr *PegawaiRepoImpl) UploadLaporanAngg(ctx context.Context, tx *sql.Tx, laporan entity.Laporan) (entity.Laporan, error) {
	SQL := "INSERT TO `laporan_anggaran` (`surat_tugas_id`, `user_id`, `dok_laporan_name`, `dok_laporan_pdf`) VALUES(?, ?, ?, ?);"
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
		return laporan, nil
	}
}
