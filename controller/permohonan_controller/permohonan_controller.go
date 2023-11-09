package permohonancontroller

import (
	"net/http"
)

type PermohonanController interface {
	Index(w http.ResponseWriter, r *http.Request)
	CreatePermohonan(w http.ResponseWriter, r *http.Request)
}
