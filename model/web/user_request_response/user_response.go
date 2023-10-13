package userrequestresponse

type UserResponse struct {
	Id    int    `json:"id"`
	NIP   string `json:"nip"`
	Rank  int    `json:"rank"`
	Email string `json:"email"`
}
