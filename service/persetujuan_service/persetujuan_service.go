package persetujuanservice

import (
	"context"

	persetujuanreqres "github.com/fzialam/workAway/model/req_res/persetujuan_req_res"
)

type PersetujuanService interface {
	SetApproved(ctx context.Context, request persetujuanreqres.PersetujuanRequest) persetujuanreqres.PersetujuanResponse
	GetAllSuratTugasJOINApprovedUser(ctx context.Context) []persetujuanreqres.SuratTugasJOINApprovedUserResponse
	GetSuratTugasById(ctx context.Context, suratId int) persetujuanreqres.SuratTugasJOINApprovedUserParticipanResponse
}
