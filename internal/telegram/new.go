package telegram

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	pb "github.com/vrabber/telegram-client/gen/client"
)

func NewClient(ctx context.Context, token string, in <-chan *pb.DownloadStatusResponse, out chan<- *pb.StartDownloadRequest) (*Client, error) {
	//todo move buffer size to config
	client := &Client{
		ctx: ctx,
		in:  in,
		out: out,
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(client.handler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to init telegram client: %w", err)
	}

	client.bot = b
	return client, nil
}
