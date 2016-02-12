package utils

import (
	"net/http/httptest"
	"net/http"
	"io"
	"fmt"
	"github.com/gorilla/mux"
)

func MakeRequest(method, urlStr string, body io.Reader, router *mux.Router) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, urlStr, body)

	if err != nil {
		fmt.Println(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	return w
}
