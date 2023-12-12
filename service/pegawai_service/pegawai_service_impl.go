package pegawaiservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	dokumenreqres "github.com/fzialam/workAway/model/req_res/dokumen_req_res"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	pegawaireqres "github.com/fzialam/workAway/model/req_res/pegawai_req_res"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	presensireqres "github.com/fzialam/workAway/model/req_res/presensi_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
	pegawairepository "github.com/fzialam/workAway/repository/pegawai_repository"
	"github.com/go-playground/validator/v10"
)

type PegawaiServiceImpl struct {
	PegawaiRepo pegawairepository.PegawaiRepo
	DB          *sql.DB
	Validate    *validator.Validate
}

func NewPegawaiService(pegawaiRepo pegawairepository.PegawaiRepo, db *sql.DB, validate *validator.Validate) PegawaiService {
	return &PegawaiServiceImpl{
		PegawaiRepo: pegawaiRepo,
		DB:          db,
		Validate:    validate,
	}
}

// Index implements PegawaiService.
func (ps *PegawaiServiceImpl) Index(ctx context.Context, userId int) (pegawaireqres.IndexPegawai, error) {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	index, err := ps.PegawaiRepo.Index(ctx, tx, userId)
	helper.PanicIfError(err)

	return index, nil
}

// CreatePermohonan implements PegawaiService.
func (ps *PegawaiServiceImpl) CreatePermohonan(ctx context.Context, request permohonanreqres.PermohonanRequest) permohonanreqres.PermohonanResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := entity.SuratTugas{
		Tipe:             0,
		UserId:           request.UserPemohonId,
		LokasiTujuan:     request.LokasiTujuan,
		JenisProgram:     request.JenisProgram,
		DokPendukungName: request.DokPendukungName,
		DokPendukungPdf:  request.DokPendukungPdf,
		TglAwal:          request.TglAwal,
		TglAkhir:         request.TglAkhir,
	}
	surat, err = ps.PegawaiRepo.CreateSurat(ctx, tx, surat)
	helper.PanicIfError(err)

	if len(request.ParticipansId) > 0 {
		participan := entity.Participan{
			UserId:       request.ParticipansId,
			SuratTugasId: surat.Id,
		}

		_, err = ps.PegawaiRepo.AddParticipans(ctx, tx, participan)
		helper.PanicIfError(err)
	}

	err = ps.PegawaiRepo.Set0Approved(ctx, tx, surat.Id)
	helper.PanicIfError(err)

	return helper.ToPermohonanResponse(surat)
}

// GetAllUserId implements PegawaiService.
func (ps *PegawaiServiceImpl) GetAllUserId(ctx context.Context, userId int) []userreqres.UserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := ps.PegawaiRepo.GetAllUserID(ctx, tx, userId)
	return helper.ToUserResponses(user)
}

// PresensiFoto implements PegawaiService.
func (ps *PegawaiServiceImpl) PresensiFoto(ctx context.Context, request presensireqres.PresensiFotoRequest) presensireqres.PresensiFotoResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	presensi := entity.Presensi{
		UserId:       request.UserId,
		SuratTugasId: request.SuratTugasId,
		Gambar:       request.Gambar,
		Lokasi:       request.Lokasi,
		Koordinat:    request.Koordinat,
	}

	presensi, err = ps.PegawaiRepo.PresensiFoto(ctx, tx, presensi)
	helper.PanicIfError(err)

	return helper.ToPresensiResponse(presensi)
}

// GetSurat implements PegawaiService.
func (ps *PegawaiServiceImpl) GetSurat(ctx context.Context, request surattugasreqres.GetSuratRequest) []surattugasreqres.SuratTugasJOINSPPDApprovedAnggaranResponse {

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat, err := ps.PegawaiRepo.GetSurat(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSuratTugasJOINSPPDApprovedAnggaranResponses(surat)
}

func (ps *PegawaiServiceImpl) GetSuratPresensi(ctx context.Context, userId int) []surattugasreqres.SuratTugasJOINPresensiResponse {

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := ps.PegawaiRepo.GetSuratPresensi(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSuratTugasJOINPresensiResponses(surat)
}

// GetSuratById implements PegawaiService.
func (ps *PegawaiServiceImpl) GetSuratById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINUserParticipanResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat, err := ps.PegawaiRepo.GetSuratById(ctx, tx, suratId)
	helper.PanicIfError(err)

	participan, err := ps.PegawaiRepo.GetAllParticipanBySPPDId(ctx, tx, suratId)
	helper.PanicIfError(err)

	surat.Participans = participan

	return helper.ToSuratTugasJOINUserParticipanResponse(surat)
}

// LaporanGetAllSPPDByUserId implements PegawaiService.
func (ps *PegawaiServiceImpl) LaporanGetAllSPPDByUserId(ctx context.Context, userId int) []surattugasreqres.SuratTugasJOINApprovedLaporanResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surats, err := ps.PegawaiRepo.LaporanGetAllSPPDByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	for i := 0; i < len(surats); i++ {
		laporan := ps.PegawaiRepo.GetStatusLaporanAkBySPPDId(ctx, tx, surats[i].Id)
		surats[i].StatusPimpinan = laporan.Status
		if (laporan.Message != "OK") && (laporan.Message != "0") {
			surats[i].Message += laporan.Message
		}

		laporan = ps.PegawaiRepo.GetStatusLaporanAngBySPPDId(ctx, tx, surats[i].Id)
		surats[i].StatusKeuangan = laporan.Status
		if (laporan.Message != "OK") && (laporan.Message != "0") {
			if surats[i].Message == "" {
				surats[i].Message += " " + laporan.Message
			} else {
				surats[i].Message += laporan.Message
			}
		}
	}

	return helper.ToSuratTugasJOINApprovedLaporanResponses(surats)
}

