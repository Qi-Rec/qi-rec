from router.models.song_recommender import SongRecommender
from router.schema.schemas import Playlist, Song
from extractor import DatasetExtractor


class Predictor:
	def __init__(self, dataset_name: str = "tracks_features.csv"):
		self.song_recommender = SongRecommender()
		self.dataset_name = dataset_name

	def predict(self, playlist: Playlist) -> Song:
		self.song_recommender.load_model(
			model_path="router/models/model.joblib",
			scaler_path="router/models/scaler.joblib"
		)
		full_songs = DatasetExtractor(dataset_name=self.dataset_name).extract()
		recommended_songs = self.song_recommender.predict(playlist, full_songs)
		return recommended_songs.iloc[0].to_dict()
