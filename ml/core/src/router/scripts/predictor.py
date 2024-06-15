import os

from ..models.song_recommender import SongRecommender
from ..schema.schemas import Playlist
from ..scripts.extractor import DatasetExtractor


class Predictor:
	def __init__(self, dataset_name: str = "tracks_features.csv"):
		self.song_recommender = SongRecommender()
		self.dataset_name = dataset_name

	def predict(self, playlist: Playlist):
		base_dir = os.path.dirname(os.path.abspath(__file__))
		model_path = os.path.join(base_dir, "../models/dumps/model.joblib")
		scaler_path = os.path.join(base_dir, "../models/dumps/scaler.joblib")
		self.song_recommender.load_model(
			model_path=model_path,
			scaler_path=scaler_path
		)
		full_songs = DatasetExtractor(dataset_name=self.dataset_name).extract()
		recommended_songs = self.song_recommender.predict(playlist, full_songs)
		return recommended_songs.iloc[1].to_dict()
