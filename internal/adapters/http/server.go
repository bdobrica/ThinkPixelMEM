package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/bdobrica/ThinkPixelMEM/internal/config"
	"github.com/bdobrica/ThinkPixelMEM/internal/domain"
	"github.com/bdobrica/ThinkPixelMEM/internal/telemetry"
	"github.com/prometheus/client_golang/prometheus"
)

type ReadinessCheck func(context.Context) error
type Server struct {
	http            *http.Server
	shutdownTimeout time.Duration
}

// NewServer assembles operational endpoints and an application handler. A nil
// readiness check fails closed; liveness only reports process responsiveness.
func NewServer(cfg config.HTTPConfig, runtime *telemetry.Runtime, ids *domain.IDGenerator, ready ReadinessCheck, application http.Handler) (*Server, error) {
	if runtime == nil || ids == nil {
		return nil, errors.New("http: telemetry runtime and ID generator are required")
	}
	if cfg.MaxBodyBytes < 1 || cfg.MaxHeaderBytes < 1 || cfg.ShutdownTimeout <= 0 {
		return nil, errors.New("http: invalid limits or shutdown timeout")
	}
	if application == nil {
		application = http.NotFoundHandler()
	}
	requests := prometheus.NewCounterVec(prometheus.CounterOpts{Name: "thinkpixelmem_http_requests_total", Help: "HTTP requests by method and status class."}, []string{"method", "status_class"})
	duration := prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "thinkpixelmem_http_request_duration_seconds", Help: "HTTP request duration by method and status class."}, []string{"method", "status_class"})
	if err := runtime.Registry.Register(requests); err != nil {
		return nil, err
	}
	if err := runtime.Registry.Register(duration); err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /livez", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		if ready == nil || ready(r.Context()) != nil {
			WriteProblem(w, r, errNotReady)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.Handle("GET /metrics", runtime.MetricsHandler())
	mux.Handle("/", application)
	mw := &middleware{ids: ids, propagator: runtime.Propagator, tracer: runtime.TracerProvider.Tracer("thinkpixelmem/http"), maxBody: cfg.MaxBodyBytes, requests: requests, duration: duration}
	return &Server{http: &http.Server{Addr: cfg.Address, Handler: mw.wrap(mux), ReadHeaderTimeout: cfg.ReadHeaderTimeout, ReadTimeout: cfg.ReadTimeout, WriteTimeout: cfg.WriteTimeout, IdleTimeout: cfg.IdleTimeout, MaxHeaderBytes: cfg.MaxHeaderBytes}, shutdownTimeout: cfg.ShutdownTimeout}, nil
}
func (s *Server) Handler() http.Handler { return s.http.Handler }

// Run drains in-flight requests after context cancellation.
func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.http.Addr)
	if err != nil {
		return err
	}
	return s.run(ctx, listener)
}

func (s *Server) run(ctx context.Context, listener net.Listener) error {
	serveErr := make(chan error, 1)
	go func() { serveErr <- s.http.Serve(listener) }()
	select {
	case err := <-serveErr:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()
		if err := s.http.Shutdown(shutdownCtx); err != nil {
			_ = s.http.Close()
			return err
		}
		err := <-serveErr
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
