package service

import (
	"context"

	api "github.com/shlyapos/echelon/api/proto"
	"github.com/shlyapos/echelon/cmd/server/internal/cache"
	"github.com/shlyapos/echelon/cmd/server/internal/thumbnail"
	"google.golang.org/api/youtube/v3"
)

type thumbnailService struct {
	api.UnimplementedThumbnailsServer
	YoutubeService *youtube.Service
	Cache          *cache.RedisCache
}

func (s thumbnailService) Get(ctx context.Context, r *api.GetRequest) (*api.GetResponse, error) {
	link := r.GetLink()

	thumbBytes, err := s.Cache.GetThumbnail(link)

	if err != nil {
		part := []string{"snippet"}
		videoInfo, err := s.YoutubeService.Videos.List(part).Fields("items/snippet/thumbnails").Id(link).Do()

		if err != nil {
			return nil, err
		}

		thumbs := videoInfo.Items[0].Snippet.Thumbnails
		thumbBytes, err = thumbnail.ThumbnailHandler(thumbs)

		if err != nil {
			return nil, err
		}

		if err := s.Cache.SetThumbnail(link, thumbBytes); err != nil {
			return nil, err
		}
	}

	thumbPath, err := thumbnail.SaveThumbnail(thumbBytes, link)
	if err != nil {
		return nil, err
	}

	response := api.GetResponse{Thumbnail: thumbPath}

	return &response, nil
}

func NewThumbnailService(youtubeService *youtube.Service, cache *cache.RedisCache) *thumbnailService {
	return &thumbnailService{YoutubeService: youtubeService, Cache: cache}
}
