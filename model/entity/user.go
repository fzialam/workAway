package entity

type User struct {
	Id       int    `json:"id"`
	NIP      string `json:"nip"`
	Rank     int    `json:"rank"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
