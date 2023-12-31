package tuservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	tureqres "github.com/fzialam/workAway/model/req_res/tu_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
	turepository "github.com/fzialam/workAway/repository/tu_repository"
	"github.com/go-playground/validator/v10"
)

type TUServiceImpl struct {
	TURepo   turepository.TURepo
	DB       *sql.DB
	Validate *validator.Validate
}

func NewTUService(tuRepo turepository.TURepo, db *sql.DB, validate *validator.Validate) TUService {
	return &TUServiceImpl{
		TURepo:   tuRepo,
		DB:       db,
		Validate: validate,
	}
}

// IndexTU implements TUService.
func (ts *TUServiceImpl) IndexTU(ctx context.Context) tureqres.IndexTU {
	tx, err := ts.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	index, err := ts.TURepo.Index(ctx, tx)
	helper.PanicIfError(err)

	return index
}

// CreateSPPD implements TUService.
func (ts *TUServiceImpl) CreateSPPD(ctx context.Context, request tureqres.CreateSPPDRequest) tureqres.CreateSPPDResponse {
	err := ts.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ts.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := entity.SuratTugas{
		Id:          request.SuratTugasId,
		DokumenName: request.DokumenName,
		DokumenPDF:  request.DokumenPDF,
	}

	sppd, err := ts.TURepo.CreateSPPD(ctx, tx, surat)
	helper.PanicIfError(err)

	err = ts.TURepo.SetNULLStatus(ctx, tx, sppd.Id)
	helper.PanicIfError(err)

	return tureqres.CreateSPPDResponse{
		DokumenName: sppd.DokumenName,
		Message:     "success",
	}

}

// GetAllSuratTugasJOINApprovedUser implements TUService.
func (ts *TUServiceImpl) GetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse {
	tx, err := ts.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	suratJA, err := ts.TURepo.ListSurat(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToSuratTugasJOINApprovedUserResponses(suratJA)
}

// GetSuratTugasById implements TUService.
func (ts *TUServiceImpl) GetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINDoubleApprovedUserParticipanResponse {
	tx, err := ts.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat, err := ts.TURepo.GetSuratTugasById(ctx, tx, suratId)
	helper.PanicIfError(err)

	participan, _ := ts.TURepo.GetAllParticipanJOINUserBySuratId(ctx, tx, suratId)
	surat.Participans = participan

	return helper.ToSuratTugasJOINDoubleApprovedUserParticipanResponse(surat)
}

// Profile implements TUService.
func (ts *TUServiceImpl) Profile(ctx context.Context, userId int) userreqres.UserResponse {
	tx, err := ts.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := ts.TURepo.Profile(ctx, tx, userId)

	return helper.ToUserResponse(user)
}
