package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cmd := ""
	if len(os.Args) > 0 {
		cmd = os.Args[1]
	}
	if cmd == "" {
		log.Fatal("no command given")
	}

	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	err := travisWait(contextWithSigterm(context.Background()), cmd, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func travisWait(ctx context.Context, command string, args ...string) error {
	start := time.Now()
	timeout := 20 * time.Minute
	interval := time.Minute

	ticker := time.NewTicker(interval)
	go func() {
		for t := range ticker.C {
			fmt.Printf("travis-wait waiting %s elapsed %s\n", t.Format(time.RFC1123Z), time.Since(start).Round(time.Second))
		}
	}()
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func contextWithSigterm(ctx context.Context) context.Context {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()

		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

		select {
		case <-signalCh:
		case <-ctx.Done():
		}
	}()

	return ctxWithCancel
}
