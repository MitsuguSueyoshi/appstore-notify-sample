package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/appstore-notify-sample/pkg/cmd/api/registry"
)

func Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

	server, err := registry.NewServer()
	if err != nil {
		log.Panicf("failed to create server: %+v", err)
	}

	group, egctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
		return nil
	})
	<-egctx.Done()

	tctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(tctx); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	if err := group.Wait(); err != nil {
		log.Printf("failed to serve: %+v", err)
	}
}
