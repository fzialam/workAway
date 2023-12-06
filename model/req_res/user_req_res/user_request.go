package userreqres

type ChangePasswordReq struct {
	Id          int    `validate:"required" json:"id"`
	Password    string `validate:"required" json:"password"`
	NewPassword string `validate:"required" json:"new_password"`
}
type ChangeImageReq struct {
	Id     int    `validate:"required" json:"id"`
	Gambar string `validate:"required" json:"gambar"`
}

type RankChangeRequest struct {
	Id   int `validate:"required" json:"id"`
	Rank int `validate:"required" json:"rank"`
}
