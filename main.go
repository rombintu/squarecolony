package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rombintu/square_colony/game"
)

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
	wg.Add(2)
	defer wg.Done()
	go programExit(ctx)
	go runGame()
	wg.Wait()
	cancelFunc()
}
