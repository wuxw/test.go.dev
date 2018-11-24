package main

import (
	"fmt"
	"net/http"

	"test.go.dev/gin-blog/models"
	"test.go.dev/gin-blog/pkg/logging"
	"test.go.dev/gin-blog/pkg/setting"
	"test.go.dev/gin-blog/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	//var fns []*models.Functionalities
	maps := make(map[string]interface{})
	fns, _ := models.GetAllFunctionalities(maps)
	fmt.Println(fns)
	//logging.Info(fns)
	json := models.GetTree(fns)
	//logging.Info(json)
	fmt.Println(json)

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
