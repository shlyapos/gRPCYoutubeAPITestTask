package thumbnail

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"google.golang.org/api/youtube/v3"
)

func ThumbnailHandler(thumb *youtube.ThumbnailDetails) ([]byte, error) {
	var thumbLink string

	if thumb.High != nil {
		thumbLink = thumb.High.Url
	} else if thumb.Maxres != nil {
		thumbLink = thumb.Maxres.Url
	} else if thumb.Medium != nil {
		thumbLink = thumb.Medium.Url
	} else if thumb.Standard != nil {
		thumbLink = thumb.Standard.Url
	} else if thumb.Default != nil {
		thumbLink = thumb.Default.Url
	} else {
		return nil, errors.New("video does not have thumbnail")
	}

	return downloadThumbnail(thumbLink)
}

func downloadThumbnail(thumbLink string) ([]byte, error) {
	res, err := http.Get(thumbLink)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func SaveThumbnail(thumbnail []byte, videoId string) (string, error) {
	if err := checkImageDir(); err != nil {
		return "", err
	}

	pathFile := fmt.Sprintf("./images/%s.jpg", videoId)
	file, err := os.Create(pathFile)

	if err != nil {
		return "", err
	}

	defer file.Close()

	file.Write(thumbnail)

	return pathFile, nil
}

func checkImageDir() error {
	info, err := os.Stat("./images")

	if os.IsNotExist(err) || !info.IsDir() {
		if err := os.Mkdir("images", 0777); err != nil {
			return err
		}
	}

	return nil
}
