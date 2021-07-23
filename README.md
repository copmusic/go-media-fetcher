# go-media-fetcher

A simple GoLang microservice for one of my side-projects to get a direct link to audio from YouTube (and other in the future) videos

## Usage
Go to https://console.cloud.google.com/apis/library and get your YouTube Data API v3 key. More details: https://developers.google.com/youtube/v3/getting-started?hl=ru

Insert the received key into .env.local:

    YOUTUBE_DATA_API_V3_KEY={YOUR_API_KEY}

After that run:

    ./run-dev.sh

Then open ./index.html in browser and try to enter a link to a video from YouTube or the name of a song that exists on YouTube

The audio player will start playing the audio track in the best quality presented on YouTube

## TODO

- [ ] Ability to choose the quality of the audio track
- [ ] Recognition of links to playlists
- [ ] Recognition of Spotify links (both single tracks and playlists)
- [ ] Add other sources besides YouTube