package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	pb "github.com/bonefabric/vrabber-client-telegram/gen/client"
	"github.com/bonefabric/vrabber-client-telegram/internal/config"
	"github.com/bonefabric/vrabber-client-telegram/internal/telegram"
	"github.com/bonefabric/vrabber-client-telegram/internal/vrabber"
	"github.com/bonefabric/vrabber-client-telegram/setup"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := config.MustLoad()
	setup.ConfigureLogLevel(cfg.LogLevel)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	req, resp := make(chan *pb.StartDownloadRequest, cfg.MessagesBuffer), make(chan *pb.DownloadStatusResponse, cfg.MessagesBuffer)

	client, err := telegram.NewClient(ctx, cfg.TgToken, resp, req)
	if err != nil {
		slog.Error("failed to create telegram client", "err", err)
		return
	}

	cn := vrabber.NewClient(ctx, cfg.ServerHost, cfg.ServerPort, req, resp)

	eg := &errgroup.Group{}
	eg.Go(client.Start)
	eg.Go(cn.Start)

	if err = eg.Wait(); err != nil {
		slog.Error("application exited with error", "err", err)
	} else {
		slog.Info("application exited without errors")
	}
}
