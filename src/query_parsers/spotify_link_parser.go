package query_parsers

type SpotifyLinkParser struct {}

func (parser SpotifyLinkParser) Support(query Query) bool {
	return false
}

func (parser SpotifyLinkParser) GetMediaId(query Query) (MediaId, error) {
	return "", nil
}
