package server

import (
	"context"
	"fmt"
	"github.com/romycode/phrased/internal/platform/gorilla/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/romycode/phrased/internal/platform/gin/server/handler/system"
	userhandler "github.com/romycode/phrased/internal/platform/gin/server/handler/user"
	"github.com/romycode/phrased/internal/platform/gin/server/handler/websocket"
	"github.com/romycode/phrased/internal/platform/gin/server/middleware"
	"github.com/romycode/phrased/internal/user"
)

type Server struct {
	httpAddr string
	Engine   *gin.Engine

	shutdownTimeout time.Duration
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration) (context.Context, Server) {
	srv := Server{
		Engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
	}

	srv.registerRoutes()
	srv.registerMiddleware()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	v1 := s.Engine.Group("/v1")
	// System
	system.RegisterStatusHandler(v1)
	system.RegisterWelcomeHandler(v1)

	// Api endpoints
	cu := user.NewCreateUserService(user.NewInMemoryRepository(map[string]*user.User{}))
	userhandler.RegisterCreateUserHandler(v1, cu)

	// Websocket
	websocket.RegisterWebsocketHandler(v1, server.NewGorillaWebsocketPromoter())
}

func (s *Server) registerMiddleware() {
	// Recovery middleware
	middleware.RegisterRecovery(s.Engine)
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.Engine,
	}

	go func() {
		if err := srv.ListenAndServeTLS("./go-server.crt", "./go-server.key"); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shutdown", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
