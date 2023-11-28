package tests

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	pimpinanrepository "github.com/fzialam/workAway/repository/pimpinan_repository"
	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	db, err := sql.Open("mysql", "rootsql:@tcp(localhost:3306)/workaway")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// user := entity.User{
	// 	Email:    "email@unesa.ac.id",
	// 	Password: "password",
	// }

	// `presensi`.name, `presensi`.gambar, `presensi`.lokasi, `presensi`.koordinat

	// SQL := "SELECT `laporan_aktivitas`.id, `laporan_aktivitas`.dok_laporan_name, `laporan_aktivitas`.dok_laporan_pdf, `laporan_anggaran`.id,`laporan_anggaran`.dok_laporan_name, `laporan_anggaran`.dok_laporan_pdf "
	// SQL += "FROM `surat_tugas` "
	// SQL += "INNER JOIN `laporan_aktivitas` on `surat_tugas`.id = `laporan_aktivitas`.surat_tugas_id "
	// SQL += "INNER JOIN `laporan_anggaran` on `surat_tugas`.id = `laporan_anggaran`.surat_tugas_id "
	// SQL += "WHERE `surat_tugas`.id = ?;"

	// var laporan entity.Laporan
	ctx := context.Background()
	tx, err := db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	stjf := pimpinanrepository.NewPimpinanRepo().GetFotoKetuaSPPDById(ctx, tx, entity.SuratTugasJOINUserFoto{
		UserId: 1,
		Id:     32,
	})
	// helper.PanicIfError(err)

	fmt.Println(stjf.UserGambar)
	// row := db.QueryRow(SQL, 70)

	// err = row.Scan(
	// 	&laporan.DokAktivitasId,
	// 	&laporan.DokAktivitasName,
	// 	&laporan.DokAktivitasPDF,
	// 	&laporan.DokAnggaranId,
	// 	&laporan.DokAnggaranName,
	// 	&laporan.DokAnggaranPDF,
	// )
	// // helper.PanicIfError(err)
	// log.Println(err)

	// // return laporan, nil
	// fmt.Println(laporan.DokAktivitasName)
	// fmt.Println(laporan.DokAnggaranName)
}
