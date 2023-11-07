package userreqres

type UserRegisterRequest struct {
	NIP      string `validate:"required" json:"nip"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required, alphanum, min=8" json:"password"`
}
