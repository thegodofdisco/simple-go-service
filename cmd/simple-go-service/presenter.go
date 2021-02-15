package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Presenter struct {
}

func NewPresenter() *Presenter {
	return &Presenter{}
}

type SuccessfulResponse struct {
	Message Message `json:"success"`
}

type Message struct {
	Message string `json:"message"`
}

type ErrorRespose struct {
	Message Message `json:"error"`
}

func NewBadRequestResponse() ErrorRespose {
	return ErrorRespose{
		Message: Message{http.StatusText(http.StatusBadRequest)},
	}
}

func NewSuccessfulResponse() SuccessfulResponse {
	return SuccessfulResponse{
		Message: Message{http.StatusText(http.StatusOK)},
	}
}

func (p *Presenter) GetData(ctx *gin.Context) {
	lastModifiedHeader := ctx.GetHeader("If-Modified-Since")
	lastModified, err := http.ParseTime(lastModifiedHeader)
	if err != nil {
		logrus.Error("Problem parsing If-Modified-Since header ", err)
		ctx.JSON(http.StatusBadRequest, NewBadRequestResponse())
		return
	}
	fmt.Println(lastModified)

	ctx.Header("Last-Modified", time.Now().String())
	ctx.JSON(http.StatusOK, NewSuccessfulResponse())
}
