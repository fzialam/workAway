package tests

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
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
	SQL := "SELECT `s`.*, `ala`.status as 'status_ala', `alg`.status as 'status_alg' "
	SQL += "FROM `surat_tugas` `s` "
	SQL += "LEFT JOIN `participan` `p` on `s`.id = `p`.surat_tugas_id "
	SQL += "INNER JOIN `user` `u` on `u`.id =`s`.user_id "
	SQL += "INNER JOIN `approved_lap_ak` `ala` on `ala`.surat_tugas_id =`s`.id  "
	SQL += "INNER JOIN `approved_lap_angg` `alg` on `alg`.surat_tugas_id =`s`.id "
	SQL += "LEFT JOIN `approved` `a` on `s`.id = `a`.surat_tugas_id "
	SQL += "WHERE (`s`.user_id = 1 OR `p`.user_id = 1) AND `a`.status_ttd = '1' AND `s`.tgl_awal > NOW();"
	// fmt.Println(SQL)
	rows, err := db.Query(SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	surats := []entity.SuratTugasJOINApprovedLaporan{}
	for rows.Next() {
		surat := entity.SuratTugasJOINApprovedLaporan{}
		err := rows.Scan(
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
		helper.PanicIfError(err)
		surats = append(surats, surat)
	}

	for _, i := range surats {
		fmt.Println(i.Id)
	}
}
