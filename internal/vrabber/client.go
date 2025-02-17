package vrabber

import (
	"context"
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
	host string
	port int
	ctx  context.Context
	in   <-chan *pb.StartDownloadRequest
	out  chan<- *pb.DownloadStatusResponse
}

func NewClient(ctx context.Context, host string, port int, in <-chan *pb.StartDownloadRequest, out chan<- *pb.DownloadStatusResponse) *Client {
	return &Client{
		host: host,
		port: port,
		ctx:  ctx,
		in:   in,
		out:  out,
	}
}

func (c *Client) Start() error {
	defer close(c.out)

	conn, err := grpc.NewClient(c.host+":"+strconv.Itoa(c.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to create grpc client: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		if err = conn.Close(); err != nil {
			slog.Error("failed to close grpc client connection", "err", err)
		}
	}(conn)

	cl := pb.NewDownloadServiceClient(conn)

	stream, err := cl.DownloadVideo(c.ctx)
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
			c.out <- in
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for m := range c.in {
			if err := stream.Send(m); err != nil {
				slog.Error("failed to send to grpc stream", "err", err)
				return
			}
		}
	}()

	wg.Wait()
	return nil
}
