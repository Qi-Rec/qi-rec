# Qi-Rec - brand new Spotify Recommendation System

## Main System Design
![Link to PDF](docs/SMC.svg)

## Frontend (Artemii Kister)

The frontend component of the Qi-Rec system provides a user-friendly interface for interacting with the song recommendation engine. Built with modern web technologies, it ensures a seamless experience for users to discover new music based on their current playlists.

### Features

- User Interface: A clean and intuitive design for users to easily input their playlists and view song recommendations.
- Interactive Elements: Dynamic elements that enhance user interaction, such as real-time search and playlist management.

### Technologies

- HTML/CSS: For structuring and styling the web pages.
- JavaScript: To handle user interactions and communicate with the backend API.
- Vue.js: A progressive JavaScript framework for building user interfaces, ensuring a dynamic and responsive application.
- Axios: For making HTTP requests to the backend API, facilitating seamless data exchange between the frontend and backend.


### Prerequisites

- Node.js
- npm (Node Package Manager)

### Installing

1. Clone the repository:

    ```bash 
    git clone https://github.com/your-repository.git
    ```
   
2. Navigate to the `frontend` directory:

    ```bash
    cd frontend
    ```
   
3. Install the required dependencies:

    ```bash
    npm install
    ```
   
### Running the Frontend

4. To start the development server and run the frontend application, use the following command:
    ```bash
    npm run serve
    ```
   

This will launch the application on http://localhost:8080, where you can interact with the Qi-Rec system’s user interface.

### Built With

- Vue.js
- Node.js
- Axios

By following these instructions, you will be able to set up and run the frontend part of the Qi-Rec project, providing a user-friendly interface for song recommendations.

## Backend (Dmitry Krasnoyarov)

The backend component of the Qi-Rec system is responsible for handling user authentication, storing recommendation history, and communicating with the frontend and ML components. It is built using Go, a statically typed, compiled language that is efficient and reliable for backend development.

### Features

- User Authentication: Handles user sign up, sign in, and session management using PostgreSQL.
- Recommendation History: Stores and retrieves user recommendation history using MongoDB.
- API Endpoints: Provides endpoints for the frontend to interact with, facilitating seamless data exchange.

### Technologies

- Go: The backend is written in Go, known for its efficiency and strong support for concurrent programming.
- Gin-Gonic: A web framework written in Go. It features a robust set of libraries for building web applications.
- PostgreSQL: A powerful, open-source object-relational database system used for user authentication and session management.
- MongoDB: A source-available cross-platform document-oriented database program, used for storing recommendation history.

### Prerequisites

- Go
- PostgreSQL
- MongoDB

### Built With

- Go
- Gin-Gonic
- PostgreSQL
- MongoDB

By following these instructions, you will be able to set up and run the backend part of the Qi-Rec project, providing necessary services for user authentication and recommendation history.

## ML + MLOps (Arman Tovmasian)

ML Core is a machine learning project that uses Python and the FastAPI framework. It is designed to provide a RESTful API for song recommendation based on a given playlist.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Python
- pip
- MLFlow

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
