package main

import (
	"bufio"
	"fmt"
	"time"

	"github.com/Microsoft/go-winio"
)

func main() {
	timeout := 10 * time.Second
	conn, err := winio.DialPipe("\\\\.\\pipe\\mynamedpipe", &timeout)
	if err != nil {
		panic(err)
	}

	path := "/count"

	fmt.Fprintf(conn, "GET %s HTTP/1.0\r\n\r\n", path)
	reader := bufio.NewReader(conn)
	//res, err := ioutil.ReadAll(reader)

	// read until error happens.
	// for http request one needs to read until content-length is reached
	for {
		res, myerr := reader.ReadString('\n')
		fmt.Print(res)
		if myerr != nil {
			break
		}
	}
}
