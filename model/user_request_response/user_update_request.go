package userrequestresponse

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	NIP      string `validate:"required" json:"nip"`
	Rank     int    `validate:"required" json:"rank"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
