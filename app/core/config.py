from typing import Literal

from pydantic import PostgresDsn
from pydantic_settings import BaseSettings, SettingsConfigDict

from app.logger.custom import LogLevel


class Settings(BaseSettings):
    ENVIRONMENT: Literal["development", "production"] = "development"
    LOG_LEVEL: LogLevel = "INFO"

    DATABASE_URL: PostgresDsn = PostgresDsn(
        "postgresql://postgres:postgres@localhost:5432/postgres"
    )

    model_config: SettingsConfigDict = {
        "env_file": ".env",
        "env_file_encoding": "utf-8",
        "extra": "ignore",
    }


settings = Settings()
