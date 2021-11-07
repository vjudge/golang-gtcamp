package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main () {
	g, ctx := errgroup.WithContext(context.Background())
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	})
	serv := &http.Server{Addr: ":9000", Handler: handler}

	// 启动 server
	g.Go(func() error {
		fmt.Println("---start---")
		go func() {
			<-ctx.Done()
			fmt.Println("---done---")
			serv.Shutdown(ctx)
		}()
		return serv.ListenAndServe()
	})

	// Signal
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-sigs:
			return fmt.Errorf("termin signal: %v", sigs)
		}
	})
	err := g.Wait()
	fmt.Println(err)
}
