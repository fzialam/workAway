package userreqres

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	NIP      string `validate:"required" json:"nip"`
	NIK      string `validate:"required" json:"nik"`
	NPWP     string `validate:"required" json:"npwp"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	NoTelp   string `validate:"required" json:"no_telp"`
	TglLahir string `validate:"required" json:"tgl_lahir"`
	Alamat   string `validate:"required" json:"alamat"`
}
