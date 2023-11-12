package persetujuanservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	persetujuanreqres "github.com/fzialam/workAway/model/req_res/persetujuan_req_res"
	persetujuanrepository "github.com/fzialam/workAway/repository/persetujuan_repository"
	"github.com/go-playground/validator/v10"
)

type PersetujuanServiceImpl struct {
	PersetujuanRepo persetujuanrepository.PersetujuanRepo
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewPersetujuanService(persetujuanRepo persetujuanrepository.PersetujuanRepo, db *sql.DB, validate *validator.Validate) PersetujuanService {
	return &PersetujuanServiceImpl{
		PersetujuanRepo: persetujuanRepo,
		DB:              db,
		Validate:        validate,
	}
}

// SetApproved implements PersetujuanService.
func (ps *PersetujuanServiceImpl) SetApproved(ctx context.Context, request persetujuanreqres.PersetujuanRequest) persetujuanreqres.PersetujuanResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	izin := entity.Izin{
		SuratTugasId: request.SuratTugasId,
		Status:       request.Status,
		CreateAt:     request.CreateAt,
	}

	izin = ps.PersetujuanRepo.SetApproved(ctx, tx, izin)

	return helper.ToPersetujuanResponses(izin)
}

// GetAllSuratTugasJOINApprovedPemohonUser implements PersetujuanService.
func (ps *PersetujuanServiceImpl) GetAllSuratTugasJOINApprovedUser(ctx context.Context) []persetujuanreqres.SuratTugasJOINApprovedUserResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := ps.PersetujuanRepo.GetAllSuratTugasJOINApprovedUser(ctx, tx)
	helper.PanicIfError(err)
	return helper.ToSuratTugasJOINApprovedUserResponses(result)
}

// GetSuratTugasById implements PersetujuanService.
func (ps *PersetujuanServiceImpl) GetSuratTugasById(ctx context.Context, suratId int) persetujuanreqres.SuratTugasJOINApprovedUserParticipanResponse {
	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := ps.PersetujuanRepo.GetSuratTugasById(ctx, tx, suratId)
	helper.PanicIfError(err)

	participans, err := ps.PersetujuanRepo.GetAllParticipanJOINUserBySuratId(ctx, tx, suratId)
	helper.PanicIfError(err)

	result.Participans = participans

	return helper.ToSuratTugasJOINApprovedUserParticipanResponse(result)
}