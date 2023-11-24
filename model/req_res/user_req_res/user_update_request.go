package userreqres

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Name     string `validate:"required" json:"name"`
	NoTelp   string `validate:"required" json:"no_telp"`
	TglLahir string `validate:"required" json:"tgl_lahir"`
	Alamat   string `validate:"required" json:"alamat"`
}
