package goserver

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type Run interface {
	Start() error
	StartSecure() error
	Stop() error
}

var _ Run = (*GoServer)(nil)

// StartSecure implements Run.
func (g *GoServer) StartSecure() error {
	if g.server == nil {
		if err := g.BuildServer(); err != nil {
			return fmt.Errorf("failed to build server: %w", err)
		}
	}
	g.secure = true
	g.Log()
	return g.server.ListenAndServeTLS("cert.pem", "key.pem")
}

// Start implements Run.
func (g *GoServer) Start() error {

	if g.server == nil {
		if err := g.BuildServer(); err != nil {
			return fmt.Errorf("failed to build server: %w", err)
		}
	}
	g.secure = false

	g.Log() // Log server details before starting
	return g.server.ListenAndServe()
}

// Stop implements Run.
func (g *GoServer) Stop() error {
	panic("unimplemented")
}

func (g *GoServer) BuildServer() error {

	localip, err := GetLocalIP()
	if err != nil {
		return fmt.Errorf("failed to get local IP: %w", err)
	}

	g.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", localip, g.port),
		Handler:      g.mux,
		ReadTimeout:  g.readTimeout,
		WriteTimeout: g.writeTimeout,
		IdleTimeout:  g.idleTimeout,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		},
	}
	if g.server == nil {
		return fmt.Errorf("failed to create server: server is nil")
	}
	if g.mux == nil {
		return fmt.Errorf("failed to create server: mux is nil")
	}

	return nil
}

func (g *GoServer) Log() {

	proto := "http"
	if g.secure {
		proto = "https"
	}

	ip, err := GetLocalIP()

	if err != nil {
		fmt.Printf("Error getting local IP: %v\n", err)
		ip = "localhost"
	}

	fmt.Printf("GoServer is running on %s://%s:%d\n", proto, ip, g.port)
	fmt.Printf("GoServer is running on %s:localhost:%d\n", proto, g.port)
	fmt.Printf("Read Timeout: %s, Write Timeout: %s, Idle Timeout: %s\n",
		g.readTimeout, g.writeTimeout, g.idleTimeout)
}

