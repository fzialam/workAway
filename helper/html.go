package helper

import (
	"html/template"
	"net/http"
)

func AddIndex(x int) int {
	return x + 1
}

func UnauthorizedTemplate(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/unauthorized.html")

	PanicIfError(err)

	err = temp.Execute(w, nil)
	PanicIfError(err)
}
