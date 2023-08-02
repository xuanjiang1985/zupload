package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"zupload/config"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed static/assets/* templates/*
	f embed.FS
)

func main() {
	configFile := flag.String("c", "", "-c x.yaml")
	flag.Parse()

	conf, err := config.InitConfig(*configFile)
	if err != nil {
		log.Fatalf("load config err: %v", err)
	}

	if conf.Env == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.Default()

	assets, _ := fs.Sub(f, "static/assets")
	frontTmpl, _ := fs.Sub(f, "templates/front")
	app.StaticFS("/static/assets/", http.FS(assets))

	app.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	app.GET("/", func(c *gin.Context) {
		c.FileFromFS("/", http.FS(frontTmpl))
	})

	fmt.Println("open http site 127.0.0.1:8283")
	_ = app.Run(":8283")
}
