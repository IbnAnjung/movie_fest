package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IbnAnjung/movie_fest/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	terminalHandler := make(chan os.Signal, 1)
	signal.Notify(
		terminalHandler,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	cleanup, err := app.Start(ctx)
	defer cleanup()

	if err != nil {
		panic(err)
	}

	go func() {
		in := <-terminalHandler
		log.Printf("SYSTEM CALL: %+v", in)
		cancel()
	}()

	<-ctx.Done()

}
