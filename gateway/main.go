package main

import (
	"context"
	"log"
	"moonChat/gateway/handler"
	"moonChat/gateway/route"
	wb "moonChat/gateway/websocket"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	handler.UserCenterConnPool = handler.InitUserCenterGrpcConnPool()

	// 创建WebSocket连接池
	handler.WsPool = wb.NewConnectionPool()
	go handler.WsPool.Run()

	router := gin.Default()
	router.GET("/ws", handler.UpgradeHttpToWs)
	route.RouteUserCenter(router)

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    ":8091",
		Handler: router,
	}

	// 在单独的 goroutine 中启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 创建一个用于超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 使用 WaitGroup 并发关闭所有服务
	var wg sync.WaitGroup
	wg.Add(3) // 三个关闭操作：HTTP服务器、WebSocket连接池、gRPC连接池

	// 并发关闭 HTTP 服务器
	go func() {
		defer wg.Done()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server forced to shutdown: %v", err)
		} else {
			log.Println("HTTP server shutdown successfully")
		}
	}()

	// 并发关闭 WebSocket 连接池
	go func() {
		defer wg.Done()
		handler.WsPool.Shutdown(ctx)
		log.Println("WebSocket pool shutdown successfully")
	}()

	// 并发关闭 gRPC 连接池
	go func() {
		defer wg.Done()
		if handler.UserCenterConnPool != nil {
			handler.UserCenterConnPool.ClearPool()
			log.Println("gRPC connection pool cleared successfully")
		}
	}()

	// 等待所有关闭操作完成或超时
	waitCh := make(chan struct{})
	go func() {
		wg.Wait()
		close(waitCh)
	}()

	select {
	case <-waitCh:
		log.Println("All services shutdown successfully")
	case <-ctx.Done():
		log.Println("Shutdown timeout, some services might not have shutdown properly")
	}

	log.Println("Server exiting")
}
