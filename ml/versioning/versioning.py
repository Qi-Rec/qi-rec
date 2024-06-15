import os
import pandas as pd
import mlflow
import mlflow.sklearn


from loguru import logger
from core.src.router.models.song_recommender import SongRecommender
from .mlflow_setup import init_mlflow


class Versioner:
    base_path = os.path.dirname(os.path.abspath(__file__))

    def __init__(self):
        self.song_recommender = SongRecommender()
        self.df = pd.read_csv(os.path.join(self.base_path, "../core/src/data/tracks_features.csv"))

    def commit(self):
        logger.info(f"Versioner: Committing model")
        init_mlflow()

        logger.info("Fitting model")
        self.song_recommender.fit(self.df)
        logger.info("Optimizing hyperparameters")

        model_path = os.path.join(self.base_path, "model.joblib")
        scaler_path = os.path.join(self.base_path, "scaler.joblib")
        pca_path = os.path.join(self.base_path, "pca.joblib")

        logger.info(f"Saving model to {model_path}")
        logger.info(f"Saving scaler to {scaler_path}")
        logger.info(f"Saving pca to {pca_path}")

        self.song_recommender.save_model(
            model_path=model_path,
            scaler_path=scaler_path,
            pca_path=pca_path
        )
        logger.info("Logging model to MLflow")
        with mlflow.start_run() as run:
            mlflow.log_param("n_neighbors", self.song_recommender.n_neighbors)
            mlflow.sklearn.log_model(self.song_recommender.model, "model")
            mlflow.log_artifact(scaler_path)
            mlflow.log_artifact(pca_path)
