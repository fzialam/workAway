package userrequestresponse

type UserRegisterRequest struct {
	NIP      string `validate:"required" json:"nip"`
	Rank     int    `validate:"required" json:"rank"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required, alphanum, min=8" json:"password"`
}