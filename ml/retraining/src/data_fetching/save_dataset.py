import os

from ..settings import settings


def save_dataset(dataset_name: str) -> None:
	bucket = S3Bucket(
		bucket_name=settings.bucket_name,
		region=settings.region,
		access_key=settings.access_key,
		secret_key=settings.secret_key
	)

	dataset = bucket.retrieve_dataset(dataset_name)
	base_path = os.path.dirname(os.path.abspath(__file__))
	dataset_path = os.path.join(base_path, f"../../../core/src/data/{dataset_name}.csv")
	with open(
			dataset_path, "wb"
	) as f:
		f.write(dataset)
