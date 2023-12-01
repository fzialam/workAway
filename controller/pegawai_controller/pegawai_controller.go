package pegawaicontroller

import "net/http"

type PegawaiController interface {
	Index(w http.ResponseWriter, r *http.Request)
	IndexPermohonan(w http.ResponseWriter, r *http.Request)
	CreatePermohonan(w http.ResponseWriter, r *http.Request)

	Presensi(w http.ResponseWriter, r *http.Request)
	GetSurat(w http.ResponseWriter, r *http.Request)

	IndexSPPD(w http.ResponseWriter, r *http.Request)

	IndexLaporan(w http.ResponseWriter, r *http.Request)
	UploadLapAktivitas(w http.ResponseWriter, r *http.Request)
	UploadLapAnggaran(w http.ResponseWriter, r *http.Request)

	SetLapAktivitas(w http.ResponseWriter, r *http.Request)
	SetLapAnggaran(w http.ResponseWriter, r *http.Request)
}
