import os
import pandas as pd
import mlflow
import mlflow.sklearn

from core.src.router.models.song_recommender import SongRecommender
from .mlflow_setup import init_mlflow


class Versioner:
    base_path = os.path.dirname(os.path.abspath(__file__))

    def __init__(self):
        self.song_recommender = SongRecommender()
        self.df = pd.read_csv(os.path.join(self.base_path, "../core/src/data/tracks_features.csv"))

    def commit(self):
        init_mlflow()
        self.song_recommender.fit(self.df)
        model_path = os.path.join(self.base_path, "model.joblib")
        scaler_path = os.path.join(self.base_path, "scaler.joblib")
        self.song_recommender.save_model(
            model_path=model_path,
            scaler_path=scaler_path
        )
        with mlflow.start_run() as run:
            mlflow.log_param("n_neighbors", self.song_recommender.n_neighbors)
            mlflow.sklearn.log_model(self.song_recommender.model, "model")
            mlflow.log_artifact(scaler_path)
