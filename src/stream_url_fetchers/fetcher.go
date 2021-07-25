package stream_url_fetchers

import "fmt"

type StreamUrl string

type StreamUrlFetcher interface {
	Fetch() (StreamUrl, error)
}

type FetchError struct {
	err string
}

func (e FetchError) Error() string {
	return fmt.Sprintf("%s", e.err)
}