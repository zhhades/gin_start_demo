package initialize

import (
	"context"
	"gin_start_demo/global"
	"gin_start_demo/logger"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func Run() {
	InitConfig()
	router := Routers()

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(global.Config.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Sugar.Fatal("listen: " + err.Error())
		}
	}()
	logger.Sugar.Info("Started server success on port[" + strconv.Itoa(global.Config.Port) + "]")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Sugar.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Sugar.Fatal("Server Shutdown:" + err.Error())
	}
	logger.Sugar.Info("Server exiting")
}
