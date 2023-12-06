package keuangancontroller

import "net/http"

type KeuanganController interface {
	Index(w http.ResponseWriter, r *http.Request)
	IndexPermohonan(w http.ResponseWriter, r *http.Request)
	UploadRincian(w http.ResponseWriter, r *http.Request)
	SetRincian(w http.ResponseWriter, r *http.Request)

	IndexSPPD(w http.ResponseWriter, r *http.Request)
	SetFullAnggaran(w http.ResponseWriter, r *http.Request)

	IndexLaporan(w http.ResponseWriter, r *http.Request)
	SetApprovedLaporan(w http.ResponseWriter, r *http.Request)

	Profile(w http.ResponseWriter, r *http.Request)
}
