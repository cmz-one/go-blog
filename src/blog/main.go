package main

import (
	"net/http"
	"fmt"
	"blog/pkg/setting"
	"blog/routers"
	"blog/models"
)

func main() {

	router := routers.InitRouter()

	defer models.CloseDB()
	s := &http.Server{
		Addr:fmt.Sprintf(":%d",setting.HTTPPort),
		Handler:router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}
	s.ListenAndServe()

}
