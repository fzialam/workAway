package tucontroller

import "net/http"

type TUController interface {
	Index(w http.ResponseWriter, r *http.Request)
	IndexSPPD(w http.ResponseWriter, r *http.Request)
	DetailSurat(w http.ResponseWriter, r *http.Request)
	CreateSPPD(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
}
