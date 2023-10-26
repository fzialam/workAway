package presensiservice

import (
	"context"
	"database/sql"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	presensireqres "github.com/fzialam/workAway/model/presensi_request_response"
	presensirepository "github.com/fzialam/workAway/repository/presensi_repository"
	"github.com/go-playground/validator/v10"
)

type PresensiServiceImpl struct {
	PresensiRepo presensirepository.PresensiRepo
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewPresensiService(presensiRepo presensirepository.PresensiRepo, db *sql.DB, validate *validator.Validate) PresensiService {
	return &PresensiServiceImpl{
		PresensiRepo: presensiRepo,
		DB:           db,
		Validate:     validate,
	}
}

// PresensiFoto implements PresensiService.
func (ps *PresensiServiceImpl) PresensiFoto(ctx context.Context, request presensireqres.PresensiFotoRequest) presensireqres.PresensiFotoResponse {
	err := ps.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ps.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = ps.PresensiRepo.CheckIzin(ctx, tx, request.SuratTugasId)
	helper.PanicIfError(err)

	presensi := entity.Presensi{
		UserId:       request.UserId,
		SuratTugasId: request.SuratTugasId,
		Gambar:       request.Gambar,
		Lokasi:       request.Lokasi,
	}

	presensi, err = ps.PresensiRepo.PresensiFoto(ctx, tx, presensi)
	helper.PanicIfError(err)

	return helper.ToPresensiResponse(presensi)
}