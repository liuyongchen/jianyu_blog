package main

import (
	"net/http"
	"time"

	"blog-service/internal/routers"
)

func main()  {
	router := routers.NewRouter()

	s := http.Server{
		Addr:              "：8080",
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	s.ListenAndServe()
}
