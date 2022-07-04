package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	pb "github.com/youyuanwu/scratch-go/grpc/helloworld/helloworld"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/proto"
)

func main() {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DumpRequest(r)
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		hellorequest := pb.HelloRequest{}
		err = proto.Unmarshal(reqBody[5:], &hellorequest)
		if err != nil {
			log.Fatal(err)
		}

		respBody := pb.HelloReply{Message: hellorequest.GetName()}
		msg, err := proto.Marshal(&respBody)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Trailer", "grpc-status")

		w.Header().Add("Content-Type", "application/grpc+proto")
		w.Header().Add("Connection", "close") // this header not get delivered to client
		w.WriteHeader(http.StatusOK)

		bs := make([]byte, 5)
		binary.BigEndian.PutUint32(bs[1:], uint32(len(msg)))
		//body := io.MultiReader(bytes.NewReader(bs), bytes.NewReader(msg))
		w.Write(bs)
		w.Write(msg)

		// Note that golang http dump will ignore this field.
		w.Header().Set("grpc-status", "0")
	})
	h2s := &http2.Server{
		// ...
	}
	h1s := &http.Server{
		Addr:    ":50051",
		Handler: h2c.NewHandler(handler, h2s),
	}
	log.Printf("server listening at %v", h1s.Addr)
	log.Fatal(h1s.ListenAndServe())
}

func DumpRequest(req *http.Request) {
	reqDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q", reqDump)
	//fmt.Print("\n\n")
}
