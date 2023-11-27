package helper

import (
	"github.com/fzialam/workAway/model/entity"
	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	presensireqres "github.com/fzialam/workAway/model/req_res/presensi_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

// User Section
func ToUserResponse(user entity.User) userreqres.UserResponse {
	return userreqres.UserResponse{
		Id:       user.Id,
		NIP:      user.NIP,
		NIK:      user.NIK,
		NPWP:     user.NPWP,
		Name:     user.Name,
		Rank:     user.Rank,
		NoTelp:   user.NoTelp,
		TglLahir: user.TglLahir,
		Status:   user.Status,
		Gender:   user.Gender,
		Alamat:   user.Alamat,
		Email:    user.Email,
	}
}

func ToUserResponses(users []entity.User) []userreqres.UserResponse {
	var userResponses []userreqres.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

// Izin Section
func ToIzinResponses(izin entity.Izin) izinreqres.IzinResponse {
	return izinreqres.IzinResponse{
		Status:  izin.Status,
		Message: "Success",
	}
}

// Surat Tugas Section
func ToSuratTugasResponse(surat entity.SuratTugas) surattugasreqres.SuratTugasResponse {
	return surattugasreqres.SuratTugasResponse{
		Id:               surat.Id,
		Tipe:             surat.Tipe,
		UserId:           surat.UserId,
		LokasiTujuan:     surat.LokasiTujuan,
		JenisProgram:     surat.JenisProgram,
		DokumenName:      surat.DokumenName,
		DokumenPDF:       surat.DokumenPDF,
		DokPendukungName: surat.DokPendukungName,
		DokPendukungPdf:  surat.DokPendukungPdf,
		TglAwal:          surat.TglAwal,
		TglAkhir:         surat.TglAkhir,
	}
}

// Join Model Section
func ToSuratTugasJOINApprovedUserResponse(surat entity.SuratTugasJOINApprovedUser) surattugasreqres.SuratTugasJOINApprovedUserResponse {
	return surattugasreqres.SuratTugasJOINApprovedUserResponse{
		Id:               surat.Id,
		UserId:           surat.UserId,
		LokasiTujuan:     surat.LokasiTujuan,
		JenisProgram:     surat.JenisProgram,
		DokumenName:      surat.DokumenName,
		DokumenPDF:       surat.DokumenPDF,
		DokPendukungName: surat.DokPendukungName,
		DokPendukungPdf:  surat.DokPendukungPdf,
		TglAwal:          surat.TglAwal,
		TglAkhir:         surat.TglAkhir,
		CreateAt:         surat.CreateAt,
		Status:           surat.Status,
		StatusTTD:        surat.StatusTTD,
		UserNIP:          surat.UserNIP,
		UserName:         surat.UserName,
		UserNoTelp:       surat.UserNoTelp,
		UserEmail:        surat.UserEmail,
	}
}

func ToSuratTugasJOINApprovedUserResponses(surats []entity.SuratTugasJOINApprovedUser) []surattugasreqres.SuratTugasJOINApprovedUserResponse {
	var suratTugasJOINApprovedUserResponses []surattugasreqres.SuratTugasJOINApprovedUserResponse
	for _, surat := range surats {
		suratTugasJOINApprovedUserResponses = append(suratTugasJOINApprovedUserResponses, ToSuratTugasJOINApprovedUserResponse(surat))
	}
	return suratTugasJOINApprovedUserResponses
}

func ToSuratTugasJOINApprovedUserParticipanResponse(req entity.SuratTugasJOINApprovedUserParticipan) surattugasreqres.SuratTugasJOINApprovedUserParticipanResponse {
	return surattugasreqres.SuratTugasJOINApprovedUserParticipanResponse{
		Id:               req.Id,
		UserId:           req.UserId,
		LokasiTujuan:     req.LokasiTujuan,
		JenisProgram:     req.JenisProgram,
		DokumenName:      req.DokumenName,
		DokumenPDF:       req.DokumenPDF,
		DokPendukungName: req.DokPendukungName,
		DokPendukungPdf:  req.DokPendukungPdf,
		TglAwal:          req.TglAwal,
		TglAkhir:         req.TglAkhir,
		CreateAt:         req.CreateAt,
		Status:           req.Status,
		UserNIP:          req.UserNIP,
		UserName:         req.UserName,
		UserNoTelp:       req.UserNoTelp,
		UserEmail:        req.UserEmail,
		Participans:      req.Participans,
	}
}

func ToSuratTugasJOINApprovedResponse(surat entity.SuratTugasJOINApproved) surattugasreqres.SuratTugasJOINApprovedResponse {
	return surattugasreqres.SuratTugasJOINApprovedResponse{
		Id:               surat.Id,
		UserId:           surat.UserId,
		LokasiTujuan:     surat.LokasiTujuan,
		JenisProgram:     surat.JenisProgram,
		DokumenName:      surat.DokumenName,
		DokumenPDF:       surat.DokumenPDF,
		DokPendukungName: surat.DokPendukungName,
		DokPendukungPdf:  surat.DokPendukungPdf,
		TglAwal:          surat.TglAwal,
		TglAkhir:         surat.TglAkhir,
		Status:           surat.Status,
	}
}

func ToSuratTugasJOINApprovedResponses(surats []entity.SuratTugasJOINApproved) []surattugasreqres.SuratTugasJOINApprovedResponse {
	var suratResponses []surattugasreqres.SuratTugasJOINApprovedResponse
	for _, surat := range surats {
		suratResponses = append(suratResponses, ToSuratTugasJOINApprovedResponse(surat))
	}
	return suratResponses
}

// Presensi Section
func ToPresensiResponse(presensi entity.Presensi) presensireqres.PresensiFotoResponse {
	return presensireqres.PresensiFotoResponse{
		Id:           presensi.Id,
		UserId:       presensi.UserId,
		SuratTugasId: presensi.SuratTugasId,
		Gambar:       "Sukses Upload Gambar",
		Lokasi:       presensi.Lokasi,
	}
}

// Permohonan Section
func ToPermohonanResponse(surat entity.SuratTugas, participan entity.Participan) permohonanreqres.PermohonanResponse {
	return permohonanreqres.PermohonanResponse{
		UserPemohonId:    surat.UserId,
		LokasiTujuan:     surat.LokasiTujuan,
		JenisProgram:     surat.JenisProgram,
		DokPendukungName: surat.DokumenName,
		ParticipansId:    participan.UserId,
		TglAwal:          surat.TglAwal,
		TglAkhir:         surat.TglAkhir,
	}
}

// Penugasan Section
func ToPenugasanResponse(surat entity.SuratTugas, participan entity.Participan) penugasanreqres.PenugasanResponse {
	return penugasanreqres.PenugasanResponse{
		UserKetuaId:   surat.UserId,
		LokasiTujuan:  surat.LokasiTujuan,
		JenisProgram:  surat.JenisProgram,
		ParticipansId: participan.UserId,
		TglAwal:       surat.TglAwal,
		TglAkhir:      surat.TglAkhir,
	}
}

func ToSuratTugasJOINApprovedLaporanResponse(surat entity.SuratTugasJOINApprovedLaporan) surattugasreqres.SuratTugasJOINApprovedLaporanResponse {
	return surattugasreqres.SuratTugasJOINApprovedLaporanResponse{
		Id:               surat.Id,
		Tipe:             surat.Tipe,
		UserId:           surat.UserId,
		LokasiTujuan:     surat.LokasiTujuan,
		JenisProgram:     surat.JenisProgram,
		DokumenName:      surat.DokumenName,
		DokumenPDF:       surat.DokumenPDF,
		DokPendukungName: surat.DokPendukungName,
		DokPendukungPdf:  surat.DokPendukungPdf,
		TglAwal:          surat.TglAwal,
		TglAkhir:         surat.TglAkhir,
		CreateAt:         surat.CreateAt,
		StatusPimpinan:   surat.StatusPimpinan,
		StatusKeuangan:   surat.StatusKeuangan,
	}

}

func ToSuratTugasJOINApprovedLaporanResponses(surats []entity.SuratTugasJOINApprovedLaporan) []surattugasreqres.SuratTugasJOINApprovedLaporanResponse {
	var suratResponses []surattugasreqres.SuratTugasJOINApprovedLaporanResponse
	for _, surat := range surats {
		suratResponses = append(suratResponses, ToSuratTugasJOINApprovedLaporanResponse(surat))
	}
	return suratResponses
}

func ToSuratTugasJOINUserParticipanLaporanResponse(
	surat entity.SuratTugasJOINUserParticipan,
	presensi entity.Presensi,
	lapAkAngg entity.LaporanAktivitasAnggaran) surattugasreqres.SuratTugasJOINUserParticipanLaporanResponse {
	return surattugasreqres.SuratTugasJOINUserParticipanLaporanResponse{
		Id:                   surat.Id,
		Tipe:                 surat.Tipe,
		UserId:               surat.UserId,
		LokasiTujuan:         surat.LokasiTujuan,
		JenisProgram:         surat.JenisProgram,
		DokumenName:          surat.DokumenName,
		DokumenPDF:           surat.DokumenPDF,
		DokPendukungName:     surat.DokPendukungName,
		DokPendukungPdf:      surat.DokPendukungPdf,
		TglAwal:              surat.TglAwal,
		TglAkhir:             surat.TglAkhir,
		CreateAt:             surat.CreateAt,
		UserNIP:              surat.UserNIP,
		UserName:             surat.UserName,
		UserNoTelp:           surat.UserNoTelp,
		UserEmail:            surat.UserEmail,
		Participans:          surat.Participans,
		LaporanAktivitasName: lapAkAngg.DokAktivitasName,
		LaporanAktivitasPDF:  lapAkAngg.DokAktivitasPDF,
		LaporanAnggaranName:  lapAkAngg.DokAnggaranName,
		LaporanAnggaranPDF:   lapAkAngg.DokAnggaranPDF,
		Gambar:               presensi.Gambar,
		Lokasi:               presensi.Lokasi,
		Koordinat:            presensi.Koordinat,
	}
}
