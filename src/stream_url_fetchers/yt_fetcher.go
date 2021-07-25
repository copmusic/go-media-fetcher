package stream_url_fetchers

import (
	"github.com/kkdai/youtube/v2"
	"go-media-fetcher/src/query_parsers"
)

type YtFetcher struct {
	YtClient youtube.Client
	Query query_parsers.Query
	parsers []query_parsers.QueryParser
}

func NewYtFetcher(Query query_parsers.Query) *YtFetcher {
	return &YtFetcher{
		YtClient: youtube.Client{},
		Query: Query,
		parsers: []query_parsers.QueryParser{
			query_parsers.SpotifyLinkParser{},
			query_parsers.YtLinkParser{},
			query_parsers.YtSearchParser{},
		},
	}
}

func (fetcher *YtFetcher) Fetch() (StreamUrl, error) {
	tries := 0

	for _, parser := range fetcher.parsers {
		if parser.Support(fetcher.Query) == false {
			tries++
			continue
		}

		mediaId, err := parser.GetMediaId(fetcher.Query)

		if err != nil {
			return "", err
		}

		video, err := fetcher.getVideo(mediaId)
		format := fetcher.getFormat(video)

		streamUrl, err := fetcher.getStreamUrl(video, format)

		if err != nil {
			return "", err
		}

		return streamUrl, nil
	}

	if tries == len(fetcher.parsers) {
		return "", &FetchError{"Wrong or unsupported query"}
	}

	return "", &FetchError{"No parsers registered"}
}

func (fetcher *YtFetcher) getVideo(mediaId query_parsers.MediaId) (*youtube.Video, error) {
	video, err := fetcher.YtClient.GetVideo(string(mediaId))

	return video, err
}

func (fetcher *YtFetcher) getFormat(video *youtube.Video) *youtube.Format  {
	formats := video.Formats.Type("audio/webm")
	formats.Sort()
	format := formats[0]

	return &format
}

func (fetcher *YtFetcher) getStreamUrl(video *youtube.Video, format *youtube.Format) (StreamUrl, error) {
	if len(format.URL) > 0 {
		return StreamUrl(format.URL), nil
	}

	streamUrl, err := fetcher.YtClient.GetStreamURL(video, format)

	return StreamUrl(streamUrl), err
}