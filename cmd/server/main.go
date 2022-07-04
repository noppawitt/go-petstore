package main

import (
	"context"
	"log"

	"github.com/noppawitt/go-petstore/server"
)

func main() {
	if err := server.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
