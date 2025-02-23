package vrabber

import (
	"context"

	pb "github.com/vrabber/telegram-client/gen/client"
)

type Opts struct {
	Ctx  context.Context
	Host string
	Port int
	In   <-chan *pb.StartDownloadRequest
	Out  chan<- *pb.DownloadStatusResponse
}
