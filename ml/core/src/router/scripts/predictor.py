import os
import mlflow
import joblib

from loguru import logger

from ..models.song_recommender import SongRecommender
from ..schema.schemas import Playlist
from ..scripts.extractor import DatasetExtractor


class Predictor:
	def __init__(self, dataset_name: str = "tracks_features.csv"):
		self.song_recommender = SongRecommender()
		self.dataset_name = dataset_name
		self.model = None
		self.scaler = None

	def load_best_model(self):
		experiment_name = "song_recommender_experiment"
		tracking_uri = "http://mlflow:5001"

		client = mlflow.tracking.MlflowClient(tracking_uri=tracking_uri)
		experiment = client.get_experiment_by_name(experiment_name)
		best_run = client.search_runs(
			experiment_ids=[experiment.experiment_id],
			max_results=1
		)[0]

		model_path = os.path.join(best_run.info.artifact_uri, "model")
		scaler_path = os.path.join(best_run.info.artifact_uri, "scaler.joblib")

		logger.info(f"Loading model from {model_path}")

		self.model = mlflow.sklearn.load_model(model_path)
		self.scaler = joblib.load(scaler_path)

	def predict(self, playlist: Playlist):
		if not self.model or not self.scaler:
			self.load_best_model()

		self.song_recommender.model = self.model
		self.song_recommender.scaler = self.scaler

		full_songs = DatasetExtractor(dataset_name=self.dataset_name).extract()
		self.song_recommender.fit(full_songs)
		recommended_songs = self.song_recommender.predict(playlist, full_songs)
		logger.info(f"Recommended songs: {recommended_songs}")
		return recommended_songs.iloc[0].to_dict()
