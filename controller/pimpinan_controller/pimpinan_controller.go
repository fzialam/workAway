package pimpinancontroller

import "net/http"

type PimpinanController interface {
	CreatePenugasan(w http.ResponseWriter, r *http.Request)

	IndexSPPD(w http.ResponseWriter, r *http.Request)
	SPPDDetailSurat(w http.ResponseWriter, r *http.Request)
	SPPDSetApproved(w http.ResponseWriter, r *http.Request)

	IndexPermohonan(w http.ResponseWriter, r *http.Request)
	PermohonanDetailSurat(w http.ResponseWriter, r *http.Request)
	PermohonanSetApproved(w http.ResponseWriter, r *http.Request)

	IndexLaporan(w http.ResponseWriter, r *http.Request)
	LaporanDetail(w http.ResponseWriter, r *http.Request)
	LaporanSetAprroved(w http.ResponseWriter, r *http.Request)
}
