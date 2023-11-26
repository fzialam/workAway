package pegawaiservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
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

	participan := entity.Participan{
		UserId:       request.ParticipansId,
		SuratTugasId: surat.Id,
	}

	participan, err = ps.PegawaiRepo.AddParticipans(ctx, tx, participan)
	helper.PanicIfError(err)

	err = ps.PegawaiRepo.Set0Approved(ctx, tx, surat.Id)
	helper.PanicIfError(err)

	return helper.ToPermohonanResponse(surat, participan)
}

// GetAllUserId implements PegawaiService.
func (ps *PegawaiServiceImpl) GetAllUserId(ctx context.Context) []userreqres.UserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := ps.PegawaiRepo.GetAllUserID(ctx, tx)
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
	}

	err = ps.PegawaiRepo.CheckIzin(ctx, tx, presensi)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	presensi, err = ps.PegawaiRepo.PresensiFoto(ctx, tx, presensi)
	helper.PanicIfError(err)

	return helper.ToPresensiResponse(presensi)
}

// GetSurat implements PegawaiService.
func (ps *PegawaiServiceImpl) GetSurat(ctx context.Context, request presensireqres.GetSuratForPresensiRequest) []surattugasreqres.SuratTugasJOINApprovedResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat, err := ps.PegawaiRepo.GetSurat(ctx, tx, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSuratTugasJOINApprovedResponses(surat)
}
