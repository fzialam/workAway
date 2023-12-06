package tuservice

import (
	"context"

	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	tureqres "github.com/fzialam/workAway/model/req_res/tu_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type TUService interface {
	IndexTU(ctx context.Context) tureqres.IndexTU
	CreateSPPD(ctx context.Context, request tureqres.CreateSPPDRequest) tureqres.CreateSPPDResponse
	GetAllSuratTugasJOINApprovedUser(ctx context.Context) []surattugasreqres.SuratTugasJOINApprovedUserResponse
	GetSuratTugasById(ctx context.Context, suratId int) surattugasreqres.SuratTugasJOINDoubleApprovedUserParticipanResponse
	Profile(ctx context.Context, userId int) userreqres.UserResponse
}
