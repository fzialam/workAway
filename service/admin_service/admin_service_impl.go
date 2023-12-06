package adminservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
	adminrepository "github.com/fzialam/workAway/repository/admin_repository"
	"github.com/go-playground/validator/v10"
)

type AdminServiceImpl struct {
	AdminRepo adminrepository.AdminRepo
	DB        *sql.DB
	Validate  *validator.Validate
}

func NewAdminService(adminRepo adminrepository.AdminRepo, db *sql.DB, validate *validator.Validate) AdminService {
	return &AdminServiceImpl{
		AdminRepo: adminRepo,
		DB:        db,
		Validate:  validate,
	}
}

// Permohonan implements AdminService.
func (as *AdminServiceImpl) Permohonan(ctx context.Context) ([]surattugasreqres.SuratTugasJOINUserResponse, error) {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat, err := as.AdminRepo.Permohonan(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToSuratTugasJoinUserResponses(surat), nil
}

// Penugasan implements AdminService.
func (as *AdminServiceImpl) Penugasan(ctx context.Context) ([]surattugasreqres.SuratTugasJOINUserResponse, error) {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat, err := as.AdminRepo.Penugasan(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToSuratTugasJoinUserResponses(surat), nil
}

// LapAKK implements AdminService.
func (as *AdminServiceImpl) LapAKK(ctx context.Context) ([]entity.Laporan, error) {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	laporan, err := as.AdminRepo.LapAKK(ctx, tx)
	helper.PanicIfError(err)

	return laporan, nil
}

// LapAGG implements AdminService.
func (as *AdminServiceImpl) LapAGG(ctx context.Context) ([]entity.Laporan, error) {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	laporan, err := as.AdminRepo.LapAGG(ctx, tx)
	helper.PanicIfError(err)

	return laporan, nil
}

// UserGET implements AdminService.
func (as *AdminServiceImpl) UserGET(ctx context.Context) ([]userreqres.UserResponse, error) {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users, err := as.AdminRepo.UserGET(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToUserResponses(users), nil
}

// UserGETById implements AdminService.
func (as *AdminServiceImpl) UserGETById(ctx context.Context, userId int) (userreqres.UserResponse, error) {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := as.AdminRepo.UserGETById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user), nil
}

// UserPOST implements AdminService.
func (as *AdminServiceImpl) UserPOST(ctx context.Context, request userreqres.RankChangeRequest) error {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Id:   request.Id,
		Rank: request.Rank,
	}

	err = as.AdminRepo.UserPOST(ctx, tx, user)
	helper.PanicIfError(err)

	return nil
}

// Profile implements AdminService.
func (as *AdminServiceImpl) Profile(ctx context.Context, userId int) userreqres.UserResponse {
	tx, err := as.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := as.AdminRepo.Profile(ctx, tx, userId)

	return helper.ToUserResponse(user)
}
