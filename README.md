# Qi-Rec - brand new Spotify Recommendation System

## Main System Design
![Link to PDF](docs/SMC.svg)

## Frontend (Artemii Kister)

## Backend (Dmitry Krasnoyarov)

## ML + MLOps (Arman Tovmasian)

ML Core is a machine learning project that uses Python and the FastAPI framework. It is designed to provide a RESTful API for song recommendation based on a given playlist.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Python
- pip

### Installing

```Future```

## API Endpoints

- `POST /predict`: Predicts the next song based on the given playlist. The request body should be a JSON object that matches the `Playlist` schema. The response will be a `SongResponse` object.

- `POST /commit_model`: Commits a new version of the song recommender model straight into the MLFlow registry. The response will be a JSON object with a message indicating the success of the operation.

## Versioning

The project uses MLflow for versioning the machine learning models. The `Versioner` class in `ml/versioning/versioning.py` is responsible for committing new versions of the model.

## Retraining

The project uses Apache Airflow for scheduled retraining, if there's some changes in original dataset

## Built With

- [Python](https://www.python.org/)
- [FastAPI](https://fastapi.tiangolo.com/)
- [MLflow](https://mlflow.org/)
- [Apache Airflow](https://airflow.apache.org/)
