package controllers

import (
	"github.com/gin-gonic/gin"
	"go-media-fetcher/src/query_parsers"
	"go-media-fetcher/src/stream_url_fetchers"
	"net/http"
)

type StreamController struct {
	c *gin.Context
}

func (sc StreamController) Get(c *gin.Context)  {
	query := c.Query("s")

	fetcher := stream_url_fetchers.NewYtFetcher(query_parsers.Query(query))
	streamUrl, err := fetcher.Fetch()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{
		"url": streamUrl,
	})
}
