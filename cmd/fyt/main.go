package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"fyt/internal"
	"fyt/internal/api"
	"fyt/internal/app/service"
	"fyt/internal/config"
	"fyt/internal/logging"
	"fyt/internal/storage/postgresql"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var env string

	// parse flags, default flag value is `dev`
	flag.StringVar(&env, "env", "dev", "Environment: dev or prod")
	flag.Parse()
	cfgName := env + ".yaml"

	errC, err := run(cfgName)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(cfgName string) (<-chan error, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// config file must be in project root dir, compiled bin must be in /bin dir!!!
	configPath, err := filepath.Abs(dir + "/../" + cfgName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config path:", configPath)

	var cfg *config.Config
	// Create config
	cfg, err = config.NewConfig(configPath)
	if err != nil {
		fmt.Printf("getting config: %s\n", err)
		os.Exit(1)
	}

	// Create logger
	logger, err := logging.GetLogger(cfg.Logger.LoggingLevel)
	if err != nil {
		fmt.Printf("getting logger: %s\n", err)
		os.Exit(1)
	}

	pool, err := NewPostgreSQL(cfg)
	if err != nil {
		// return nil, fmt.Errorf("newDB %w", err)
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "internal.NewPostgreSQL")
	}

	serverAddr := cfg.Server.ServerURL

	errC := make(chan error, 1)

	srv := newServer(cfg, pool)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		logger.Info("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			// logger.Sync()
			pool.Close()
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		logger.Info("Shutdown completed")
	}()

	go func() {
		// logger.Info("Listening and serving", zap.String("address", serverAddr))
		logger.Info("Listening and serving", "address", serverAddr)

		// "ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is
		// ErrServerClosed."
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC, nil

}

func newServer(cfg *config.Config, db *pgxpool.Pool, mws ...mux.MiddlewareFunc) *http.Server {
	r := mux.NewRouter()

	for _, mw := range mws {
		r.Use(mw)
	}

	repoProjects := postgresql.NewProjectRepo(db)          // Project Repository
	svcProjects := service.NewProjectService(repoProjects) // Project Service

	api.NewProjectHandler(svcProjects).Register(r)

	address := cfg.Server.ServerURL

	return &http.Server{
		Handler:           r,
		Addr:              address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
}

func NewPostgreSQL(cfg *config.Config) (*pgxpool.Pool, error) {
	// conn, err := pgx.Connect(context.Background(), cfg.Postgres.PostgresURL)
	// if err != nil {
	// 	return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgx.Connect")
	// }

	pool, err := pgxpool.New(context.Background(), cfg.Postgres.PostgresURL)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgxpool.New")
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pool.Ping")
	}

	return pool, nil
}
