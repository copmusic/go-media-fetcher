package query_parsers

import (
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"net/http"
	"os"
)

// Should be always included after all other, this is the "fallback" parser
type YtSearchParser struct {}

func (parser YtSearchParser) Support(query Query) bool {
	return true
}

func (parser YtSearchParser) GetMediaId(query Query) (MediaId, error) {
	youtubeDataApiV3Key, exists := os.LookupEnv("YOUTUBE_DATA_API_V3_KEY")

	if youtubeDataApiV3Key == "" || exists == false {
		panic("YOUTUBE_DATA_API_V3_KEY not found")
	}

	client := &http.Client{
		Transport: &transport.APIKey{Key: youtubeDataApiV3Key},
	}

	service, err := youtube.New(client)

	if err != nil {
		return "", err
	}

	call := service.Search.List([]string{"id"}).
		Q(string(query)).
		MaxResults(1).
		Type("video")

	response, err := call.Do()

	if err != nil {
		return "", err
	}

	return MediaId(response.Items[0].Id.VideoId), nil
}