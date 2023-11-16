package presensireqres

type GetSuratForPresensiRequest struct {
	UserId int `validate:"required" json:"user_id"`
}
