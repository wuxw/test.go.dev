package main

import (
	"fmt"
	"test.go.dev/curl/curl"
)

func main()  {
	br := curl.NewBrowser()
/*	byteAry , statusCode := br.Get("http://www.qqt2.ph/auth/signin")
	fmt.Println(string(byteAry[:]), statusCode)*/
	params := map[string]string {
		// "_token": "1ov4aRl6YizrNLssDpR7WJy14CynuejxhzNWu0fY",
		//"_random": "$2y$10$tnteo66tNKefK47CRpZueOjH/MZkNVZkHS0pjxFSV4Bbrn7rb9ccy_MZazE3",
		"username": "12112211",
		"password": "6131306562363963353334303639653037393963383039353464313361343065",
	}
	/*headers := map[string]string {
		"Referer" : "http://www.qqt2.ph/auth/signin",
	}*/
	byteAry2 := br.Post("http://www.qqt2.ph/auth/signin?type=mobile", params ,  map[string]string {
		"Referer" : "http://www.qqt2.ph/auth/signin",
		"X-Requested-With" : "XMLHttpRequest", //ajax请求
		"User-Agent" : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	} )
	fmt.Println(string(byteAry2[:]))
}
