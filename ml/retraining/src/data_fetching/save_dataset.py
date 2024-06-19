import os
from typing import Optional

from ..settings import settings
from retraining.aws.s3_bucket_setup import S3Bucket

from loguru import logger


def save_dataset(dataset_name: str) -> Optional[bool]:
	bucket = S3Bucket(
		bucket_name=settings.bucket_name,
		region=settings.region,
		access_key=settings.access_key,
		secret_key=settings.secret_key
	)

	try:
		dataset = bucket.retrieve_dataset(dataset_name)
	except Exception as e:
		logger.error(f"Failed to retrieve dataset {dataset_name}: {e}")
		raise e

	if not dataset:
		logger.error(f"Dataset {dataset_name} is empty")
		return False

	base_path = os.path.dirname(os.path.abspath(__file__))
	dataset_path = os.path.join(base_path, f"../../../core/src/data/{dataset_name}")
	with open(
			dataset_path, "wb"
	) as f:
		f.write(dataset)
		return True
