package helper

import (
	"github.com/fzialam/workAway/model/entity"
	approvedreqres "github.com/fzialam/workAway/model/req_res/approved_req_res"
	keuanganreqres "github.com/fzialam/workAway/model/req_res/keuangan_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	presensireqres "github.com/fzialam/workAway/model/req_res/presensi_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

// User Section
func ToLoginResponse(user entity.User, token string) userreqres.LoginResponse {
	return userreqres.LoginResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Rank:  user.Rank,
		Token: token,
	}
}

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
		Gambar:   user.Gambar,
	}
}

func ToUserResponses(users []entity.User) []userreqres.UserResponse {
	var userResponses []userreqres.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToSuratTugasJoinUserResponse(surat entity.SuratTugasJOINUser) surattugasreqres.SuratTugasJOINUserResponse {
	return surattugasreqres.SuratTugasJOINUserResponse{
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
		UserNIP:          surat.UserNIP,
		UserName:         surat.UserName,
		UserNoTelp:       surat.UserNoTelp,
		UserEmail:        surat.UserEmail,
	}
}

func ToSuratTugasJoinUserResponses(users []entity.SuratTugasJOINUser) []surattugasreqres.SuratTugasJOINUserResponse {
	var responses []surattugasreqres.SuratTugasJOINUserResponse
	for _, user := range users {
		responses = append(responses, ToSuratTugasJoinUserResponse(user))
	}
	return responses
}

// Izin Section
func ToIzinResponses(izin entity.Approved) approvedreqres.ApprovedResponse {
	return approvedreqres.ApprovedResponse{
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
		CreateAt:         surat.CreateAt,
	}
}

func ToSuratTugasResponses(surats []entity.SuratTugas) []surattugasreqres.SuratTugasResponse {
	var responses []surattugasreqres.SuratTugasResponse

	for _, surat := range surats {
		responses = append(responses, ToSuratTugasResponse(surat))
	}

	return responses

}

func ToSuratTugasJOINRincianResponse(surat entity.SuratTugasJOINRincian) surattugasreqres.SuratTugasJOINRincianResponse {
	return surattugasreqres.SuratTugasJOINRincianResponse{
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
		Rincian:          surat.Rincian,
	}
}

func ToSuratTugasJOINRincianResponses(surats []entity.SuratTugasJOINRincian) []surattugasreqres.SuratTugasJOINRincianResponse {
	var responses []surattugasreqres.SuratTugasJOINRincianResponse
	for _, surat := range surats {
		responses = append(responses, ToSuratTugasJOINRincianResponse(surat))
	}
	return responses
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

func ToSuratTugasJOINUserParticipanResponse(req entity.SuratTugasJOINUserParticipan) surattugasreqres.SuratTugasJOINUserParticipanResponse {
	return surattugasreqres.SuratTugasJOINUserParticipanResponse{
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
		UserNIP:          req.UserNIP,
		UserName:         req.UserName,
		UserNoTelp:       req.UserNoTelp,
		UserEmail:        req.UserEmail,
		Participans:      req.Participans,
	}
}

func ToSuratTugasJOINDoubleApprovedUserParticipanResponse(req entity.SuratTugasJOINDoubleApprovedUserParticipan) surattugasreqres.SuratTugasJOINDoubleApprovedUserParticipanResponse {
	return surattugasreqres.SuratTugasJOINDoubleApprovedUserParticipanResponse{
		Id:               req.Id,
		Tipe:             req.Tipe,
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
		OtherStatus:      req.OtherStatus,
		UserNIP:          req.UserNIP,
		UserName:         req.UserName,
		UserNoTelp:       req.UserNoTelp,
		UserEmail:        req.UserEmail,
		Participans:      req.Participans,
	}
}

func ToSuratTugasJOINApprovedUserParticipanLaporanResponse(
	req entity.SuratTugasJOINApprovedUserParticipan,
	laporan entity.Laporan) surattugasreqres.SuratTugasJOINApprovedUserParticipanLaporanResponse {
	return surattugasreqres.SuratTugasJOINApprovedUserParticipanLaporanResponse{
		Id:               req.Id,
		Tipe:             req.Tipe,
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
		Laporan:          laporan,
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

func ToSuratTugasJOINSPPDApprovedAnggaranResponse(surat entity.SuratTugasJOINSPPDApprovedAnggaran) surattugasreqres.SuratTugasJOINSPPDApprovedAnggaranResponse {
	return surattugasreqres.SuratTugasJOINSPPDApprovedAnggaranResponse{
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
		Status:           surat.Status,
		Rincian:          surat.Rincian,
	}
}

func ToSuratTugasJOINSPPDApprovedAnggaranResponses(surats []entity.SuratTugasJOINSPPDApprovedAnggaran) []surattugasreqres.SuratTugasJOINSPPDApprovedAnggaranResponse {
	var suratResponses []surattugasreqres.SuratTugasJOINSPPDApprovedAnggaranResponse
	for _, surat := range surats {
		suratResponses = append(suratResponses, ToSuratTugasJOINSPPDApprovedAnggaranResponse(surat))
	}
	return suratResponses
}

func ToSuratTugasJOINPresensiResponse(surat entity.SuratTugasJOINPresensi) surattugasreqres.SuratTugasJOINPresensiResponse {
	return surattugasreqres.SuratTugasJOINPresensiResponse{
		Id:           surat.Id,
		LokasiSurat:  surat.LokasiSurat,
		JenisProgram: surat.JenisProgram,
		TglAwal:      surat.TglAwal,
		TglAkhir:     surat.TglAkhir,
		GambarId:     surat.GambarId,
		NameGambar:   surat.NameGambar,
		Gambar:       surat.Gambar,
		Lokasi:       surat.Lokasi,
		Koordinat:    surat.Koordinat,
	}
}

func ToSuratTugasJOINPresensiResponses(surats []entity.SuratTugasJOINPresensi) []surattugasreqres.SuratTugasJOINPresensiResponse {
	var suratResponses []surattugasreqres.SuratTugasJOINPresensiResponse
	for _, surat := range surats {
		suratResponses = append(suratResponses, ToSuratTugasJOINPresensiResponse(surat))
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
func ToPermohonanResponse(surat entity.SuratTugas) permohonanreqres.PermohonanResponse {
	return permohonanreqres.PermohonanResponse{
		LokasiTujuan:     surat.LokasiTujuan,
		JenisProgram:     surat.JenisProgram,
		DokPendukungName: surat.DokumenName,
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
		Message:          surat.Message,
	}
}

func ToSuratTugasJOINApprovedLaporanResponses(surats []entity.SuratTugasJOINApprovedLaporan) []surattugasreqres.SuratTugasJOINApprovedLaporanResponse {
	var suratResponses []surattugasreqres.SuratTugasJOINApprovedLaporanResponse
	for _, surat := range surats {
		suratResponses = append(suratResponses, ToSuratTugasJOINApprovedLaporanResponse(surat))
	}
	return suratResponses
}

func ToSuratTugasJOINUserParticipanLaporanJOINApprovedResponse(
	surat entity.SuratTugasJOINUserParticipan,
	presensi entity.Presensi,
	laporan []entity.LaporanJoinApproved) surattugasreqres.SuratTugasJOINUserParticipanLaporanJOINApprovedResponse {
	return surattugasreqres.SuratTugasJOINUserParticipanLaporanJOINApprovedResponse{
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
		UserNIP:          surat.UserNIP,
		UserName:         surat.UserName,
		UserNoTelp:       surat.UserNoTelp,
		UserEmail:        surat.UserEmail,
		Participans:      surat.Participans,
		Laporan:          laporan,
		Gambar:           presensi.Gambar,
		Lokasi:           presensi.Lokasi,
		Koordinat:        presensi.Koordinat,
	}
}

func ToSuratTugasJOINApprovedUserFotoParticipanFotoResponse(
	surat entity.SuratTugasJOINUserFoto,
	laporan entity.Laporan,
	isApproved string,
	participansFoto []entity.ParticipanJoinUserFoto,
) surattugasreqres.SuratTugasJOINUserFotoParticipanFotoLaporanStatusResponse {

	return surattugasreqres.SuratTugasJOINUserFotoParticipanFotoLaporanStatusResponse{
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
		UserNIP:          surat.UserNIP,
		UserName:         surat.UserName,
		UserNoTelp:       surat.UserNoTelp,
		UserEmail:        surat.UserEmail,
		Participans:      participansFoto,
		LaporanId:        laporan.Id,
		LaporanDokName:   laporan.DokLaporanName,
		LaporanDokPDF:    laporan.DokLaporanPDF,
		Status:           isApproved,
		NameGambar:       surat.UserNameGambar,
		Gambar:           surat.UserGambar,
		Lokasi:           surat.UserLokasi,
		Koordinat:        surat.UserKoordinat,
	}
}

func ToSuratTugasJOINUserLaporanApprovedResponse(surat entity.SuratTugasJOINUserLaporanApproved) surattugasreqres.SuratTugasJOINUserLaporanApprovedResponse {
	return surattugasreqres.SuratTugasJOINUserLaporanApprovedResponse{
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
		UserName:         surat.UserName,
		Laporan:          surat.Laporan,
	}
}

func ToSuratTugasJOINUserLaporanApprovedResponses(surats []entity.SuratTugasJOINUserLaporanApproved) []surattugasreqres.SuratTugasJOINUserLaporanApprovedResponse {
	var responses []surattugasreqres.SuratTugasJOINUserLaporanApprovedResponse

	for _, surat := range surats {
		responses = append(responses, ToSuratTugasJOINUserLaporanApprovedResponse(surat))
	}
	return responses
}

func ToSuratTugasJOINApprovedUserOtherIdResponse(surat entity.SuratTugasJOINApprovedUserOtherId) keuanganreqres.SuratTugasJOINApprovedUserOtherIdResponse {
	return keuanganreqres.SuratTugasJOINApprovedUserOtherIdResponse{
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
		Status:           surat.Status,
		OtherStatus:      surat.OtherStatus,
		UserNIP:          surat.UserNIP,
		UserName:         surat.UserName,
		UserNoTelp:       surat.UserNoTelp,
		UserEmail:        surat.UserEmail,
		OtherId:          surat.OtherId,
	}
}

func ToSuratTugasJOINApprovedUserOtherIdResponses(surats []entity.SuratTugasJOINApprovedUserOtherId) []keuanganreqres.SuratTugasJOINApprovedUserOtherIdResponse {
	var responses []keuanganreqres.SuratTugasJOINApprovedUserOtherIdResponse
	for _, surat := range surats {
		responses = append(responses, ToSuratTugasJOINApprovedUserOtherIdResponse(surat))
	}
	return responses
}
