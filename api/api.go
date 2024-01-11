package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"jchhay/go-rest-api-gin/config"
	"jchhay/go-rest-api-gin/pkg/db"
)

func init() {
	// Setup environment variables
	config.LoadEnv()
	config.SetupConfig()
}

func Run() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	web := NewRouter()

	fmt.Println("Go API REST Running on port " + config.GetConfig().Server.Port)
	fmt.Println("==================>")

	srv := &http.Server{
		Addr:    ":" + config.GetConfig().Server.Port,
		Handler: web,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish as oppose to waiting indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// Clean up
		db, _ := db.GetDB().DB()
		db.Close()
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server exiting")

}
