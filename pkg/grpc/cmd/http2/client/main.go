// send http2 request to grpc server

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/golang/protobuf/proto"
	pb "github.com/youyuanwu/scratch-go/pkg/grpc/helloworld"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	// http2 by default use tls.
	// so we need h2c to by pass tls.
	// see https://github.com/golang/go/issues/14141

	t := &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(netw, addr)
		},
	}
	c := &http.Client{
		Transport: t,
	}

	payload := pb.HelloRequest{Name: "hello"}
	msg, err := proto.Marshal(&payload)
	if err != nil {
		fmt.Println(err)
	}

	// 1 byte compressiong flag, 4 byte length prefix
	bs := make([]byte, 5)
	binary.BigEndian.PutUint32(bs[1:], uint32(len(msg)))
	body := io.MultiReader(bytes.NewReader(bs), bytes.NewReader(msg))

	r, err := http.NewRequest("POST",
		"http://localhost:50051/helloworld.Greeter/SayHello",
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
	defer resp.Body.Close()

	DumpHttpResp(resp)
	//HandleRespHello(resp)
}

func HandleRespHello(resp *http.Response) {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if status := resp.Header.Get("grpc-status"); status != "" {
		var c int
		if c, err = strconv.Atoi(status); err != nil {
			log.Fatal(err)
		}
		err = grpc.Errorf(codes.Code(c), resp.Header.Get("grpc-message"))
		//status.Errorf(codes.Code(c), resp.Header.Get("grpc-message"))
		if err != nil {
			log.Fatal(err)
		}
	}

	helloOut := pb.HelloReply{}
	// 1 byte compressiong flag, 4 byte length prefix
	// ideally we should inspect those 5 bytes.
	err = proto.Unmarshal(respBody[5:], &helloOut)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s", helloOut.GetMessage())
}

func DumpHttpResp(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q", dump)
}
