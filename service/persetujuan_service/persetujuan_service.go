package persetujuanservice

import (
	"context"

	persetujuanreqres "github.com/fzialam/workAway/model/req_res/persetujuan_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
)

type PersetujuanService interface {
	SetApproved(ctx context.Context, request persetujuanreqres.PersetujuanRequest) persetujuanreqres.PersetujuanResponse
	GetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	GetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINApprovedUserParticipanResponse
}
