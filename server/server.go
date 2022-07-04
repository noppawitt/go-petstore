package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	oapimw "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/noppawitt/go-petstore/mongo"
	"github.com/noppawitt/go-petstore/restapi"
	"github.com/noppawitt/go-petstore/service"
	"github.com/noppawitt/go-petstore/spec"
)

func Run(ctx context.Context) error {
	srv, err := new()
	if err != nil {
		return err
	}

	serverCtx, shutdown := context.WithCancel(ctx)

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		shutdown()
	}()

	srv.run(serverCtx)

	return nil
}

type server struct {
	srv *http.Server
}

func (s *server) run(ctx context.Context) {
	errCh := make(chan error)

	// Run the server
	go func() {
		log.Printf("Starting a server at %s\n", s.srv.Addr)

		err := s.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	// Wait for server context to be stopped
	select {
	case err := <-errCh:
		log.Printf("Server error: %s\n", err)
		break
	case <-ctx.Done():
		break
	}

	log.Println("Shutting down a server")

	// Shutdown signal with grace period of 30 seconds
	shutdownCtx, cancelShutdownCtx := context.WithTimeout(context.Background(), 30*time.Second)
	shutdownErrCh := make(chan error)

	// Trigger graceful shutdown
	go func() {
		if err := s.srv.Shutdown(shutdownCtx); err != nil {
			shutdownErrCh <- err
		}

		cancelShutdownCtx()
	}()

	select {
	case err := <-shutdownErrCh:
		log.Printf("Graceful shutdown error: %s\n", err)
	case <-shutdownCtx.Done():
		if err := shutdownCtx.Err(); err == context.DeadlineExceeded {
			log.Println("Graceful shutdown timed out.. forcing exit")
		} else {
			log.Println("Graceful shutdown success")
		}
	}
}

func new() (*server, error) {
	swagger, err := spec.GetSwagger()
	if err != nil {
		return nil, err
	}

	swagger.Servers = nil

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(oapimw.OapiRequestValidator(swagger))

	spec.HandlerFromMux(handler(), r)

	srv := &http.Server{Handler: r, Addr: "0.0.0.0:8080"}

	return &server{
		srv: srv,
	}, nil
}

func handler() *restapi.Handler {
	store := mongo.NewStore()

	h := restapi.NewHandler()
	h.PetService = service.NewPetService(store)

	return h
}
