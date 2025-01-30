package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/config"
	"github.com/sahidhossen/synmail/src/middleware"
	"github.com/sahidhossen/synmail/src/route"
	"golang.org/x/sync/errgroup"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: gin.DefaultWriter})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Err(err).Msg("unable to load service configuration")
	}
	// config.ConfigureEmail(cfg)
	// gin setup
	e := gin.Default()
	e.Use(middleware.CORSMiddleware())
	e.Use(gin.Recovery())
	e.Use(middleware.Logging())
	e.Use(middleware.ErrorHandlerMiddleware())
	// setup routing and our custom middleware
	SchServer := e.Group("/api")
	// Start V1 group
	V1 := SchServer.Group("/v1")
	route.Router(V1, ctx, cfg)

	http.Handle("/api/", e)
	// start http server
	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        e,
		MaxHeaderBytes: 2 << 20,
	}

	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		// start listening to http connections
		return httpServer.ListenAndServe()
	})

	log.Info().Msg(fmt.Sprintf("Server is listening at port:%d \n", cfg.Port))

	serviceErrors := make(chan error, 1)
	go func() {
		// wait for either the gin server or the http server to return an error
		serviceErrors <- errGroup.Wait()
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	select {
	// wait for service errors
	case err := <-serviceErrors:
		log.Err(err).Msg("service error received")
	// wait for iterrupt/termination signals
	case <-osSignals:
		log.Info().Msg("shutdown signal received")
	}

	// shutdown http server
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Err(err).Msg("error when shutting down http server gracefully")
		if err := httpServer.Close(); err != nil {
			log.Err(err).Msg("error closing http server")
		}
	}
}
