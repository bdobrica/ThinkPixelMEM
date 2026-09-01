// Command thinkpixelmem runs the ThinkPixelMEM HTTP service.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	httpadapter "github.com/bdobrica/ThinkPixelMEM/internal/adapters/http"
	"github.com/bdobrica/ThinkPixelMEM/internal/config"
	"github.com/bdobrica/ThinkPixelMEM/internal/domain"
	"github.com/bdobrica/ThinkPixelMEM/internal/telemetry"
	"github.com/bdobrica/ThinkPixelMEM/internal/telemetry/logging"
)

func main() {
	os.Exit(run(os.Args[1:], os.Environ()))
}

func run(args, environ []string) int {
	cfg, err := config.Load(args, environ)
	if err != nil {
		_, _ = os.Stderr.WriteString("thinkpixelmem: invalid configuration\n")
		return 2
	}
	logger, err := logging.New(os.Stderr, cfg.Log.Level)
	if err != nil {
		_, _ = os.Stderr.WriteString("thinkpixelmem: initialize logging failed\n")
		return 1
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	runtime, err := telemetry.New(ctx, cfg.Telemetry)
	if err != nil {
		logger.Error("telemetry initialization failed", slog.String("reason_code", "telemetry_initialization"))
		return 1
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
		defer cancel()
		if shutdownErr := runtime.Shutdown(shutdownCtx); shutdownErr != nil {
			logger.Error("telemetry shutdown failed", slog.String("reason_code", "telemetry_shutdown"))
		}
	}()

	// Readiness intentionally fails closed until the canonical PostgreSQL adapter
	// is introduced in Phase 2. Liveness still proves process responsiveness.
	server, err := httpadapter.NewServer(cfg.HTTP, runtime, domain.NewSystemIDGenerator(), nil, nil)
	if err != nil {
		logger.Error("HTTP initialization failed", slog.String("reason_code", "http_initialization"))
		return 1
	}
	logger.Info("service starting", slog.String("operation", "service.start"))
	if err := server.Run(ctx); err != nil {
		logger.Error("HTTP service stopped unexpectedly", slog.String("reason_code", "http_runtime"))
		return 1
	}
	logger.Info("service stopped", slog.String("operation", "service.stop"))
	return 0
}
