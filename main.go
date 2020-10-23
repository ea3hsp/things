package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ea3hsp/test/api"
	rest "github.com/ea3hsp/test/api/http"
	"github.com/ea3hsp/test/repository"
	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	gohandlers "github.com/gorilla/handlers"
)

const (
	// Default config definitions
	defProcName    = "l3-backend-api"
	defBindAddress = "0.0.0.0:3020"
	defRepoName    = "sqlite3"

	// Environment variable names
	envProcName    = "PROC_NAME"
	envBindAddress = "HTTP_BIND_ADDRESS"
	envRepoName    = "REPO_NAME"
)

var (
	// ErrorRepoNotFound repository not found
	ErrorRepoNotFound = errors.New("Repository not found")
)

// config struct definition
type config struct {
	processName  string
	httpBindAddr string
	repoName     string
}

func main() {
	// parse os args
	cfg := loadConfig()
	// Creates logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	level.Info(logger).Log("msg", "hermes things service started")
	// create context
	ctx := context.Background()
	// get repository by os enviroment variable
	rep, err := GetRepo(cfg.repoName, "", "")
	if err != nil {
		level.Error(logger).Log("msg", err.Error())
		os.Exit(1)
	}
	defer func() {
		rep.Close(ctx)
		level.Info(logger).Log("msg", "hermes things service ended")

	}()
	var srv api.Service
	{
		repo := rep
		srv = api.NewService(repo, logger)
	}
	// enpoints creation
	endpoints := api.MakeEndpoints(srv)
	// creates REST API Server
	go func() {
		// banner
		level.Info(logger).Log("msg", fmt.Sprintf("hermes things server listening: %s", cfg.httpBindAddr))
		// service handler
		handler := rest.NewHTTPServer(ctx, endpoints)
		// CORS
		ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
		// Recovery
		rh := gohandlers.RecoveryHandler()
		s := http.Server{
			Addr:         cfg.httpBindAddr,  // configure the bind address
			Handler:      ch(rh(handler)),   // handlers Endpoints + CORS + Recovery
			ReadTimeout:  5 * time.Second,   // max time to read request from the client
			WriteTimeout: 10 * time.Second,  // max time to write response to the client
			IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		}
		// Listen and serve
		s.ListenAndServe()
	}()
	sig := WaitSignal()
	// exit banner
	logger.Log("[Info]", "Exit", "signal", sig.String())
}

// WaitSignal waits for os signal
func WaitSignal() os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			return sig
		}
	}
}

// GetRepo creates repository based on os environment
func GetRepo(repoName, user, password string) (api.Repository, error) {
	switch repoName {
	case "sqlite3":
		r, err := repository.NewSqlite()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return r, nil
	default:
		return nil, ErrorRepoNotFound
	}
}

// env get environment variable or fallback to default one
func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// load config parameters
func loadConfig() *config {
	return &config{
		processName:  env(envProcName, defProcName),
		httpBindAddr: env(envBindAddress, defBindAddress),
		repoName:     env(envRepoName, defRepoName),
	}
}
