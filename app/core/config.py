from typing import Literal

from pydantic_settings import BaseSettings

from app.logger.custom import LogLevel


class Settings(BaseSettings):
    ENVIRONMENT: Literal["development", "production"] = "development"
    LOG_LEVEL: LogLevel = "INFO"


settings = Settings()
