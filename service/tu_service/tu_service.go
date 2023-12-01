package tuservice

import (
	"context"

	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	tureqres "github.com/fzialam/workAway/model/req_res/tu_req_res"
)

type TUService interface {
	CreateSPPD(ctx context.Context, request tureqres.CreateSPPDRequest) tureqres.CreateSPPDResponse
	GetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	GetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINDoubleApprovedUserParticipanResponse
}
