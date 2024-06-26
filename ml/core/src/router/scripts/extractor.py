import os
import pandas as pd

from dataclasses import dataclass
from pandas import DataFrame
from loguru import logger

from retraining.src.data_fetching.save_dataset import save_dataset



@dataclass
class DatasetExtractor:
	"""
	Extracts data from a dataset.
	"""

	def __init__(self, dataset_name: str):
		self.dataset_name = dataset_name

	def extract(self) -> DataFrame:
		"""
		Extracts data from the dataset.
		"""
		try:
			base_path = os.path.dirname(os.path.abspath(__file__))
			df_path = os.path.join(base_path, f"../../data/{self.dataset_name}")
			logger.info(f"Reading dataset {self.dataset_name} from {df_path}")
			df = pd.read_csv(df_path)
		except FileNotFoundError:
			logger.warning(f"Dataset {self.dataset_name} not found in the data folder. Trying to retrieve it from the bucket.")
			try:
				save_dataset(self.dataset_name)
				df = pd.read_csv(df_path)
			except Exception as e:
				raise FileNotFoundError(f"Dataset {self.dataset_name} not found in the data folder or the bucket.") from e
		return df
