package penugasancontroller

import (
	"net/http"
)

type PenugasanController interface {
	Index(w http.ResponseWriter, r *http.Request)
	CreatePenugasan(w http.ResponseWriter, r *http.Request)
	DetailSurat(w http.ResponseWriter, r *http.Request)
	SetApproved(w http.ResponseWriter, r *http.Request)
}
