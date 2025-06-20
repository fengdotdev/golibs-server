package main

import (
	"net/http"

	"github.com/fengdotdev/golibs-server/sandbox/draft1/goserver"
)

func main() {

	err := goserver.GenerateCertForLocalHostIfNotExists()
	if err != nil {
		panic("Failed cert or pem" + err.Error())
	}

	myServer := goserver.NewGoServer()

	myServer.RegisterHandler("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Go Server!"))
	})

	myServer.RegisterHandler("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	myServer.StartSecure()
}
