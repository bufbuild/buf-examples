# Copyright 2020-2025 Buf Technologies, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import (
    ClassVar as _ClassVar,
    Iterable as _Iterable,
    Mapping as _Mapping,
    Optional as _Optional,
    Union as _Union,
)

DESCRIPTOR: _descriptor.FileDescriptor

class Invoice(_message.Message):
    __slots__ = ("invoice_id", "account_id", "invoice_date", "line_items")
    INVOICE_ID_FIELD_NUMBER: _ClassVar[int]
    ACCOUNT_ID_FIELD_NUMBER: _ClassVar[int]
    INVOICE_DATE_FIELD_NUMBER: _ClassVar[int]
    LINE_ITEMS_FIELD_NUMBER: _ClassVar[int]
    invoice_id: str
    account_id: str
    invoice_date: _timestamp_pb2.Timestamp
    line_items: _containers.RepeatedCompositeFieldContainer[LineItem]
    def __init__(
        self,
        invoice_id: _Optional[str] = ...,
        account_id: _Optional[str] = ...,
        invoice_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...,
        line_items: _Optional[_Iterable[_Union[LineItem, _Mapping]]] = ...,
    ) -> None: ...

class LineItem(_message.Message):
    __slots__ = ("line_item_id", "product_id", "quantity", "unit_price")
    LINE_ITEM_ID_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_ID_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    UNIT_PRICE_FIELD_NUMBER: _ClassVar[int]
    line_item_id: str
    product_id: str
    quantity: int
    unit_price: int
    def __init__(
        self,
        line_item_id: _Optional[str] = ...,
        product_id: _Optional[str] = ...,
        quantity: _Optional[int] = ...,
        unit_price: _Optional[int] = ...,
    ) -> None: ...
