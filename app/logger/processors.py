from structlog.types import EventDict, WrappedLogger


def drop_message_color_key(
    _logger: WrappedLogger, _method: str, event_dict: EventDict
) -> EventDict:
    """
    Uvicorn logs the message a second time in the extra `color_message`, but we don't
    need it. This processor drops the key from the event dict if it exists.
    """
    event_dict.pop("color_message", None)
    return event_dict


def replace_level_with_severity(
    _logger: WrappedLogger, _method: str, event_dict: EventDict
) -> EventDict:
    """
    Replace the level field with a severity field as understood by Stackdriver
    logs.
    """
    if "level" in event_dict:
        event_dict["severity"] = event_dict.pop("level").upper()
    return event_dict
