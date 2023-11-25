package permohonanservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
	permohonanrepository "github.com/fzialam/workAway/repository/permohonan_repository"
	"github.com/go-playground/validator/v10"
)

type PermohonanServiceImpl struct {
	PermohonanRepo permohonanrepository.PermohonanRepo
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPermohonanService(permohonanRepo permohonanrepository.PermohonanRepo, db *sql.DB, validate *validator.Validate) PermohonanService {
	return &PermohonanServiceImpl{
		PermohonanRepo: permohonanRepo,
		DB:             db,
		Validate:       validate,
	}
}

// CreatePermohonan implements PermohonanService.
func (ps *PermohonanServiceImpl) CreatePermohonan(ctx context.Context, request permohonanreqres.PermohonanRequest) permohonanreqres.PermohonanResponse {
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
	surat, err = ps.PermohonanRepo.CreateSurat(ctx, tx, surat)
	helper.PanicIfError(err)

	participan := entity.Participan{
		UserId:       request.ParticipansId,
		SuratTugasId: surat.Id,
	}

	participan, err = ps.PermohonanRepo.AddParticipans(ctx, tx, participan)
	helper.PanicIfError(err)

	err = ps.PermohonanRepo.Set0Approved(ctx, tx, surat.Id)
	helper.PanicIfError(err)

	return helper.ToPermohonanResponse(surat, participan)
}

// GetAllUserId implements PermohonanService.
func (ps *PermohonanServiceImpl) GetAllUserId(ctx context.Context) []userreqres.UserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := ps.PermohonanRepo.GetAllUserID(ctx, tx)
	return helper.ToUserResponses(user)
}
