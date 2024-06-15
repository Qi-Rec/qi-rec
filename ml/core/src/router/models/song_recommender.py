import pandas as pd
import numpy as np
import joblib

from ..schema.schemas import Playlist, Song

from sklearn.preprocessing import StandardScaler
from sklearn.neighbors import NearestNeighbors


class SongRecommender:
    def __init__(self, n_neighbors=7):
        self.n_neighbors = n_neighbors
        self.model = NearestNeighbors(n_neighbors=self.n_neighbors, algorithm='auto')
        self.scaler = StandardScaler()
        self.features: list[str] = ['danceability', 'energy', 'key', 'mode',
                                    'speechiness', 'acousticness', 'instrumentalness', 'valence', 'tempo',
                                    'time_signature']

    def fit(self, songs: pd.DataFrame):
        X = songs[self.features]
        X_scaled = self.scaler.fit_transform(X)
        self.model.fit(X_scaled)

    def predict(self, playlist: Playlist, full_songs: pd.DataFrame) -> Song:
        playlist_df = pd.DataFrame([song.__dict__ for song in playlist.songs])
        full_df = full_songs

        playlist_features = playlist_df[self.features]
        playlist_features_scaled = self.scaler.transform(playlist_features)
        distances, indices = self.model.kneighbors(playlist_features_scaled)

        recommended_song_indices = np.unique(indices.flatten())
        recommended_songs = full_df.iloc[recommended_song_indices]

        return recommended_songs

    def save_model(self, model_path, scaler_path):
        joblib.dump(self.model, model_path)
        joblib.dump(self.scaler, scaler_path)

    def load_model(self, model_path, scaler_path):
        self.model = joblib.load(model_path)
        self.scaler = joblib.load(scaler_path)


if __name__ == '__main__':
    song_recommender = SongRecommender()

    # Load the dataset
    dataset = pd.read_csv("../../data/tracks_features.csv")

    # Fit the model
    song_recommender.fit(dataset)

    # Save the model
    song_recommender.save_model("dumps/model.joblib", "dumps/scaler.joblib")
