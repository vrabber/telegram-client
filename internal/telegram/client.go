package telegram

import (
	"context"
	"log/slog"

	pb "github.com/bonefabric/vrabber-client-telegram/gen/client"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Client struct {
	ctx context.Context
	bot *bot.Bot
	in  <-chan *pb.DownloadStatusResponse
	out chan<- *pb.StartDownloadRequest
}

func (c *Client) Listen() error {
	go c.listenInput()
	c.bot.Start(c.ctx)
	return nil
}

func (c *Client) handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	select {
	case <-c.ctx.Done():
		close(c.out)
		return
	default:
		slog.Info("received update", "id", update.Message.Chat.ID, "url", update.Message.Text)
		c.out <- &pb.StartDownloadRequest{
			Url:       update.Message.Text,
			RequestId: update.Message.Chat.ID,
		}
	}
}

func (c *Client) Setup() error {
	return c.setupCommands()
}

// todo request id used as client id, change this in future. Only dev feature for testing
func (c *Client) listenInput() {
	for r := range c.in {
		if _, err := c.bot.SendMessage(c.ctx, &bot.SendMessageParams{
			ChatID: r.RequestId,
			Text:   r.Status.String(),
		}); err != nil {
			slog.Error("failed to send message to telegram:", "err", err, "chat_id", r.RequestId, "text", r.Status.String())
		}
	}
}
