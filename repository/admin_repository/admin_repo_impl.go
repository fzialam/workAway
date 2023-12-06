package adminrepository

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
)

type AdminRepoImpl struct {
}

func NewAdminRepo() AdminRepo {
	return &AdminRepoImpl{}
}

// Permohonan implements AdminRepo.
func (ar *AdminRepoImpl) Permohonan(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINUser, error) {
	SQL := "SELECT `s`.*, `u`.name from `surat_tugas` `s` "
	SQL += "INNER JOIN `user` `u` ON `s`.user_id = `u`.id "
	SQL += "WHERE `s`.tipe=0;"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	surats := []entity.SuratTugasJOINUser{}
	for rows.Next() {
		var surat entity.SuratTugasJOINUser
		err := rows.Scan(
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
		helper.PanicIfError(err)

		surats = append(surats, surat)
	}
	return surats, nil
}

// Penugasan implements AdminRepo.
func (ar *AdminRepoImpl) Penugasan(ctx context.Context, tx *sql.Tx) ([]entity.SuratTugasJOINUser, error) {
	SQL := "SELECT `s`.*, `u`.name from `surat_tugas` `s` "
	SQL += "INNER JOIN `user` `u` ON `s`.user_id = `u`.id "
	SQL += "WHERE `s`.tipe=1;"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	surats := []entity.SuratTugasJOINUser{}
	for rows.Next() {
		var surat entity.SuratTugasJOINUser
		err := rows.Scan(
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
		helper.PanicIfError(err)

		surats = append(surats, surat)
	}
	return surats, nil
}

// LapAKK implements AdminRepo.
func (ar *AdminRepoImpl) LapAKK(ctx context.Context, tx *sql.Tx) ([]entity.Laporan, error) {
	SQL := "SELECT * FROM `laporan_aktivitas`;"

	var laporans []entity.Laporan
	row, err := tx.QueryContext(ctx, SQL)

	helper.PanicIfError(err)

	for row.Next() {
		laporan := entity.Laporan{}
		row.Scan(
			&laporan.Id,
			&laporan.SuratTugasId,
			&laporan.UserId,
			&laporan.DokLaporanName,
			&laporan.DokLaporanPDF,
			&laporan.CreateAt,
		)
		laporan.CreateAt = helper.ConvertSQLTimeStamp(laporan.CreateAt)

		laporans = append(laporans, laporan)
	}

	return laporans, nil
}

// LapAGG implements AdminRepo.
func (ar *AdminRepoImpl) LapAGG(ctx context.Context, tx *sql.Tx) ([]entity.Laporan, error) {
	SQL := "SELECT * FROM `laporan_anggaran`;"

	var laporans []entity.Laporan
	row, err := tx.QueryContext(ctx, SQL)

	helper.PanicIfError(err)

	for row.Next() {
		laporan := entity.Laporan{}
		row.Scan(
			&laporan.Id,
			&laporan.SuratTugasId,
			&laporan.UserId,
			&laporan.DokLaporanName,
			&laporan.DokLaporanPDF,
			&laporan.CreateAt,
		)
		laporan.CreateAt = helper.ConvertSQLTimeStamp(laporan.CreateAt)

		laporans = append(laporans, laporan)
	}

	return laporans, nil
}

// UserGET implements AdminRepo.
func (ar *AdminRepoImpl) UserGET(ctx context.Context, tx *sql.Tx) ([]entity.User, error) {
	SQL := "select * from `user` where `rank`!= 4"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	users := []entity.User{}
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
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
		helper.PanicIfError(err)

		users = append(users, user)
	}
	return users, nil
}

// UserGETById implements AdminRepo.
func (ar *AdminRepoImpl) UserGETById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error) {
	SQL := "select * from `user` where id =?"
	row := tx.QueryRowContext(ctx, SQL, userId)
	user := entity.User{}
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

	return user, nil
}

// UserPOST implements AdminRepo.
func (ar *AdminRepoImpl) UserPOST(ctx context.Context, tx *sql.Tx, user entity.User) error {
	SQL := "UPDATE `user` set `rank`=? where `id`=?;"
	_, err := tx.ExecContext(ctx, SQL, user.Rank, user.Id)
	helper.PanicIfError(err)

	return nil
}

// Profile implements AdminRepo.
func (ar *AdminRepoImpl) Profile(ctx context.Context, tx *sql.Tx, userId int) entity.User {
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
