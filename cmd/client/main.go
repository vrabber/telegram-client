package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	pb "github.com/vrabber/telegram-client/gen/client"
	"github.com/vrabber/telegram-client/internal/config"
	"github.com/vrabber/telegram-client/internal/telegram"
	"github.com/vrabber/telegram-client/internal/vrabber"
	"github.com/vrabber/telegram-client/setup"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := config.MustLoad()
	setup.ConfigureLogLevel(cfg.LogLevel)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	req := make(chan *pb.StartDownloadRequest, cfg.MessagesBuffer)
	resp := make(chan *pb.DownloadStatusResponse, cfg.MessagesBuffer)

	client, err := telegram.NewClient(ctx, cfg.TgToken, resp, req)
	if err != nil {
		slog.Error("failed to create telegram client", "err", err)
		panic(err)
	}

	cn := vrabber.NewClient(ctx, cfg.ServerHost, cfg.ServerPort, req, resp)
	if err = client.Setup(); err != nil {
		slog.Error("failed to setup telegram client", "err", err)
		panic(err)
	}

	eg := &errgroup.Group{}
	eg.Go(client.Listen)
	eg.Go(cn.Start)

	if err = eg.Wait(); err != nil {
		slog.Error("application exited with error", "err", err)
		panic(err)
	} else {
		slog.Info("application exited without errors")
	}
}
