from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetWeatherRequest(_message.Message):
    __slots__ = ("latitude", "longitude", "forecast_date")
    LATITUDE_FIELD_NUMBER: _ClassVar[int]
    LONGITUDE_FIELD_NUMBER: _ClassVar[int]
    FORECAST_DATE_FIELD_NUMBER: _ClassVar[int]
    latitude: float
    longitude: float
    forecast_date: _timestamp_pb2.Timestamp
    def __init__(self, latitude: _Optional[float] = ..., longitude: _Optional[float] = ..., forecast_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
