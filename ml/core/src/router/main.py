from fastapi import FastAPI
from schema.schemas import Song, Playlist, SongResponse
from models.song_recommender import SongRecommender


app = FastAPI(
	title="ML Core",
	description="Machine Learning Core API",
	version="0.1.0",
)


@app.on_event("startup")
async def startup_event():
	with open("openapi.json", "w") as file:
		import json
		from fastapi.openapi.utils import get_openapi
		openapi_schema = get_openapi(
			title="ML Core",
			version="0.1.0",
			routes=app.routes
		)
		json.dump(openapi_schema, file)


@app.post("/predict")
async def predict(playlist: Playlist) -> SongResponse:
	song_recommender = SongRecommender()
	song_recommender.load_model(
		model_path="../router/models/dumps/model.joblib",
		scaler_path="../router/models/dumps/scaler.joblib"
	)
	recommended_songs = song_recommender.predict(playlist, playlist.full_songs)
	return recommended_songs.iloc[0].to_dict()
