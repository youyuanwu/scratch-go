package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/Microsoft/go-winio"
)

func main() {

	t := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			timeout := 10 * time.Second
			//println(network)
			//println(addr)
			return winio.DialPipe("\\\\.\\pipe\\mynamedpipe", &timeout)
		},
	}

	c := &http.Client{
		Transport: t,
	}

	r, err := http.NewRequest("GET",
		"http://someurl:80/count",
		nil)
	if err != nil {
		panic(err)
	}

	resp, err := c.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", dump)
}
