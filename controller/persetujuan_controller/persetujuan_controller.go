package persetujuancontroller

import "net/http"

type PersetujunanController interface {
	Index(w http.ResponseWriter, r *http.Request)
	DetailSurat(w http.ResponseWriter, r *http.Request)
	SetApproved(w http.ResponseWriter, r *http.Request)
}
