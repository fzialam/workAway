package app

import (
	"github.com/fzialam/workAway/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	r := httprouter.New()

	// Define an endpoint to fetch all data
	r.GET("/", controller.Index)

	return r
}
