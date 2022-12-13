package main

import (
	"context"
	"fmt"

	immuclient "github.com/codenotary/immudb/pkg/client"
	"google.golang.org/grpc/metadata"
)

func Run() {
	client, _ := immuclient.NewImmuClient(immuclient.DefaultOptions())

	ctx := context.Background()
	lr, _ := client.Login(ctx, []byte(`immudb`), []byte(`immudb`))

	md := metadata.Pairs("authorization", lr.Token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	vtx, _ := client.VerifiedSet(ctx, []byte(`welcome`), []byte(`gophers`))

	fmt.Printf("Succesfully set key: %d\n", vtx.Id)

	ventry, _ := client.VerifiedGet(ctx, []byte(`welcome`))

	fmt.Printf("Succesful entry \n%v\n", ventry)
}

func main() {
	Run()
}
