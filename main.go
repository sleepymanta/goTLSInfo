package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, _ := tls.Dial("tcp", "www.google.com:443", conf)
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Printf("Common Name: %s\n", cert.Subject.CommonName)
		for _, san := range cert.DNSNames {
			fmt.Printf("SAN: %v\n", san)
		}
		fmt.Printf("Serial Number: %s \n", cert.SerialNumber)
		fmt.Printf("Expires On: %s \n", cert.NotAfter.Format(time.RFC1123))
		fmt.Printf("Issuer Name: %s\n", cert.Issuer)
		fmt.Printf("Issuer CN: %s \n\n", cert.Issuer.CommonName)
	}
}
