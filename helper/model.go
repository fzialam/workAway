package helper

import (
	"github.com/fzialam/workAway/model/entity"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	persetujuanreqres "github.com/fzialam/workAway/model/req_res/persetujuan_req_res"
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

func ToGetSuratResponse(surat entity.SuratTugasJOINApproved) surattugasreqres.SuratTugasJOINApprovedResponse {
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

func ToGetSuratResponses(surats []entity.SuratTugasJOINApproved) []surattugasreqres.SuratTugasJOINApprovedResponse {
	var suratResponses []surattugasreqres.SuratTugasJOINApprovedResponse
	for _, surat := range surats {
		suratResponses = append(suratResponses, ToGetSuratResponse(surat))
	}
	return suratResponses
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

// Persetujuan Section
func ToPersetujuanResponses(izin entity.Izin) persetujuanreqres.PersetujuanResponse {
	return persetujuanreqres.PersetujuanResponse{
		Status:  izin.Status,
		Message: "Success",
	}

}
func ToSuratTugasJOINApprovedUserResponse(surat entity.SuratTugasJOINApprovedUser) persetujuanreqres.SuratTugasJOINApprovedUserResponse {
	return persetujuanreqres.SuratTugasJOINApprovedUserResponse{
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
		Create_at:        surat.Create_at,
		Status:           surat.Status,
		UserNIP:          surat.UserNIP,
		UserName:         surat.UserName,
		UserNoTelp:       surat.UserNoTelp,
		UserEmail:        surat.UserEmail,
	}
}

func ToSuratTugasJOINApprovedUserResponses(surats []entity.SuratTugasJOINApprovedUser) []persetujuanreqres.SuratTugasJOINApprovedUserResponse {
	var suratTugasJOINApprovedUserResponses []persetujuanreqres.SuratTugasJOINApprovedUserResponse
	for _, surat := range surats {
		suratTugasJOINApprovedUserResponses = append(suratTugasJOINApprovedUserResponses, ToSuratTugasJOINApprovedUserResponse(surat))
	}
	return suratTugasJOINApprovedUserResponses
}
