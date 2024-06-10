package cmd

// TPL is the template for the server.go file
var TPL = `package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"{{.ProjectName}}/pkg/config"
	"{{.ProjectName}}/pkg/model"
	"{{.ProjectName}}/pkg/router"
)

var (
	configFile = flag.String("c", "../etc/dev.yaml", "config file")
)

func main() {
	flag.Parse()

	config.ParseConfig(*configFile)
	config.InitLogger(config.Config().LogFile)

	model.Connect(
		model.DefaultDriver,
		config.Config().Mysql.Master,
		config.Config().Mysql.Slave)

	if !config.Config().Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(cors.Default())
	router.URLs(r)

	qs := make(chan os.Signal, 1)
	signal.Notify(qs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(context.Background())

	server := http.Server{
		Addr:    config.Config().Address,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err == http.ErrServerClosed {
			log.Info("Stopping everything...")
		} else if err != nil {
			log.Errorf("listen and serve failed: %+v", err)
			cancel()
		}
	}()

	select {
	case <-qs:
		log.Infof("Signal captured, exiting....")
		server.Shutdown(ctx)
		cancel()
	}
}
`
