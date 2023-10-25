package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis"
	api "github.com/shlyapos/echelon/api/proto"
	"github.com/shlyapos/echelon/cmd/server/internal/cache"
	"github.com/shlyapos/echelon/cmd/server/service"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"google.golang.org/grpc"
)

func main() {
	if err := loadConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(viper.GetString("api_key")))

	if err != nil {
		log.Fatalf("error while create youtube service: %s", err.Error())
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: "",
		DB:       0,
	})
	redisCache := cache.NewRedisCache(redisClient)

	grpcServer := grpc.NewServer()
	grpcService := service.NewThumbnailService(youtubeService, redisCache)

	api.RegisterThumbnailsServer(grpcServer, grpcService)

	tcpAddress := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	conn, err := net.Listen("tcp", tcpAddress)
	if err != nil {
		log.Fatalf("error while listening port: %s", err)
	}

	log.Printf("Starting Ports GRPC Server on port %s", viper.GetString("port"))

	err = grpcServer.Serve(conn)
	if err != nil {
		log.Fatalf("error while running grpc server: %s", err)
	}
}

func loadConfig() error {
	viper.SetConfigFile("./cmd/configs/config.yml")
	return viper.ReadInConfig()
}
