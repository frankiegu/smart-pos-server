package handler

import (
	"net/http"
	"github.com/CardInfoLink/log"
)

func Activate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	log.Debug("activate")
}
