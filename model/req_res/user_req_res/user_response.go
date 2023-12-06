package userreqres

type UserResponse struct {
	Id       int    `json:"id"`
	NIP      string `json:"nip"`
	NIK      string `json:"nik"`
	NPWP     string `json:"npwp"`
	Name     string `json:"name"`
	Rank     int    `json:"rank"`
	NoTelp   string `json:"no_telp"`
	TglLahir string `json:"tgl_lahir"`
	Status   string `json:"status"`
	Gender   int    `json:"gender"`
	Alamat   string `json:"alamat"`
	Email    string `json:"email"`
	Gambar   string `json:"gambar"`
}

type LoginResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Rank  int    `json:"rank"`
	Token string `json:"token"`
}
