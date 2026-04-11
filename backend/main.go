package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"libro-backend/controllers/mainController"
	"libro-backend/middleware/requestctx"
	"libro-backend/repositories"
	"libro-backend/repositories/initRepositories"
	"libro-backend/services/core"
	"libro-backend/statics/configs"
)

func main() {
	_ = godotenv.Load("dev.env")

	cfg, err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}
	logger := requestctx.NewLogger()

	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err = repositories.AssertSchema(db); err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr, Password: cfg.RedisPassword, DB: cfg.RedisDB})
	if err = rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	deps := initRepositories.New(db, rdb)
	ir := repositories.NewInitialRepositories(deps)
	server := core.NewServer(cfg, mainController.DepsFromInitialRepositories(ir), logger)

	listenErrCh := make(chan error, 1)
	go func() {
		listenErrCh <- server.Listen(":" + cfg.AppPort)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		log.Printf("received signal %s", sig)
	case listenErr := <-listenErrCh:
		if listenErr != nil {
			log.Fatal(listenErr)
		}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = server.ShutdownWithContext(ctx)
	sqlDB, _ := db.DB()
	if sqlDB != nil {
		_ = sqlDB.Close()
	}
	_ = rdb.Close()
}
