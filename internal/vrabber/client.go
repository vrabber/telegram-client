package vrabber

import (
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"sync"

	pb "github.com/vrabber/telegram-client/gen/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	pb.UnimplementedDownloadServiceServer
	opts Opts
}

func NewClient(opts Opts) *Client {
	return &Client{opts: opts}
}

func (c *Client) Start() error {
	defer close(c.opts.Out)

	conn, err := grpc.NewClient(c.opts.Host+":"+strconv.Itoa(c.opts.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to create grpc client: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		if err = conn.Close(); err != nil {
			slog.Error("failed to close grpc client connection", "err", err)
		}
	}(conn)

	cl := pb.NewDownloadServiceClient(conn)

	stream, err := cl.DownloadVideo(c.opts.Ctx)
	if err != nil {
		return fmt.Errorf("failed to download video: %v", err)
	}

	defer func(stream grpc.BidiStreamingClient[pb.StartDownloadRequest, pb.DownloadStatusResponse]) {
		if err = stream.CloseSend(); err != nil {
			slog.Error("failed to close grpc client stream", "err", err)
		}
	}(stream)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			in, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					slog.Error("failed to receive from grpc stream", "err", err)
				}
				break
			}
			c.opts.Out <- in
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for m := range c.opts.In {
			if err := stream.Send(m); err != nil {
				slog.Error("failed to send to grpc stream", "err", err)
				return
			}
		}
	}()

	wg.Wait()
	return nil
}
