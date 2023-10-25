package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	api "github.com/shlyapos/echelon/api/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	if err := loadConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	address := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error while dial server: %s", err.Error())
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while connection closing: %s", err.Error())
		}
	}()

	client := api.NewThumbnailsClient(conn)

	asyncFlag := flag.Bool("async", false, "implements asynchronous execution")
	flag.Parse()

	handleArgs(os.Args, *asyncFlag, client)
}

func handleArgs(args []string, async bool, client api.ThumbnailsClient) {
	if async {
		args = args[2:]
	} else {
		args = args[1:]
	}

	var wg sync.WaitGroup
	for _, arg := range args {
		if async {
			wg.Add(1)
			go requestServer(arg, client, &wg)
		} else {
			requestServer(arg, client, nil)
		}
	}

	if async {
		wg.Wait()
	}
}

func requestServer(videoUrl string, client api.ThumbnailsClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	link := getLink(videoUrl)
	response, err := client.Get(context.Background(), &api.GetRequest{Link: link})

	if err != nil {
		log.Fatalf("error while requesting gRPC-server: %s", err.Error())
	}

	fmt.Printf("Thumbnail-file of video %s: %s\n", videoUrl, response.Thumbnail)
}

func getLink(videoUrl string) string {
	urlParse, err := url.Parse(videoUrl)

	if err != nil {
		log.Fatalf("error while parsing url: %s", err.Error())
	}

	urlParams, err := url.ParseQuery(urlParse.RawQuery)

	if err != nil {
		log.Fatalf("error while parsing url params: %s", err.Error())
	}

	link := urlParams.Get("v")

	return link
}

func loadConfig() error {
	viper.SetConfigFile("./cmd/configs/config.yml")
	return viper.ReadInConfig()
}
