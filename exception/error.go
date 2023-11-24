package exception

import (
	"net/http"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	"github.com/go-playground/validator/v10"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Tangani panic dengan memanggil ErrorHandler
				ErrorHandler(w, r, err)
			}
		}()

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if unauthorizedError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	if duplicatedError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(Unauthorized)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		response := model.Response{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := model.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := model.Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func duplicatedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(DuplicatedError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := model.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := model.Response{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, response)
}
