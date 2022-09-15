package main

import (
	"context"
	"fmt"
	"log"

	immuclient "github.com/codenotary/immudb/pkg/client"

	"google.golang.org/grpc/metadata"
)

func Run() error {
	client, err := immuclient.NewImmuClient(immuclient.DefaultOptions())
	if err != nil {
		log.Fatal(err)
		//return err
	}
	ctx := context.Background()
	lr, err := client.Login(ctx, []byte(`immudb`), []byte(`immudb`))
	if err != nil {
		log.Fatal(err)
		return err
	}

	md := metadata.Pairs("authorization", lr.Token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	vtx, err := client.VerifiedSet(ctx, []byte(`welcome`), []byte(`gophers`))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Successfully set key with transaction ID %d", vtx.Id)

	ventry, err := client.VerifiedSet(ctx, []byte(`welcome`), []byte(`gophers`))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Successfully retried and verified entry%v\n", ventry)
	return nil
}

func main() {
	fmt.Println("Dear friends, I'm looking for a warm home. ")
}
