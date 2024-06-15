from pydantic_settings import BaseSettings, SettingsConfigDict


class AWSSettings(BaseSettings):
	model_config = SettingsConfigDict(env_file='.env', env_file_encoding='utf-8', case_sensitive=True, extra='allow')

	bucket_name: str
	region: str
	access_key: str
	secret_key: str


settings = AWSSettings()
