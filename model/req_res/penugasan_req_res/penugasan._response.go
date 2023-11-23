package penugasanreqres

type PenugasanResponse struct {
	UserKetuaId   int    `validate:"required" json:"user_ketua_id"`
	LokasiTujuan  string `json:"lokasi_tujuan"`
	JenisProgram  string `json:"jenis_program"`
	ParticipansId []int  `json:"participans_id"`
	TglAwal       string `json:"tgl_awal"`
	TglAkhir      string `json:"tgl_akhir"`
}

type ApprovedResponse struct {
	Status  string ` json:"status"`
	Message string `json:"message"`
}
