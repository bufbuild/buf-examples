from buf.validate import validate_pb2 as _validate_pb2
from invoice.v1 import invoice_pb2 as _invoice_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateInvoiceRequest(_message.Message):
    __slots__ = ("invoice",)
    INVOICE_FIELD_NUMBER: _ClassVar[int]
    invoice: _invoice_pb2.Invoice
    def __init__(self, invoice: _Optional[_Union[_invoice_pb2.Invoice, _Mapping]] = ...) -> None: ...

class CreateInvoiceResponse(_message.Message):
    __slots__ = ("invoice_id", "version")
    INVOICE_ID_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    invoice_id: str
    version: int
    def __init__(self, invoice_id: _Optional[str] = ..., version: _Optional[int] = ...) -> None: ...