// LaporanGetSPPDById implements PegawaiService.
func (ps *PegawaiServiceImpl) LaporanGetSPPDById(ctx context.Context, request laporanreqres.LaporanGetSPPDByIdRequest) surattugasreqres.SuratTugasJOINUserParticipanLaporanJOINApprovedResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat, err := ps.PegawaiRepo.LaporanGetSPPDById(ctx, tx, request.SuratTugasId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	presensi := ps.PegawaiRepo.GetFotoByUserIdAndSPPDId(ctx, tx, request)

	laporan := entity.Laporan{
		SuratTugasId: request.SuratTugasId,
		UserId:       request.UserId,
	}
	var laporans []entity.LaporanJoinApproved

	participan, err := ps.PegawaiRepo.GetAllParticipanBySPPDId(ctx, tx, surat.Id)
	helper.PanicIfError(err)

	surat.Participans = participan

	laporanAk := ps.PegawaiRepo.GetLaporanAktivitasByUserIdAndSPPDId(ctx, tx, laporan)

	laporans = append(laporans, laporanAk)

	laporanAg := ps.PegawaiRepo.GetLaporanAnggaranByUserIdAndSPPDId(ctx, tx, laporan)
	laporans = append(laporans, laporanAg)

	return helper.ToSuratTugasJOINUserParticipanLaporanJOINApprovedResponse(surat, presensi, laporans)
}

// UploadLapAktivitas implements PegawaiService.
func (ps *PegawaiServiceImpl) UploadLapAktivitas(ctx context.Context, request laporanreqres.UploadLaporanRequest) dokumenreqres.UploadDokumenResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	laporan := entity.Laporan{
		SuratTugasId:   request.SuratTugasId,
		UserId:         request.UserId,
		DokLaporanName: request.DokLaporanName,
		DokLaporanPDF:  request.DokLaporanPDF,
	}

	uploadLaporan, err := ps.PegawaiRepo.UploadLaporanAct(ctx, tx, laporan)
	helper.PanicIfError(err)

	return dokumenreqres.UploadDokumenResponse{
		DokLaporanName: uploadLaporan.DokLaporanName,
		Message:        "Success",
	}
}

// UploadLapAnggaran implements PegawaiService.
func (ps *PegawaiServiceImpl) UploadLapAnggaran(ctx context.Context, request laporanreqres.UploadLaporanRequest) dokumenreqres.UploadDokumenResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	laporan := entity.Laporan{
		SuratTugasId:   request.SuratTugasId,
		UserId:         request.UserId,
		DokLaporanName: request.DokLaporanName,
		DokLaporanPDF:  request.DokLaporanPDF,
	}

	uploadLaporan, err := ps.PegawaiRepo.UploadLaporanAngg(ctx, tx, laporan)
	helper.PanicIfError(err)

	return dokumenreqres.UploadDokumenResponse{
		DokLaporanName: uploadLaporan.DokLaporanName,
		Message:        "Success",
	}
}

// SetLapAktivitas implements PegawaiService.
func (ps *PegawaiServiceImpl) SetLapAktivitas(ctx context.Context, request laporanreqres.UploadLaporanRequest) dokumenreqres.UploadDokumenResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	laporan := entity.Laporan{
		SuratTugasId:   request.SuratTugasId,
		UserId:         request.UserId,
		DokLaporanName: request.DokLaporanName,
		DokLaporanPDF:  request.DokLaporanPDF,
	}

	laporan, err = ps.PegawaiRepo.SetLaporanAct(ctx, tx, laporan)
	helper.PanicIfError(err)

	return dokumenreqres.UploadDokumenResponse{
		DokLaporanName: laporan.DokLaporanName,
		Message:        "Success",
	}

}

// SetLapAnggaran implements PegawaiService.
func (ps *PegawaiServiceImpl) SetLapAnggaran(ctx context.Context, request laporanreqres.UploadLaporanRequest) dokumenreqres.UploadDokumenResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	laporan := entity.Laporan{
		SuratTugasId:   request.SuratTugasId,
		UserId:         request.UserId,
		DokLaporanName: request.DokLaporanName,
		DokLaporanPDF:  request.DokLaporanPDF,
	}

	laporan, err = ps.PegawaiRepo.SetLaporanAngg(ctx, tx, laporan)
	helper.PanicIfError(err)

	return dokumenreqres.UploadDokumenResponse{
		DokLaporanName: laporan.DokLaporanName,
		Message:        "Success",
	}

}

// Profile implements PegawaiService.
func (ps *PegawaiServiceImpl) Profile(ctx context.Context, userId int) userreqres.UserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := ps.PegawaiRepo.Profile(ctx, tx, userId)

	return helper.ToUserResponse(user)
}
