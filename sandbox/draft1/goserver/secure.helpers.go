package goserver

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func GenerateCertForLocalHost() error {

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Error al generar clave privada: ", err)
	}

	ip, err := GetLocalIP()
	if err != nil {
		log.Fatal("Error al obtener IP local: ", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Localhost"},
			CommonName:   "localhost",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 90),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{ip},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		log.Fatal("Error al crear certificado: ", err)
	}

	// Guardar certificado
	certOut, err := os.Create("cert.pem")
	if err != nil {
		log.Fatal("Error al crear cert.pem: ", err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// Guardar clave privada
	keyOut, err := os.Create("key.pem")
	if err != nil {
		log.Fatal("Error al crear key.pem: ", err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()

	log.Println("Certificados generados: cert.pem y key.pem")
	return nil
}

func GenerateCertForLocalHostIfNotExists() error {

	cernExists := func() bool {
		_, err := os.Stat("cert.pem")
		return !os.IsNotExist(err)
	}

	pemExists := func() bool {
		_, err := os.Stat("key.pem")
		return !os.IsNotExist(err)
	}

	if !cernExists() || !pemExists() {

		if err := GenerateCertForLocalHost(); err != nil {
			return err
		}
		log.Println("Certificados generados: cert.pem y key.pem")
	} else {
		log.Println("Los certificados ya existen: cert.pem y key.pem")
		return nil
	}

	log.Println("Certificados ya existen, no es necesario generarlos.")
	return nil
}
