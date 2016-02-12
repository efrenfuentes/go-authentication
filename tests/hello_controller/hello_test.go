package hello_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/efrenfuentes/go-authentication/core/settings"
	"github.com/efrenfuentes/go-authentication/routers"
	"github.com/efrenfuentes/go-authentication/core/utils"
	"os"
	"net/http"

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
		var jsonResponse interface{}

		w := utils.MakeRequest("GET", "/api/v1/hello", nil, router)

		Expect(w.Code).To(Equal(http.StatusOK))

		err := json.Unmarshal([]byte(w.Body.String()), &jsonResponse)

		if err != nil {
			fmt.Println(err)
		}

		responseMap := jsonResponse.(map[string]interface{})
		Expect(responseMap["message"]).To(Equal("Hello, World!"))
	})

	It("Calling /api/v1/hello/{name}", func() {
		var jsonResponse interface{}

		w := utils.MakeRequest("GET", "/api/v1/hello/Jhon", nil, router)

		Expect(w.Code).To(Equal(http.StatusOK))

		err := json.Unmarshal([]byte(w.Body.String()), &jsonResponse)

		if err != nil {
			fmt.Println(err)
		}

		responseMap := jsonResponse.(map[string]interface{})
		Expect(responseMap["message"]).To(Equal("Hello, Jhon"))
	})
})
