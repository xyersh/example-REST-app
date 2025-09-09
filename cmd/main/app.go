package main

import (
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/xyersh/examle-REST-app/internal/config"
	"github.com/xyersh/examle-REST-app/internal/user"
	_ "github.com/xyersh/examle-REST-app/pkg/logging"
)

func main() {
	slog.Info("Create router")
	// log.Println("Create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	slog.Info("Register custom user handler")
	handler := user.NewHandler()
	handler.Register(router)

	Start(router, cfg)
}

func Start(router *httprouter.Router, cfg *config.Config) {
	slog.Info("Start application")

	if cfg.Listen.Type == "sock" {
		// если работаем на сокете
		filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		//если работает на порту

	}

	listener, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		slog.Error("Error during listener creation", "error", err)
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = server.Serve(listener)
	if err != nil {
		slog.Error("Can't serve http", "error", err)
		os.Exit(1)
	}

}
