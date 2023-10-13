package userrequestresponse

type UserLoginRequest struct {
	Id       int    `json:"id"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
