package pimpinanservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	pimpinanreqres "github.com/fzialam/workAway/model/req_res/pimpinan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
	pimpinanrepository "github.com/fzialam/workAway/repository/pimpinan_repository"
	"github.com/go-playground/validator/v10"
)

type PimpinanServiceImpl struct {
	PimpinanRepo pimpinanrepository.PimpinanRepo
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewPimpinanService(pimpinanRepo pimpinanrepository.PimpinanRepo, db *sql.DB, validate *validator.Validate) PimpinanService {
	return &PimpinanServiceImpl{
		PimpinanRepo: pimpinanRepo,
		DB:           db,
		Validate:     validate,
	}
}

// CreatePenugasan implements PimpinanService.
func (ps *PimpinanServiceImpl) CreatePenugasan(ctx context.Context, request penugasanreqres.PenugasanRequest) penugasanreqres.PenugasanResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := entity.SuratTugas{
		Tipe:             request.Tipe,
		UserId:           request.UserKetuaId,
		DokumenName:      "-",
		DokumenPDF:       "-",
		LokasiTujuan:     request.LokasiTujuan,
		JenisProgram:     request.JenisProgram,
		DokPendukungName: "-",
		DokPendukungPdf:  "-",
		TglAwal:          request.TglAwal,
		TglAkhir:         request.TglAkhir,
	}
	surat, err = ps.PimpinanRepo.CreateSurat(ctx, tx, surat)
	helper.PanicIfError(err)

	participan := entity.Participan{
		UserId:       request.ParticipansId,
		SuratTugasId: surat.Id,
	}

	if len(participan.UserId) > 0 {
		participan, err = ps.PimpinanRepo.AddParticipans(ctx, tx, participan)
		helper.PanicIfError(err)
	}

	ps.PimpinanRepo.SPPDSetApproved(ctx, tx, entity.Approved{
		SuratTugasId:      surat.Id,
		Status:            "1",
		Message:           "OK",
		StatusTTD:         "0",
		MessageTTD:        "0",
		CreateAt:          helper.TimeNowToString(),
		StatusTTDCreateAt: helper.TimeNowToString(),
	})

	return helper.ToPenugasanResponse(surat, participan)
}

// GetAllUserId implements PimpinanService.
func (ps *PimpinanServiceImpl) GetAllUserId(ctx context.Context) []userreqres.UserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := ps.PimpinanRepo.GetAllUserId(ctx, tx)
	return helper.ToUserResponses(user)
}

// PermohonanGetAllSuratTugasJOINApprovedUser implements PimpinanService.
func (ps *PimpinanServiceImpl) PermohonanGetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := ps.PimpinanRepo.PermohonanGetAllSuratTugasJOINApprovedUser(ctx, tx)
	helper.PanicIfError(err)
	return helper.ToSuratTugasJOINApprovedUserResponses(result)
}

// PermohonanGetSuratTugasById implements PimpinanService.
func (ps *PimpinanServiceImpl) PermohonanGetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINApprovedUserParticipanResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := ps.PimpinanRepo.PermohonanGetSuratTugasById(ctx, tx, suratId)
	helper.PanicIfError(err)

	participans := ps.PimpinanRepo.GetAllParticipanJOINUserBySuratId(ctx, tx, suratId)

	result.Participans = participans

	return helper.ToSuratTugasJOINApprovedUserParticipanResponse(result)
}

// PermohonanSetApproved implements PimpinanService.
func (ps *PimpinanServiceImpl) PermohonanSetApproved(ctx context.Context, request izinreqres.IzinRequest) izinreqres.IzinResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	izin := entity.Approved{
		SuratTugasId: request.SuratTugasId,
		Status:       request.Status,
		Message:      request.Message,
		StatusTTD:    request.StatusTTD,
		MessageTTD:   request.MessageTTD,
	}

	izin = ps.PimpinanRepo.PermohonanSetApproved(ctx, tx, izin)

	return helper.ToIzinResponses(izin)
}

