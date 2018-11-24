package main

import (
	"fmt"
	"test.go.dev/curl/curl"
)

func main()  {
	br := curl.NewBrowser()
/*	byteAry , statusCode := br.Get("http://www.baidu.com")
	fmt.Println(string(byteAry[:]), statusCode)*/
	params := map[string]string {
		"username": "ww12112211",
		"password": "aE3DEgMzC43Nm1F2E4vsu1/pHp+CMzFtnR3ac7w89XQpHTVaTqo0cnUV0yV/3qGhjXi/oKI3/S/kQQ1GVSnXgfZLAaarS5jN2eL5ELfNuNjBdJ5Ij04aO30dmaZYjTHyhOQ5ILSDmMRactUJRCWy/oDfQ3NnvbZSTvZGwf8k39k=",
	}
	/*headers := map[string]string {
		"Referer" : "http://www.qqt2.ph/auth/signin",
	}*/
	byteAry2 := br.Post("https://passport.baidu.com/v2/api/?login", params ,  map[string]string {
		"Referer" : "https://www.baidu.com",
		"X-Requested-With" : "XMLHttpRequest", //ajax请求
		"User-Agent" : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	} )
	fmt.Println(string(byteAry2[:]))
}
