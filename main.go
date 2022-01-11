package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rombintu/square_colony/game"
	"github.com/rombintu/square_colony/server"
)

func runServer() {
	server := server.NewServer("localhost", "5000")
	if err := server.Start(); err != nil {
		server.Logger.Fatalf("%v", err)
	}
	server.Logger.Info("Server exit")
}

func runGame() {
	game := game.NewGame()
	game.Init()
	game.Logger.Info("Game exit")
}

func programExit(ctx context.Context) {
	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh
	fmt.Println("Exit with 0")
	os.Exit(0)
	// ctx.Done()
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(3)
	defer wg.Done()
	go programExit(ctx)
	go runServer()
	go runGame()
	wg.Wait()
	cancelFunc()
}
