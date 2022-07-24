// send http2 request to grpc server

package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

// inspect trailers.
func main() {

	t := &http2.Transport{
		AllowHTTP:       true,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: t,
	}

	// 1 byte compressiong flag, 4 byte length prefix
	body := bytes.NewBufferString("Hello")

	r, err := http.NewRequest("POST",
		"https://localhost:12356/winhttpapitest/",
		body)

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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Body is %s", string(respBody))

	for key, val := range resp.Trailer {
		fmt.Println(key, val)
	}

	resp.Body.Close()
}
