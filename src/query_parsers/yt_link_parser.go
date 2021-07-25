package query_parsers

import (
	url2 "net/url"
	"strings"
)

type YtLinkParser struct {}

func (parser YtLinkParser) Support(query Query) bool {
	return strings.Contains(string(query), "youtu")
}

func (parser YtLinkParser) GetMediaId(query Query) (MediaId, error) {
	parsedUrl, err := url2.Parse(string(query))

	if err != nil {
		return "", &QueryParseError{"isn't url"}
	}

	videoId := parsedUrl.Query().Get("v")

	if videoId == "" {
		return "", &QueryParseError{"no 'v' param in url"}
	}

	return MediaId(videoId), nil
}