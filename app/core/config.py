from typing import Literal

from pydantic import computed_field
from pydantic_core import MultiHostUrl
from pydantic_settings import BaseSettings, SettingsConfigDict

from app.logger.custom import LogLevel


class Settings(BaseSettings):
    ENVIRONMENT: Literal["development", "production"] = "development"
    LOG_LEVEL: LogLevel = "INFO"

    DB_HOST: str
    DB_PORT: int
    DB_USER: str
    DB_PASSWORD: str
    DB_NAME: str

    @computed_field
    def DATABASE_URL(self) -> str:
        dsn = MultiHostUrl.build(
            scheme="postgresql",
            username=self.DB_USER,
            password=self.DB_PASSWORD,
            host=self.DB_HOST,
            port=self.DB_PORT,
            path=self.DB_NAME,
        )

        return str(dsn)

    model_config = SettingsConfigDict(
        env_file=".env", env_file_encoding="utf-8", extra="ignore"
    )


settings = Settings()  # type: ignore[reportGeneralTypeIssues]
