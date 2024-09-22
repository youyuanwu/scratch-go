package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/google/certtostore"
	"github.com/youyuanwu/scratch-go/pkg/wintls"
)

func main() {

	ctx, key, cert := wintls.GetFromStore()
	defer certtostore.FreeCertContext(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})

	certsnew := tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		Leaf:        cert,
		PrivateKey:  key,
	}

	cfg := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{certsnew},
		ClientAuth:   tls.NoClientCert,
	}
	srv := &http.Server{
		Addr:      ":12345",
		Handler:   mux,
		TLSConfig: cfg,
		// TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS("", ""))
}
