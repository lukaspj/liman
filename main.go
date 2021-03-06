// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/handlers"
)

func main() {
	cmd.CheckNotifications()

	log.Println("Listening http://0.0.0.0:8080")
	err := handlers.HTTPServer().ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
