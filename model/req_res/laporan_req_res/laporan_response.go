package laporanreqres

type UploadLaporanResponse struct {
	DokLaporanName string `json:"dok_laporan_name"`
	Message        string `json:"message"`
}

type ApprovedLaporanResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
