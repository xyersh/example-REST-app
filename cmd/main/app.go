package main

import (
	"fmt"
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

	var listenErr error
	var listener net.Listener

	if cfg.Listen.Type == "sock" {
		// если работаем на сокете
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		slog.Info("create socket")
		socketPath := filepath.Join(appDir, "app.sock")

		slog.Info("listen to unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		slog.Info(fmt.Sprintf("server is listening unix socket %s", socketPath))

	} else {
		//если работает на порту
		slog.Info("listen to TCP")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		slog.Info(fmt.Sprintf("server is listening to TCP on %s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		slog.Error("Error during listener creation", "error", listenErr)
		panic(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.Serve(listener)
	if err != nil {
		slog.Error("Can't serve http", "error", err)
		os.Exit(1)
	}

}
