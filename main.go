package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"morningo/config"
	//_ "morningo/module/schedule" // 定时任务
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.GetEnv().ACCESS_LOG {
		file, err := os.OpenFile(config.GetEnv().ACCESS_LOG_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		if config.GetEnv().DEBUG {
			gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
		} else {
			gin.DefaultWriter = io.MultiWriter(file)
		}
	}

	//fmt.Printf("\nserver pid: %d\n", os.Getpid())
	pid := fmt.Sprintf("%d", os.Getpid())
	_, openErr := os.OpenFile("pid", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openErr == nil {
		ioutil.WriteFile("pid", []byte(pid), 0)
	}

	router := initRouter() // 初始化路由
	//router.Run(":" + config.GetEnv().SERVER_PORT)

	srv := &http.Server{
		Addr:    ":" + config.GetEnv().SERVER_PORT,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
