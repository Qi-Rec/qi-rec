from pydantic import BaseModel


class Song(BaseModel):
	danceability: float
	energy: float
	key: int
	mode: int
	speechiness: float
	acousticness: float
	instrumentalness: float
	valence: float
	tempo: float
	time_signature: int


class Playlist(BaseModel):
	songs: list[Song]


class SongResponse(BaseModel):
	id: str
