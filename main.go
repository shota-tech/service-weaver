package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	ctx := context.Background()
	root := weaver.Init(ctx)

	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	reversed, err := reverser.Reverse(ctx, "!dlroW , olleH")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reversed)
}
