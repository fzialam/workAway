package keuanganservice

import (
	"context"
	"database/sql"
	"log"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	approvedreqres "github.com/fzialam/workAway/model/req_res/approved_req_res"
	dokumenreqres "github.com/fzialam/workAway/model/req_res/dokumen_req_res"
	keuanganreqres "github.com/fzialam/workAway/model/req_res/keuangan_req_res"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	keuanganrepository "github.com/fzialam/workAway/repository/keuangan_repository"
	"github.com/go-playground/validator/v10"
)

type KeuanganServiceImpl struct {
	KeuanganRepo keuanganrepository.KeuanganRepo
	DB           *sql.DB
	Validate     validator.Validate
}

func NewKeuanganService(keuanganRepo keuanganrepository.KeuanganRepo, db *sql.DB, validate validator.Validate) KeuanganService {
	return &KeuanganServiceImpl{
		KeuanganRepo: keuanganRepo,
		DB:           db,
		Validate:     validate,
	}
}

// ListPermohonanApproved implements KeuanganService.
func (ks *KeuanganServiceImpl) ListPermohonanApproved(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse {
	tx, err := ks.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := ks.KeuanganRepo.ListSurat(ctx, tx)
	return helper.ToSuratTugasJOINApprovedUserResponses(surat)
}

// PermohonanApprovedById implements KeuanganService.
func (ks *KeuanganServiceImpl) PermohonanApprovedById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINApprovedUserParticipanLaporanResponse {
	tx, err := ks.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := ks.KeuanganRepo.GetSuratTugasById(ctx, tx, suratId)
	helper.PanicIfError(err)

	participans := ks.KeuanganRepo.GetAllParticipanJOINUserBySuratId(ctx, tx, suratId)

	surat.Participans = participans

	rincian := ks.KeuanganRepo.GetAllRincianBiayaBySuratId(ctx, tx, suratId)

	return helper.ToSuratTugasJOINApprovedUserParticipanLaporanResponse(surat, rincian)
}

// UploadRincianBiaya implements KeuanganService.
func (ks *KeuanganServiceImpl) UploadRincianBiaya(ctx context.Context, request keuanganreqres.UploadRincianAnggaranRequest) dokumenreqres.UploadDokumenResponse {
	err := ks.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ks.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rinci := entity.RincianAnggaran{
		SuratTugasId: request.SuratTugasId,
		DokName:      request.DokName,
		DokPDF:       request.DokPDF,
	}

	rinci = ks.KeuanganRepo.UploadRincianAnggaran(ctx, tx, rinci)

	err = ks.KeuanganRepo.AddNULLApprovedRincian(ctx, tx, rinci)
	helper.PanicIfError(err)

	return dokumenreqres.UploadDokumenResponse{
		DokLaporanName: rinci.DokName,
		Message:        "Succsess",
	}
}

// SetRincian implements KeuanganService.
func (ks *KeuanganServiceImpl) SetRincian(ctx context.Context, request keuanganreqres.UploadRincianAnggaranRequest) dokumenreqres.UploadDokumenResponse {
	err := ks.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ks.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rinci := entity.RincianAnggaran{
		SuratTugasId: request.SuratTugasId,
		DokName:      request.DokName,
		DokPDF:       request.DokPDF,
	}

	rinci = ks.KeuanganRepo.SetRincianAnggaran(ctx, tx, rinci)

	// GetIDRincianAnggaran
	rinci = ks.KeuanganRepo.GetIDRincianAnggaran(ctx, tx, rinci)

	err = ks.KeuanganRepo.SetNullApprovedRincianAnggaran(ctx, tx, rinci)
	helper.PanicIfError(err)

	return dokumenreqres.UploadDokumenResponse{
		DokLaporanName: rinci.DokName,
		Message:        "Succsess",
	}
}

// ListSPPDApproved implements KeuanganService.
func (ks *KeuanganServiceImpl) ListSPPDApproved(ctx context.Context) []keuanganreqres.SuratTugasJOINApprovedUserOtherIdResponse {
	tx, err := ks.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := ks.KeuanganRepo.ListSPPDIsApproved(ctx, tx)
	return helper.ToSuratTugasJOINApprovedUserOtherIdResponses(surat)
}

// SetFullAnggaran implements KeuanganService.
func (ks *KeuanganServiceImpl) SetFullAnggaran(ctx context.Context, request approvedreqres.ApprovedRequest) approvedreqres.ApprovedResponse {
	err := ks.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ks.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	approved := entity.Approved{
		Id:      request.Id,
		Status:  request.Status,
		Message: request.Message,
	}

	approved = ks.KeuanganRepo.SetFullAnggaran(ctx, tx, approved)

	return approvedreqres.ApprovedResponse{
		Status:  approved.Status,
		Message: "Success",
	}
}

// ListLaporanSPPD implements KeuanganService.
func (ks *KeuanganServiceImpl) ListLaporanSPPD(ctx context.Context) []surattugasreqres.SuratTugasJOINUserLaporanApprovedResponse {
	tx, err := ks.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surat := ks.KeuanganRepo.ListLaporanSPPD(ctx, tx)

	for i := range surat {
		laporan := ks.KeuanganRepo.LaporanBySPPDId(ctx, tx, surat[i].Id)

		isApproved := ks.KeuanganRepo.IsLaporanApproved(ctx, tx, laporan.Id)
		laporan.Status = isApproved.Status
		surat[i].Laporan = laporan
	}

	return helper.ToSuratTugasJOINUserLaporanApprovedResponses(surat)
}

// SetApprovedLaporan implements KeuanganService.
func (ks *KeuanganServiceImpl) SetApprovedLaporan(ctx context.Context, request laporanreqres.ApprovedLaporanRequest) laporanreqres.ApprovedLaporanResponse {
	err := ks.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ks.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	laporan := entity.ApprovedLaporan{
		LaporanId: request.Id,
		Status:    request.Status,
		Message:   request.Message,
	}

	log.Println(laporan)

	laporan = ks.KeuanganRepo.SetApprovedLaporan(ctx, tx, laporan)

	return laporanreqres.ApprovedLaporanResponse{
		Status:  laporan.Status,
		Message: "Success",
	}
}
