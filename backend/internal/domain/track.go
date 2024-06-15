package domain

type Track struct {
	ID       string
	Name     string   `json:"name"`
	Artists  []Artist `json:"artist"`
	CoverURL string   `json:"cover_link"`
	Link     string   `json:"song_link"`
}

type Artist struct {
	Name string
}

type TrackFeatures struct {
	Danceability     float64
	Energy           float64
	Key              int
	Loudness         float64
	Mode             int
	Speechiness      float64
	Acousticness     float64
	Instrumentalness float64
	Liveness         float64
	Valence          float64
	Tempo            float64
	Type             string
	DurationMs       int
	TimeSignature    int
}
