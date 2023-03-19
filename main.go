package main

import (
	"context"
	"log"
	"time"

	"github.com/romycode/phrased/internal/platform/gin/server"
)

func main() {
	ctx, srv := server.New(context.Background(), "", 443, time.Minute)
	log.Fatal(srv.Run(ctx))
}
