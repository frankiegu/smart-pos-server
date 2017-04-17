package main

import (
	"fmt"
	"net/http"
	"github.com/phonechan/smart-pos-server/router"
	"github.com/phonechan/smart-pos-server/middleware"
)

func main() {
	fmt.Println("====== app starting... ======")
	startApp()
}

func startApp() {

	http.Handle("/", middleware.RecoverMiddleware(router.RouterV1()))

	http.ListenAndServe(":8080", nil)
}