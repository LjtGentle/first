package main

import (
	"ljtTest/myProject/db"
	"ljtTest/myProject/model"
	"ljtTest/myProject/router"
	"log"
	"net/http"
	"time"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)
func main() {

	db.InitTable(&model.DemoOrder{})
	server01 := &http.Server{
		Addr:              ":8000",
		Handler:           router.Router01(),
		TLSConfig:         nil,
		ReadTimeout:       5*time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      10*time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	server02 := &http.Server{
		Addr:              ":9111",
		Handler:           router.Router02(),
		TLSConfig:         nil,
		ReadTimeout:       5*time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      10*time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	g.Go(func()error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})
	err := g.Wait()
	if err != nil {
		log.Fatal(err)
	}

}
