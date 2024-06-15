import pandas as pd
import numpy as np
import joblib

# from ..schema.schemas import Playlist, Song

from sklearn.preprocessing import StandardScaler
from sklearn.neighbors import NearestNeighbors, KNeighborsRegressor
from sklearn.decomposition import PCA
from sklearn.model_selection import GridSearchCV

from loguru import logger


class SongRecommender:
    def __init__(self, n_neighbors=5):
        self.n_neighbors = n_neighbors
        self.model = NearestNeighbors(n_neighbors=self.n_neighbors, algorithm='auto')
        self.regressor = KNeighborsRegressor(n_neighbors=self.n_neighbors)
        self.scaler = StandardScaler()
        self.pca = None
        self.features = ['danceability', 'energy', 'key', 'mode',
                         'speechiness', 'valence', 'tempo']

    def fit(self, songs: pd.DataFrame):
        """
        Fit the model to the data.
        :param songs: DataFrame with songs.
        :return:
        """
        logger.info('Fitting model...')
        X = songs[self.features]
        X_scaled = self.scaler.fit_transform(X)
        n_components = min(len(self.features), len(songs))
        self.pca = PCA(n_components=n_components)
        X_reduced = self.pca.fit_transform(X_scaled)
        self.model.fit(X_reduced)

    def add_noise(self, data, noise_level=0.01):
        """Add noise to the data."""
        logger.info('Adding noise to the data...')
        noise = np.random.normal(0, noise_level, data.shape)
        return data + noise

    def predict(self, playlist, full_songs: pd.DataFrame, noise_level=0.01):
        """Predict songs for a playlist.

        Args:
            playlist (Playlist): Playlist object.
            full_songs (pd.DataFrame): DataFrame with all songs.
            noise_level (float): Noise level.

        Returns:
            pd.DataFrame: Recommended songs.
        """
        logger.info('Predicting songs for the playlist...')
        playlist_df = pd.DataFrame([song.__dict__ for song in playlist.songs])
        full_df = full_songs

        playlist_features = playlist_df[self.features]
        logger.info('Scaling playlist features...')
        playlist_features_scaled = self.scaler.transform(playlist_features)
        logger.info('Reducing playlist features...')
        playlist_features_reduced = self.pca.transform(playlist_features_scaled)

        playlist_features_noisy = self.add_noise(playlist_features_reduced, noise_level)

        distances, indices = self.model.kneighbors(playlist_features_noisy)

        logger.info('Calculating song scores...')
        weights = 1 / (distances + 1e-5)
        recommended_song_indices = np.unique(indices.flatten())
        recommended_songs = full_df.iloc[recommended_song_indices]

        song_scores = np.zeros(len(recommended_songs))
        for i, index in enumerate(recommended_song_indices):
            if i < weights.shape[1]:
                valid_weights = weights[:, i][weights[:, i] < len(weights)]
                song_scores[i] = np.sum(valid_weights)

        recommended_songs['score'] = song_scores
        recommended_songs = recommended_songs.sort_values(by='score', ascending=False)

        return recommended_songs

    def save_model(self, model_path, scaler_path, pca_path):
        """Save model to disk."""
        logger.info('Saving model to disk...')
        joblib.dump(self.model, model_path)
        joblib.dump(self.scaler, scaler_path)
        joblib.dump(self.pca, pca_path)

    def load_model(self, model_path, scaler_path, pca_path):
        """Load model from disk."""
        logger.info('Loading model from disk...')
        self.model = joblib.load(model_path)
        self.scaler = joblib.load(scaler_path)
        self.pca = joblib.load(pca_path)

    def optimize_hyperparameters(self, X, y):
        logger.info('Optimizing hyperparameters...')
        X_scaled = self.scaler.fit_transform(X)
        n_components = min(len(self.features), len(X))
        self.pca = PCA(n_components=n_components)
        X_reduced = self.pca.fit_transform(X_scaled)
        param_grid = {'n_neighbors': [7, 10, 15]}
        grid_search = GridSearchCV(KNeighborsRegressor(), param_grid, cv=5)

        logger.info('Fitting grid search...')

        grid_search.fit(X_reduced, y)
        self.n_neighbors = grid_search.best_params_['n_neighbors']
        self.model = NearestNeighbors(n_neighbors=self.n_neighbors, algorithm='auto')
        self.regressor = KNeighborsRegressor(n_neighbors=self.n_neighbors)
