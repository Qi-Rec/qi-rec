import os

import mlflow
import mlflow.sklearn


def init_mlflow():
	base_path = os.path.dirname(os.path.abspath(__file__))
	config_path = os.path.join(base_path, "mlflow_config.yaml")

	import yaml
	with open(config_path, 'r') as f:
		config = yaml.safe_load(f)

	mlflow.set_tracking_uri(config['tracking_uri'])
	mlflow.set_experiment(config['experiment_name'])
