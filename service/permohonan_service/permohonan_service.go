package permohonanservice

import (
	"context"

	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
)

type PermohonanService interface {
	GetAllUserId(ctx context.Context) []userreqres.UserResponse
	CreatePermohonan(ctx context.Context, request permohonanreqres.PermohonanRequest) permohonanreqres.PermohonanResponse
}
