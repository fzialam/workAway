package userrequestresponse

type UserResponse struct {
	Id    int    `json:"id"`
	NIP   string `json:"nip"`
	Email string `json:"email"`
}
