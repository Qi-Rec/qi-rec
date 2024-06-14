import datetime

from ml.retraining.aws.s3_bucket_setup import S3Bucket
from ml.retraining.src.settings import settings


def save_dataset(dataset_name: str) -> None:
	bucket = S3Bucket(
		bucket_name=settings.bucket_name,
		region=settings.region,
		access_key=settings.access_key,
		secret_key=settings.secret_key
	)

	dataset = bucket.retrieve_dataset(dataset_name)
	with open(
			f"../../../core/src/data/{dataset_name}_{datetime.datetime.now().strftime('%Y-%m-%d_%H-%M-%S')}.csv", "wb"
	) as f:
		f.write(dataset)
