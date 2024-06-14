from pydantic_settings import BaseSettings


class AWSSettings(BaseSettings):
	bucket_name: str = 'qi-rec'
	region: str = 'eu-east-1'
	access_key: str
	secret_key: str


settings = AWSSettings()
