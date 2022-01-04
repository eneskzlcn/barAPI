package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
func handleErr(t *testing.T,err error){
	if err != nil{
		t.Fatal(err)
	}
}
func TestIGetTimesFromPingPostSuccessfully(t *testing.T){
	type Request struct{
		Times int `json:"times"`
	}
	testRequest := Request{Times: 3}
	app:= fiber.New()

	app.Post("/ping", func(ctx *fiber.Ctx) error {
		body := Request{}
		if err := ctx.BodyParser(&body); err != nil{
			return err
		}
		return ctx.JSON(body)
	})

	testRequestAsByte,err := json.Marshal(testRequest)
	handleErr(t,err)

	req := httptest.NewRequest(http.MethodPost,"/ping", strings.NewReader(string(testRequestAsByte)))
	req.Header.Set(fiber.HeaderContentType,fiber.MIMEApplicationJSON)

	var resp *http.Response
	resp ,err = app.Test(req,1)
	handleErr(t,err)

	var body []byte
	body,err = ioutil.ReadAll(resp.Body)
	handleErr(t,err)
	resentRequest := Request{}
	err = json.Unmarshal(body,&resentRequest)
	handleErr(t,err)

	assert.Equalf(t, resentRequest.Times,testRequest.Times,"Sended times came correctly to the api")
}
func TestIProducePongsAsManyAsGivenTimes(t *testing.T){
	type Request struct{
		Times int `json:"times"`
	}
	testRequest := Request{Times: 3}
	app:= fiber.New()

	type PongsResponse struct{
		Pongs []string `json:"pongs"`
	}
	const PONG = "pong"
	constructPongResponseAmountOfGivenTimes := func(times int) PongsResponse{
		var pongsResponse = PongsResponse{}
		for i:=0 ; i<times; i++{
			pongsResponse.Pongs = append(pongsResponse.Pongs,PONG)
		}
		return pongsResponse
	}
	app.Post("/ping", func(ctx *fiber.Ctx) error {
		body := Request{}
		if err := ctx.BodyParser(&body); err != nil{
			return err
		}
		pongsResponse := constructPongResponseAmountOfGivenTimes(body.Times)
		return ctx.JSON(pongsResponse)
	})
	testRequestAsByte,err := json.Marshal(testRequest)
	handleErr(t,err)
	req := httptest.NewRequest(http.MethodPost,"/ping",strings.NewReader(string(testRequestAsByte)))
	req.Header.Set(fiber.HeaderContentType,fiber.MIMEApplicationJSON)

	var resp *http.Response
	resp,err = app.Test(req,1)
	handleErr(t,err)

	var body []byte
	body,err = ioutil.ReadAll(resp.Body)
	handleErr(t,err)
	sentPongsResponse := PongsResponse{}
	err = json.Unmarshal(body,&sentPongsResponse)
	handleErr(t,err)
	assert.Equalf(t, len(sentPongsResponse.Pongs),3,"Pongs response that contains 'pong' word amount of given count")
}