# Dockerfile for MLFlow
FROM python:3.10-slim

# Set the working directory
WORKDIR /mlflow

# Install MLFlow
RUN pip install mlflow

# Copy mlflow configurations and database
COPY ./ml/versioning/mlflow_config.yaml /mlflow/mlflow_config.yaml
#COPY ./ml/versioning/mlflow.db /mlflow/mlflow.db
#COPY ./ml/versioning/mlruns /mlflow/mlruns

# Expose the MLFlow port
EXPOSE 5001

# Command to run MLFlow server
CMD ["mlflow", "server", "--backend-store-uri", "sqlite:///mlflow.db", "--default-artifact-root", "./mlflow/mlruns", "--host", "0.0.0.0", "--port", "5001"]
