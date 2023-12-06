package userreqres

type UserRegisterRequest struct {
	NIP      string `validate:"required" json:"nip"`
	NIK      string `validate:"required" json:"nik"`
	NPWP     string `validate:"required" json:"npwp"`
	Name     string `validate:"required" json:"name"`
	Rank     int    `json:"rank"`
	NoTelp   string `validate:"required" json:"no_telp"`
	TglLahir string `validate:"required" json:"tgl_lahir"`
	Status   string `validate:"required" json:"status"`
	Gender   int    `validate:"required" json:"gender"`
	Alamat   string `validate:"required" json:"alamat"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
