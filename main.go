package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/y0k0ta19/go-theta/theta"
)

func main() {
	cc := github.NewClient(nil)
	cc.NewRequest("", "", nil)
	c := theta.NewClient(nil)
	ctx, cancel := context.WithCancel(context.Background())
	err := theta.Begin(ctx, c)
	if err != nil {
		fmt.Println(err)
	}
	cancel()
}
