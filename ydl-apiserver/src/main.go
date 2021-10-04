package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type request struct {
	URL  string `json:"url"`
	FMT  string `json:"fmt"`
	Type string `json:"type"`
	Crf  string `json:"crf"`
}
type response struct {
	Tasks int    `json:"tasks"`
	Type  string `json:"type"`
}
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var dl_task = []request{}
var cnv_task = []request{}

func main() {
	e := echo.New() //echoを定義

	e.GET("/request", get_request)
	e.POST("/request", get_request)
	e.GET("/get", get_stored_tasks)
	e.GET("/pop", pop_stored_request)
	if err := e.Start(":6001"); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
func get_request(c echo.Context) error {
	res := &request{FMT: "mp4", Crf: "16", Type: "dl"}
	if err := c.Bind(res); err != nil {
		return err
	}
	switch res.Type {
	case "cnv":
		cnv_task = append(cnv_task, *res)
	case "dl":
		dl_task = append(dl_task, *res)
	default:
		var apierr APIError
		apierr.Code = 100
		apierr.Message = "invalid task request: invalid type"
		return c.JSON(http.StatusBadRequest, apierr)
	}
	log.Println("[" +
		fmt.Sprintf("%x", md5.Sum([]byte(res.URL+res.FMT))) + "]" +
		" " + res.Type + " task Received: " +
		"URL=" + res.URL + " CRF=" + res.Crf + " FMT=" + res.FMT)
	return c.JSON(http.StatusOK, res)
}

func get_stored_tasks(c echo.Context) error {
	res := &response{}
	if err := c.Bind(res); err != nil {
		return err
	}
	switch res.Type {
	case "dl":
		res.Tasks = len(dl_task)
	case "cnv":
		res.Tasks = len(cnv_task)
	default:
		var apierr APIError
		apierr.Code = 110
		apierr.Message = "invalid get request: invalid type"
		return c.JSON(http.StatusBadRequest, apierr)
	}
	return c.JSON(http.StatusOK, res)
}

func pop_stored_request(c echo.Context) error {
	res := &response{}
	var task request
	if err := c.Bind(res); err != nil {
		return err
	}
	switch res.Type {
	case "dl":
		if len(dl_task) >= 1 {
			task = dl_task[0]
			dl_task = dl_task[1:]
		}
	case "cnv":
		if len(cnv_task) >= 1 {
			task = cnv_task[0]
			cnv_task = cnv_task[1:]
		}
	}
	if task.URL == "" {
		var apierr APIError
		apierr.Code = 120
		apierr.Message = "invalid pop request: URL param is empty"
		return c.JSON(http.StatusBadRequest, apierr)
	}
	log.Println("[" +
		fmt.Sprintf("%x", md5.Sum([]byte(task.URL+task.FMT))) + "]" +
		" " + res.Type + " Process start: " +
		"URL=" + task.URL + " CRF=" + task.Crf + " FMT=" + task.FMT)
	return c.JSON(http.StatusOK, task)
}
