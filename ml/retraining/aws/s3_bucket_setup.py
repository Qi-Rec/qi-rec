from dataclasses import dataclass
from loguru import logger

import boto3
import botocore.errorfactory as errorfactory


@dataclass
class S3Bucket:
	bucket_name: str
	region: str
	access_key: str
	secret_key: str

	def __post_init__(self):
		self.client = boto3.client(
			"s3",
			region_name=self.region,
			aws_access_key_id=self.access_key,
			aws_secret_access_key=self.secret_key
		)

	def retrieve_dataset(self, dataset_name: str) -> bytes:
		"""Get dataset with dataset_name from the bucket and return it as bytes"""
		try:
			obj = self.client.get_object(Bucket=self.bucket_name, Key=dataset_name)
		except errorfactory.ClientError as ex:
			if ex.response['Error']['Code'] == 'NoSuchKey':
				logger.info(f'No object ```{dataset_name}``` found - returning empty')
				return b''
			raise ex
		return obj["Body"].read()
