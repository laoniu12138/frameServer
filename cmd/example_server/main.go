package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/laoniu12138/frameServer/tools"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	tools.GoWithRecover(ctx, func(ctx context.Context) {})

	select {
	case <-sigs:
		cancel()
		time.Sleep(5 * time.Second)
	}
}
