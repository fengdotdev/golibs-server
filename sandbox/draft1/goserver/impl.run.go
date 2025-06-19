package goserver

import "net/http"

type Run interface {
	Start() error
	StartSecure() error
	Stop() error
}

var _ Run = (*GoServer)(nil)

// StartSecure implements Run.
func (g *GoServer) StartSecure() error {
	return http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
}

// Start implements Run.
func (g *GoServer) Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	return http.ListenAndServe(":8080", nil)
}

// Stop implements Run.
func (g *GoServer) Stop() error {
	panic("unimplemented")
}
