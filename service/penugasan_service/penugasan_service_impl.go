package penugasanservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
	penugasanrepository "github.com/fzialam/workAway/repository/penugasan_repository"
	"github.com/go-playground/validator/v10"
)

type PenugasanServiceImpl struct {
	PenugasanRepo penugasanrepository.PenugasanRepo
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewPenugasanService(penugasanRepo penugasanrepository.PenugasanRepo, db *sql.DB, validate *validator.Validate) PenugasanService {
	return &PenugasanServiceImpl{
		PenugasanRepo: penugasanRepo,
		DB:            db,
		Validate:      validate,
	}
}

// SetApproved implements PenugasanService.
func (ps *PenugasanServiceImpl) SetApproved(ctx context.Context, request izinreqres.IzinRequest) izinreqres.IzinResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	izin := entity.Izin{
		SuratTugasId: request.SuratTugasId,
		StatusTTD:    request.StatusTTD,
	}

	izin = ps.PenugasanRepo.SetApproved(ctx, tx, izin)

	return helper.ToIzinResponses(izin)
}

// CreatePenugasan implements PenugasanService.
func (ps *PenugasanServiceImpl) CreatePenugasan(ctx context.Context, request penugasanreqres.PenugasanRequest) penugasanreqres.PenugasanResponse {
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
	surat, err = ps.PenugasanRepo.CreateSurat(ctx, tx, surat)
	helper.PanicIfError(err)

	participan := entity.Participan{
		UserId:       request.ParticipansId,
		SuratTugasId: surat.Id,
	}
	if len(participan.UserId) > 0 {
		participan, err = ps.PenugasanRepo.AddParticipans(ctx, tx, participan)
		helper.PanicIfError(err)
	}

	ps.PenugasanRepo.SetApproved(ctx, tx, entity.Izin{
		SuratTugasId:      surat.Id,
		Status:            "1",
		StatusTTD:         "0",
		CreateAt:          helper.TimeNowToString(),
		StatusTTDCreateAt: helper.TimeNowToString(),
	})

	return helper.ToPenugasanResponse(surat, participan)
}

// GetAllUserId implements PenugasanService.
func (ps *PenugasanServiceImpl) GetAllUserId(ctx context.Context) []userreqres.UserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := ps.PenugasanRepo.GetAllUserID(ctx, tx)
	return helper.ToUserResponses(user)
}

// GetAllSuratTugasJOINApprovedUser implements PenugasanService.
func (ps *PenugasanServiceImpl) GetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := ps.PenugasanRepo.GetAllSuratTugasJOINApprovedUser(ctx, tx)
	helper.PanicIfError(err)
	return helper.ToSuratTugasJOINApprovedUserResponses(result)
}

// GetSuratTugasById implements PersetujuanService.
func (ps *PenugasanServiceImpl) GetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := ps.PenugasanRepo.GetSuratTugasById(ctx, tx, suratId)
	helper.PanicIfError(err)

	return helper.ToSuratTugasResponse(result)
}
