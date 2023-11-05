package presensiservice

import (
	"context"

	presensireqres "github.com/fzialam/workAway/model/presensi_request_response"
)

type PresensiService interface {
	PresensiFoto(ctx context.Context, request presensireqres.PresensiFotoRequest) presensireqres.PresensiFotoResponse
	GetSurat(ctx context.Context, request presensireqres.GetSuratForPresensiRequest) []presensireqres.GetSuratForPresensiResponse
}
