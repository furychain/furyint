package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	grpcda "github.com/xblackfury/furyint/da/grpc"
	"github.com/xblackfury/furyint/da/grpc/mockserv"
	"github.com/xblackfury/furyint/store"
)

func main() {
	conf := grpcda.DefaultConfig

	flag.IntVar(&conf.Port, "port", conf.Port, "listening port")
	flag.StringVar(&conf.Host, "host", "0.0.0.0", "listening address")
	flag.Parse()

	kv := store.NewDefaultKVStore(".", "db", "furyint")
	lis, err := net.Listen("tcp", conf.Host+":"+strconv.Itoa(conf.Port))
	if err != nil {
		log.Panic(err)
	}
	log.Println("Listening on:", lis.Addr())
	srv := mockserv.GetServer(kv, conf, nil)
	if err := srv.Serve(lis); err != nil {
		log.Println("error while serving:", err)
	}
}
