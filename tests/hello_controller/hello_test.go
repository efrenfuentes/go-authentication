package hello_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/efrenfuentes/go-authentication/core/settings"
	"github.com/efrenfuentes/go-authentication/routers"
	"os"
	"net/http"
	"net/http/httptest"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
)

var router *mux.Router

var _ = Describe("Hello Controller", func() {
	BeforeSuite(func() {
		os.Setenv("GO_ENV", "test")
		settings.Init()
		router = routers.InitRoutes()
	})

	It("Calling /api/v1/hello", func() {
		req, err := http.NewRequest("GET", "/api/v1/hello", nil)

		if err != nil {
			fmt.Println(err)
		}

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))

		var jsonResponse interface{}
		err = json.Unmarshal([]byte(w.Body.String()), &jsonResponse)
		if err != nil {
			fmt.Println(err)
		}
		responseMap := jsonResponse.(map[string]interface{})
		Expect(responseMap["message"]).To(Equal("Hello, World!"))
	})

	It("Calling /api/v1/hello/{name}", func() {
		req, err := http.NewRequest("GET", "/api/v1/hello/Jhon", nil)

		if err != nil {
			fmt.Println(err)
		}

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))

		var jsonResponse interface{}
		err = json.Unmarshal([]byte(w.Body.String()), &jsonResponse)
		if err != nil {
			fmt.Println(err)
		}
		responseMap := jsonResponse.(map[string]interface{})
		Expect(responseMap["message"]).To(Equal("Hello, Jhon"))
	})
})
