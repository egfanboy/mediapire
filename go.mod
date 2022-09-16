module github.com/egfanboy/mediapire-manager

go 1.16

require (
	github.com/egfanboy/mediapire-common v0.0.0-20220916222922-d4b407ac03e7
	github.com/egfanboy/mediapire-media-host v0.0.0-20220905144209-41b677b2a6a6
	github.com/go-redis/redis/v9 v9.0.0-beta.2
	github.com/gorilla/mux v1.8.0
	github.com/rs/zerolog v1.27.0

)

// uncomment for local development
// replace github.com/egfanboy/mediapire-common => ../mediapire-common
// replace github.com/egfanboy/mediapire-media-host => ../mediapire-media-host
