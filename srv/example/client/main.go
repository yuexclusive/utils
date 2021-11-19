package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/rpc/client"
	"github.com/yuexclusive/utils/srv/auth/proto/auth"
	"github.com/yuexclusive/utils/srv/basic/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func main() {

	cfg := config.MustGet()

	closer, conn, err := client.Dial(cfg.AuthServiceName, "")

	defer closer.Close()

	if err != nil {
		log.Fatal(err)
	}

	// client := hello.NewHelloClient(conn)
	// res, err := client.Send(context.Background(), &hello.Request{Name: "somebody"})
	ac := auth.NewAuthClient(conn)

	r1, err := ac.Auth(context.Background(), &auth.AuthRequest{Id: "super_admin", Key: "123"})
	if err != nil {
		log.Fatal(err)
	}

	closer, conn2, err := client.Dial("srv.basic", r1.Token)

	defer closer.Close()

	if err != nil {
		log.Fatal(err)
	}

	client := user.NewUserClient(conn2)
	res, err := client.Get(context.Background(), &user.GetRequest{Name: "super_admin"}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
