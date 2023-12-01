package surattugasreqres

type GetSuratRequest struct {
	UserId int    `validate:"required" json:"user_id"`
	Tipe   string `validate:"required" json:"tipe"`
}
