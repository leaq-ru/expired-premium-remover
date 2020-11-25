package main

import (
	"context"
	"github.com/nnqq/scr-expired-premium-remover/call"
	"github.com/nnqq/scr-expired-premium-remover/config"
	"github.com/nnqq/scr-expired-premium-remover/healthz"
	"github.com/nnqq/scr-expired-premium-remover/logger"
	"github.com/nnqq/scr-expired-premium-remover/remover"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logg, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	companyClient, err := call.NewClients(cfg.Service.Parser)
	logg.Must(err)

	go func() {
		logg.Must(healthz.NewHealthz(logg.ZL, "80").Serve())
	}()

	logg.Must(remover.NewRemover(companyClient).Do(ctx))
}
