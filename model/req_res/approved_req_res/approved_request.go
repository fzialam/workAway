package approvedreqres

type ApprovedRequest struct {
	Id      int    `validate:"required" json:"Id"`
	Status  string `validate:"required" json:"status"`
	Message string `validate:"required" json:"message"`
}
