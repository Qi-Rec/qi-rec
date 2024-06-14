from dataclasses import dataclass
from pandas import DataFrame

import pandas as pd


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
		df = pd.read_csv(f"../../data/{self.dataset_name}.csv")
		return df
