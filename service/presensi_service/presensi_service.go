package presensiservice

import (
	"context"

	presensireqres "github.com/fzialam/workAway/model/req_res/presensi_req_res"
)

type PresensiService interface {
	PresensiFoto(ctx context.Context, request presensireqres.PresensiFotoRequest) presensireqres.PresensiFotoResponse
	GetSurat(ctx context.Context, request presensireqres.GetSuratForPresensiRequest) []presensireqres.GetSuratForPresensiResponse
}
