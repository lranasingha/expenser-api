package main

import (
	"context"
	"expenser-api/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "github.com/gin-gonic/gin"
import "expenser-api/resource"

func main() {
	fmt.Println("Welcome to Expenser.")

	router := gin.Default()
	router.GET("/health", resource.HealthCheck)
	router.POST("/expenses", resource.AddExpense)
	router.GET("/expenses", resource.GetExpenses)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	server.RegisterOnShutdown(func() {
		log.Print("cleaning up resources before shutdown")
		service.CleanUp()
	})

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quitHandler := make(chan os.Signal)

	signal.Notify(quitHandler, syscall.SIGTERM, syscall.SIGINT)
	receivedSig := <-quitHandler

	log.Printf("Received signal %s. Server is shutting down...", receivedSig)
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", err)
	}
}
