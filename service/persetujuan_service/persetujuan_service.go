package persetujuanservice

import (
	"context"

	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
)

type PersetujuanService interface {
	SetApproved(ctx context.Context, request izinreqres.IzinRequest) izinreqres.IzinResponse
	GetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	GetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINApprovedUserParticipanResponse
}
