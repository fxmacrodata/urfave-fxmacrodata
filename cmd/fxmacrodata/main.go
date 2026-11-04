package main

import (
	"context"
	"fmt"
	fxmdcli "github.com/fxmacrodata/urfave-fxmacrodata"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := fxmdcli.Run(ctx, os.Args, os.Getenv, nil, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
