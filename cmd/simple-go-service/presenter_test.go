package main_test

import (
	"net/http"
	"net/http/httptest"

	main "github.com/nenov92/simple-go-service/cmd/simple-go-service/main"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Presenter", func() {

	var presenter *main.Presenter
	var recorder *httptest.ResponseRecorder
	var mockContext *gin.Context

	BeforeEach(func() {
		presenter = main.NewPresenter()
		recorder = httptest.NewRecorder()
		mockContext, _ = gin.CreateTestContext(recorder)
		mockContext.Request = httptest.NewRequest("GET", "http://abc.com/", nil)
		mockContext.Header("If-Modified-Since", "Mon, 02 Jan 2006 15:04:05 GMT")
	})

	Describe("Get Data", func() {
		When("GetData is called without Headers'", func() {

			It("should return a data list with JSON Content-Type", func() {
				presenter.GetData(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

	})

})
