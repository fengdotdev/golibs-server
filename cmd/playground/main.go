package main

import (
	"net/http"

	"github.com/fengdotdev/golibs-server/sandbox/draft1/gogroup"
	"github.com/fengdotdev/golibs-server/sandbox/draft1/gomiddlewares"
	"github.com/fengdotdev/golibs-server/sandbox/draft1/goserver"
)

func main() {
	myServer := goserver.NewGoServer()

	authMiddleware := gomiddlewares.NewGoMiddleware("AuthMiddleware", func(w http.ResponseWriter, r *http.Request) {
		// Example authentication logic
		token := r.Header.Get("Authorization")
		if token != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	})

	mygroup := gogroup.NewGoGroup("MyGroup")

	// Start the server
	if err := myServer.Start(); err != nil {
		panic(err)
	}

}
