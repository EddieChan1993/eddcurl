package main

import (
	"testing"
	."github.com/EddieChan1993/eddcurl/curl"
	"log"
)

func Test(t *testing.T) {
	url:= "http://sms-api.luosimao.com/v1/send.json"

	headers:=map[string]string{
		"Content-Type":"application/x-www-form-urlencoded",
		"Authorization":BasicAuth("api","78aac6166f23182bd2eaceae0fba6aa84"),
	}
	postData:=map[string]string{
		"mobile":"18380591566",
		"message":"go-lang test【环球娃娃】",
	}

	req:=NewRequst(url)
	result:=req.
	SetHeaders(headers).
		SetPostData(postData).
		Post()

	log.Println(result)
}
