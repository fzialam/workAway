package keuangancontroller

import "net/http"

type KeuanganController interface {
	IndexPermohonan(w http.ResponseWriter, r *http.Request)
	UploadRincian(w http.ResponseWriter, r *http.Request)
	SetRincian(w http.ResponseWriter, r *http.Request)

	IndexSPPD(w http.ResponseWriter, r *http.Request)
	SetFullAnggaran(w http.ResponseWriter, r *http.Request)

	IndexLaporan(w http.ResponseWriter, r *http.Request)
	SetApprovedLaporan(w http.ResponseWriter, r *http.Request)
}
