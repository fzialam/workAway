package helper

import (
	"html/template"
	"net/http"
)

func AddIndex(x int) int {
	return x + 1
}

func UnauthorizedTemplate(w http.ResponseWriter, r *http.Request, message interface{}) {
	temp, err := template.ParseFiles("view/unauthorized.html")

	data := map[string]interface{}{
		"message": message,
	}

	PanicIfError(err)

	err = temp.Execute(w, data)
	PanicIfError(err)
}
