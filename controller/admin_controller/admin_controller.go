package admincontroller

import "net/http"

type AdminController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Permohonan(w http.ResponseWriter, r *http.Request)
	Penugasan(w http.ResponseWriter, r *http.Request)
	LapAKK(w http.ResponseWriter, r *http.Request)
	LapAGG(w http.ResponseWriter, r *http.Request)
	UserGET(w http.ResponseWriter, r *http.Request)
	UserGETById(w http.ResponseWriter, r *http.Request)
	UserPUT(w http.ResponseWriter, r *http.Request)

	Profile(w http.ResponseWriter, r *http.Request)
}
