from pydantic import BaseModel


class Song(BaseModel):
	danceability: float
	energy: float
	key: int
	loudness: float
	mode: int
	speechiness: float
	acousticness: float
	instrumentalness: float
	liveness: float
	valence: float
	tempo: float
	duration_ms: int
	time_signature: int


class Playlist(BaseModel):
	songs: list[Song]


class SongResponse(BaseModel):
	id: str
