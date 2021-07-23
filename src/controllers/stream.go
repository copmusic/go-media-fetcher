package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	youtubeDl "github.com/kkdai/youtube/v2"
	youtubeTransport "google.golang.org/api/googleapi/transport"
	youtubeApi "google.golang.org/api/youtube/v3"
	"net/http"
	url2 "net/url"
	"os"
)

type StreamController struct {
	c *gin.Context
}

type StreamControllerError struct {
	error string
}

func (e *StreamControllerError) Error() string {
	return fmt.Sprintf("%s", e.Error())
}

func (sc StreamController) Get(c *gin.Context)  {
	s := c.Query("s")

	if s == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'s' argument not found",
		})
		return
	}

	youtubeVideoId, err := getYoutubeVideoId(s)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	youtubeStreamUrl, err := getYoutubeStreamUrl(youtubeVideoId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"url": youtubeStreamUrl,
	})
}

func getYoutubeStreamUrl(youtubeVideoId string) (string, error) {
	client := youtubeDl.Client{}

	video, _ := client.GetVideo(youtubeVideoId)
	formats := video.Formats.Type("audio/webm")
	formats.Sort()

	format := formats[0]

	var streamUrl string

	if len(format.URL) > 0 {
		streamUrl = format.URL
	} else {
		var err error
		streamUrl, err = client.GetStreamURL(video, &format)

		if err != nil {
			return "", err
		}
	}

	return streamUrl, nil
}

func getYoutubeVideoId(s string) (string, error) {
	videoId, err := getYoutubeVideoIdByLink(s)

	if err != nil {
		videoId, err = getYoutubeVideoIdByKeywords(s)
	}

	return videoId, err
}

func getYoutubeVideoIdByLink(s string) (string, error) {
	parsedUrl, err := url2.Parse(s)

	if err != nil {
		return "", &StreamControllerError{"isn't url"}
	}

	videoId := parsedUrl.Query().Get("v")

	if videoId == "" {
		return "", &StreamControllerError{"no 'v' param in url"}
	}

	return videoId, nil
}

func getYoutubeVideoIdByKeywords(s string) (string, error) {
	youtubeDataApiV3Key, exists := os.LookupEnv("YOUTUBE_DATA_API_V3_KEY")

	if youtubeDataApiV3Key == "" || exists == false {
		panic("YOUTUBE_DATA_API_V3_KEY not found")
	}

	client := &http.Client{
		Transport: &youtubeTransport.APIKey{Key: youtubeDataApiV3Key},
	}

	service, err := youtubeApi.New(client)

	if err != nil {
		return "", err
	}

	call := service.Search.List([]string{"id"}).
		Q(s).
		MaxResults(1).
		Type("video")

	response, err := call.Do()

	if err != nil {
		return "", err
	}

	return response.Items[0].Id.VideoId, nil
}