// SPPDGetAllSuratTugasJOINApprovedUser implements PimpinanService.
func (ps *PimpinanServiceImpl) SPPDGetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := ps.PimpinanRepo.SPPDGetAllSuratTugasJOINApprovedUser(ctx, tx)
	helper.PanicIfError(err)
	return helper.ToSuratTugasJOINApprovedUserResponses(result)
}

// SPPDGetSuratTugasById implements PimpinanService.
func (ps *PimpinanServiceImpl) SPPDGetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := ps.PimpinanRepo.SPPDGetSuratTugasById(ctx, tx, suratId)
	helper.PanicIfError(err)

	return helper.ToSuratTugasResponse(result)
}

// SPPDSetApproved implements PimpinanService.
func (ps *PimpinanServiceImpl) SPPDSetApproved(ctx context.Context, request pimpinanreqres.UploadSPPDRequest) izinreqres.IzinResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	izin := entity.Approved{
		SuratTugasId: request.SuratTugasId,
		StatusTTD:    request.Status,
		MessageTTD:   request.Message,
	}

	izin = ps.PimpinanRepo.SPPDSetApproved(ctx, tx, izin)

	if izin.StatusTTD == "1" {
		err = ps.PimpinanRepo.UploadSPPDAproved(ctx, tx, request)
		helper.PanicIfError(err)
	}

	return helper.ToIzinResponses(izin)
}

// LaporanGetAllSPPD implements PimpinanService.
func (ps *PimpinanServiceImpl) LaporanGetAllSPPD(ctx context.Context) []surattugasreqres.SuratTugasJOINLaporanApprovedResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := ps.PimpinanRepo.LaporanGetAllSPPD(ctx, tx)

	for i := range surat {
		laporan := ps.PimpinanRepo.LaporanBySPPDId(ctx, tx, surat[i].Id)
		statusLaporan := ps.PimpinanRepo.LaporanIsApproved(ctx, tx, laporan.Id)

		laporan.Status = statusLaporan.Status
		surat[i].Laporan = laporan
	}
	return helper.ToSuratTugasJOINLaporanApprovedResponses(surat)
}

// LaporanSPPDById implements PimpinanService.
func (ps *PimpinanServiceImpl) LaporanSPPDById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINUserFotoParticipanFotoLaporanStatusResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat, err := ps.PimpinanRepo.LaporanSPPDById(ctx, tx, suratId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	surat = ps.PimpinanRepo.GetFotoKetuaSPPDById(ctx, tx, surat)

	laporan := ps.PimpinanRepo.GetLaporanSPPDById(ctx, tx, suratId)

	isApproved := ps.PimpinanRepo.IsLaporanApproved(ctx, tx, laporan.Id)

	participans := ps.PimpinanRepo.GetAllParticipanJOINUserBySuratId(ctx, tx, suratId)

	participansFoto := []entity.ParticipanJoinUserFoto{}
	if len(participans) > 0 {
		for i := range participans {
			parFo := ps.PimpinanRepo.GetAllFotoParticipanById(ctx, tx, participans[i])
			participansFoto = append(participansFoto, parFo)
		}
	}

	return helper.ToSuratTugasJOINApprovedUserFotoParticipanFotoResponse(surat, laporan, isApproved, participansFoto)
}

// SetApprovedLaporan implements PimpinanService.
func (ps *PimpinanServiceImpl) SetApprovedLaporan(ctx context.Context, request laporanreqres.ApprovedLaporanRequest) laporanreqres.ApprovedLaporanResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	laporan := entity.ApprovedLaporan{
		LaporanId: request.LaporanId,
		UserId:    request.UserId,
		Status:    request.Status,
		Message:   request.Message,
	}

	laporan = ps.PimpinanRepo.SetApprovedLaporan(ctx, tx, laporan)

	return laporanreqres.ApprovedLaporanResponse{
		Status:  laporan.Status,
		Message: "Success",
	}
}
