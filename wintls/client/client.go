package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/google/certtostore"
	"github.com/youyuanwu/scratch-go/wintls"
)

func main() {
	ctx, key, cert := wintls.GetFromStore()
	defer certtostore.FreeCertContext(ctx)
	certsnew := tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key,
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(cert)
	t := &http.Transport{

		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            caCertPool,
			Certificates:       []tls.Certificate{certsnew}},
	}
	c := &http.Client{
		Transport: t,
	}

	r, err := http.NewRequest("GET",
		"https://localhost:12344",
		bytes.NewReader([]byte{}))

	if err != nil {
		fmt.Println(err)
	}
	r.Header.Set("Content-Type", "application/grpc+proto")
	// r.Header.Set("grpc-encoding", "gzip")
	r.Header.Set("User-Agent", "grpc-go-raw-http2")

	resp, err := c.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", dump)
}
