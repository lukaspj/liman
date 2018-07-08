package handlers

import (
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
)

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	err := parseSessionCookie(w, r)
	if err != nil {
		return
	}

	v, err := cmd.ParseVolumes()

	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	bn, _ := cmd.GetNotification()

	var data []interface{}
	data = append(data, bn)
	data = append(data, v)

	err = tpl.ExecuteTemplate(w, "volumes.tmpl", data)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
	}
	log.Println(r.Method, r.URL.Path)
}